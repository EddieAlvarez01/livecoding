package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"sync"
)

type Item struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Value string `json:"value"`
}

func main() {
	app := fiber.New()
	app.Get("/items", getItemsHandler)
	log.Fatalln(app.Listen(":8080"))
}

func getItemsHandler(c *fiber.Ctx) error {
	var items []Item
	var wg sync.WaitGroup
	client := http.Client{Transport: &http.Transport{MaxConnsPerHost: 10}}

	//15 by 15 http request until the 15 items are different
	for {
		for i := 0; i < 15; i++ {
			wg.Add(1)
			go func() {
				item, err := getRandomItem(client)
				if err != nil {
					log.Println(err.Error())
					return
				}
				addRandomItem(&items, *item)
				wg.Done()
			}()
		}
		wg.Wait()
		if len(items) == 15 {
			break
		}
	}
	return c.Status(200).JSON(items)
}

func getRandomItem(client http.Client) (*Item, error) {
	res, err := client.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var newItem Item
	err = json.Unmarshal(bodyBytes, &newItem)
	if err != nil {
		return nil, err
	}
	return &newItem, nil
}

func addRandomItem(items *[]Item, newItem Item) {

	//for the data race in the slice
	var m sync.Mutex

	if validateItem(items, newItem, &m) {
		if len(*items) < 15 {
			m.Lock()
			*items = append(*items, newItem)
			m.Unlock()
		}
	}
}

//validates that the id is different
func validateItem(items *[]Item, itemToValidate Item, m *sync.Mutex) bool {
	m.Lock()
	for _, item := range *items {
		if item.ID == itemToValidate.ID {
			m.Unlock()
			return false
		}
	}
	m.Unlock()
	return true
}
