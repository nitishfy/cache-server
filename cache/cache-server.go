package cache

import (
	"Cache-Server/types"
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

const (
	DefaultExpiration = 5 * time.Minute
	PurgeTime         = 10 * time.Minute
)

func NewCache() *types.AllCache {
	Cache := cache.New(DefaultExpiration, PurgeTime)
	return &types.AllCache{
		Products: Cache,
	}
}

func ReadFromCache(id string, c *types.AllCache) (item []byte, ok bool) {
	product, found := c.Products.Get(id)
	if found {
		res, err := json.Marshal(product.(types.Product))
		if err != nil {
			log.Fatal("Error marshalling JSON")
		}

		return res, true
	}

	return nil, false
}

func UpdateFromCache(id string, product types.Product, c *types.AllCache) {
	c.Products.Set(id, product, cache.DefaultExpiration)
}
