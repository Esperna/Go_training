// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/create", db.create)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	price := req.URL.Query().Get("price")
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no price specified to %s\n", item)
		return
	}
	f, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price %s specified to %s\n", price, item)
		return
	}
	db[item] = dollars(f)
	if f < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price %s specified to %s\n", db[item], item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item: %q already exists\n", item)
		return
	}
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no item specified\n")
		return
	}
	price := req.URL.Query().Get("price")
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no price specified to %s\n", item)
		return
	}
	f, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price %s specified to %s\n", price, item)
		return
	}
	db[item] = dollars(f)
	if f < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price %s specified to %s\n", db[item], item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}
