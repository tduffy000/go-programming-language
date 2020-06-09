package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

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
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func getFields(req *http.Request) (string, float32, error) {
	q := req.URL.Query()
	item, p := q.Get("item"), q.Get("price") // DRY by combining this
	price, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return "", 0.0, err
	} else {
		return item, float32(price), nil
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item, price, err := getFields(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "got error: %v\n", err)
		return
	}
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "item: %q already exists\n", item)
	} else {
		db[item] = dollars(price)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "item created: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item, price, err := getFields(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "got error: %v\n", err)
		return
	}
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item: %q does not exist\n", item)
	} else {
		db[item] = dollars(price)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "item updated: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item: %q does not exists\n", item)
	} else {
		delete(db, item)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "item deleted: %q\n", item)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
