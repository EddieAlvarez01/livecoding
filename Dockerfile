FROM golang:1.18-alpine as build
WORKDIR /usr/src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:3.14
WORKDIR /app
COPY --from=build /usr/src/app .
CMD ["./app"]