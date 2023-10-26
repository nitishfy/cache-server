package main

import (
	"Cache-Server/cache"
	"Cache-Server/middleware"
	"Cache-Server/types"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/patrickmn/go-cache"
	"io"
	"log"
	"net/http"
)

var C = cache.NewCache()

func main() {
	router := httprouter.New()
	router.GET("/product/:id", middleware.CheckCache(getProduct, C))
	router.POST("/products/", CreateProduct)
	log.Fatal(http.ListenAndServe(":8081", router))

}

func getProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	resp, err := http.Get("https://fakestoreapi.com/products/" + id)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	product := types.Product{}
	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Fatal(err)
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
		return
	}

	cache.UpdateFromCache(id, product, C)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func CreateProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var product types.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatal(err)
		return
	}

	jsonData, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
		return
	}

	resp, err := http.Post("https://fakestoreapi.com/products", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()
	fmt.Fprintf(w, "project created successfully")
}
