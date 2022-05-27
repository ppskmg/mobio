package apiserver

import (
	"github.com/julienschmidt/httprouter"

	"net/http"
)

type Middleware struct {
	corsMiddleware
}

type corsMiddleware struct {
	*handleResponse
}

func (c *corsMiddleware) cors(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "https://mobio.ru")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next(w, r, params)
	}
}
