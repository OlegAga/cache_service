package main

import (
	"fmt"
	"github.com/OlegAga/synccache"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"io/ioutil"
	"net/http"
	"time"
)

var caches = map[string]synccache.CacheI{}

const (
	storePath = ""
)

func CreateStore(c web.C, w http.ResponseWriter, r *http.Request) {
	store := c.URLParams["store"]
	cleanD, err := time.ParseDuration(c.URLParams["cleanD"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	saveD, err := time.ParseDuration(c.URLParams["saveD"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	file := storePath + c.URLParams["store"] + ".db"
	cache := synccache.New(cleanD, saveD, file)
	caches[store] = cache
	fmt.Fprintf(w, "Cache \"%s\" was created Ok!", store)
}

func Create(c web.C, w http.ResponseWriter, r *http.Request) {
	nameStore := c.URLParams["store"]
	key := c.URLParams["key"]
	ttl, err := time.ParseDuration(c.URLParams["ttl"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	store, ok := caches[nameStore]
	if !ok {
		http.Error(w, "Store is not exist", http.StatusNotFound)
	}
	if value, err := ioutil.ReadAll(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		err = store.Set(key, value, ttl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}
	fmt.Fprintf(w, "Key \"%s\" was setted Ok!", key)
}

func Read(c web.C, w http.ResponseWriter, r *http.Request) {
	nameStore := c.URLParams["store"]
	key := c.URLParams["key"]
	store, ok := caches[nameStore]
	if !ok {
		http.Error(w, "Store is not exist", http.StatusNotFound)
		return
	}
	result, err := store.Get(key)
	if err != nil || result == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s", result.([]uint8))
}

func Update(c web.C, w http.ResponseWriter, r *http.Request) {
	nameStore := c.URLParams["store"]
	key := c.URLParams["key"]
	store, ok := caches[nameStore]
	if !ok {
		http.Error(w, "Store is not exist", http.StatusNotFound)
	}
	if value, err := ioutil.ReadAll(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else {
		err = store.Update(key, value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}
	fmt.Fprintf(w, "Key %s was updated Ok!", key)
}

func Delete(c web.C, w http.ResponseWriter, r *http.Request) {
	nameStore := c.URLParams["store"]
	key := c.URLParams["key"]
	store, ok := caches[nameStore]
	if !ok {
		http.Error(w, "Store is not exist", http.StatusNotFound)
	}
	err := store.Remove(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintf(w, "Key %s was remove Ok!", key)
}

func main() {
	goji.Get("/newsyncstore/:store/:cleanD/:saveD", CreateStore)
	goji.Post("/syncstore/:store/:key/:ttl", Create)
	goji.Get("/syncstore/:store/:key", Read)
	goji.Put("/syncstore/:store/:key", Update)
	goji.Delete("/syncstore/:store/:key", Delete)
	goji.Serve()
}
