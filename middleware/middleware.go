package middleware

import (
	"Cache-Server/cache"
	"Cache-Server/types"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func CheckCache(f httprouter.Handle, c *types.AllCache) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		p.ByName("id")
		res, ok := cache.ReadFromCache("id", c)
		if ok {
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			return
		}

		// if not found, receive from controller and put in cache
		log.Println("From controller")
		f(w, r, p)
	}
}
