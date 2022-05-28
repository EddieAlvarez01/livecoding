# Prueba de livecoding

_Entrevista para puesto de desarrollador backend, prueba en modalidad livecoding el dia 27/05/2022_

## Requeriminetos

* docker v20+

## Instalación

1. clonar el repositorio

2. correr la aplicación

_correr con docker-compose_
```
docker-compose up -d
```

_correr manual con el Dockerfile_
```
docker build -t livecoding/eddiealvarez .
```
```
docker run -d -p 8080:8080 livecoding/eddiealvarez
```

3. probar el endpoint

```
curl http://localhost:8080/items
```
