# cache_service
Simple REST-service for CRUD in cache.

It provides to access key-value pairs stored in synccache https://github.com/OlegAga/synccache

Web Application based on Goji - A web microframework for Golang - http://goji.io/

# Project structure

`server.go`

This file starts your web application and also contains routes definition.  
     
# To start server

`go build && ./cache_service`

## Testing service ##

`go test -test.v`

# Cache service endpoint

`http://yourhost:8000/synccache`
   
## To create store ##

### Request ###

`GET`

`http://yourhost:8000/newsyncstore/:store/:cleanup_duration/:save_to_disk_duration`

### Response ###

`HTTP 200 OK` if created

`HTTP 404 Not found` if creating failed

## To retrieve a key ##

### Request ###

`GET`

`http://yourhost:8000/synccache/:store/:key`

### Response ###

`HTTP 200 OK` with data in body

`HTTP 404 Not Found` if key doesn't exist

## To store a key-value pair ##

### Request ###

`POST`

`http://yourhost:8000/synccache/:store/:key/:ttl`

### Response ###

`HTTP 200 OK` if stored

`HTTP 404 Not found`  if key or value is invalid

## To update a key-value pair ##

### Request ###

`PUT`

`http://yourhost:8000/synccache/:store/:key`

### Response ###

`HTTP 200 OK` if updated

`HTTP 404 Not found` if update key failed

## To delete a key ##

### Request ###

`DELETE`

`http://yourhost:8000/synccache/:store/:key`

### Response ###

`HTTP 200 OK` if removed

`HTTP 404 Not found` if key not removed

zxc
