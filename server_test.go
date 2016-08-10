package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const (
	statusNotFound   = "404 Not Found"
	statusBadRequest = "400 Bad Request"
	statusCreated    = "201 Created"
	statusOK         = "200 OK"
	statusNoContent  = "204 No Content"
)

func TestCreateStore(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8000/newsyncstore/test/5s/5s", bytes.NewBufferString(""))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp == nil || resp.Status != statusOK {
		t.Error("TestGetNotExistingKey() failed", err, resp)
	}
}

func TestGetNotExistingKey(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8000/syncstore/test/test_key", bytes.NewBufferString(""))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp == nil || resp.Status != statusNotFound {
		t.Error("TestGetNotExistingKey() failed", err, resp)
	}
}

func TestUpdateNotExistingKey(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8000_temp", bytes.NewBufferString("hello up"))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp == nil || resp.Status != statusNotFound {
		t.Error("TestUpdateNotExistingKey() failed", err, resp)
	}
}

func TestDeleteNotExistingKey(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8000_temp", bytes.NewBufferString(""))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp == nil || resp.Status != statusNotFound {
		t.Error("TestDeleteNotExistingKey() failed", err, resp)
	}
}

func TestStoreAndGet(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8000/syncstore/test/test_key/1m", bytes.NewBufferString("hello"))
	resp, err := client.Do(req)
	if err != nil || resp == nil || resp.Status != statusOK {
		t.Error("TestStoreAndGet() failed", err, resp)
	}
	resp.Body.Close()

	req, err = http.NewRequest("GET", "http://localhost:8000/syncstore/test/test_key", bytes.NewBufferString(""))
	resp, err = client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp == nil || resp.Status != statusOK {
		t.Error("TestStoreAndGet() failed", err, resp)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil || data == nil || string(data) != "hello" {
		t.Error("TestStoreAndGet() failed", err, resp)
	}
}

func TestStoreAndDelete(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8000/syncstore/test/test_key2/1m", bytes.NewBufferString("hello"))
	resp, err := client.Do(req)
	if err != nil || resp == nil || resp.Status != statusOK {
		t.Error("TestStoreAndDelete() failed", err, resp)
	}
	resp.Body.Close()

	req, err = http.NewRequest("DELETE", "http://localhost:8000/syncstore/test/test_key2", bytes.NewBufferString(""))
	resp, err = client.Do(req)
	if err != nil || resp == nil || resp.Status != statusOK {
		t.Error("TestStoreAndDelete() failed", err, resp)
	}
	resp.Body.Close()

	req, err = http.NewRequest("GET", "http://localhost:8000/syncstore/test/test_key2", bytes.NewBufferString(""))
	resp, err = client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp == nil || resp.Status != statusNotFound {
		t.Error("TestStoreAndDelete() failed", err, resp)
	}
}

func TestStoreAndUpdate(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8000/syncstore/test/test_key2/1m", bytes.NewBufferString("hello"))
	resp, err := client.Do(req)
	if err != nil || resp == nil || resp.Status != statusOK {
		t.Error("TestStoreAndUpdate() failed", err, resp)
	}
	resp.Body.Close()

	req, err = http.NewRequest("PUT", "http://localhost:8000/syncstore/test/test_key2", bytes.NewBufferString("hello up"))
	resp, err = client.Do(req)
	if err != nil || resp == nil || resp.Status != statusOK {
		t.Error("TestStoreAndUpdate() failed", err, resp)
	}
	resp.Body.Close()

	req, err = http.NewRequest("GET", "http://localhost:8000/syncstore/test/test_key2", bytes.NewBufferString(""))
	resp, err = client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp == nil || resp.Status != statusOK {
		t.Error("TestStoreAndUpdate() failed", err, resp)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil || data == nil || string(data) != "hello up" {
		t.Error("TestStoreAndUpdate() failed", err, resp)
	}
}

func TestTTL(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8000/syncstore/test/test_key3/1s", bytes.NewBufferString("hello"))
	resp, err := client.Do(req)
	if err != nil || resp == nil || resp.Status != statusOK {
		t.Error("TestTTL() failed", err, resp)
	}
	resp.Body.Close()

	time.Sleep(2 * time.Second)

	req, err = http.NewRequest("GET", "http://localhost:8000/syncstore/test/test_key3", bytes.NewBufferString(""))
	resp, err = client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp == nil || resp.Status != statusNotFound {
		t.Error("TestTTL() failed", err, resp)
	}
}
