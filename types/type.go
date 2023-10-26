package types

import "github.com/patrickmn/go-cache"

type Product struct {
	Price       float64 `json:"price"`
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

type AllCache struct {
	Products *cache.Cache
}
