package apiserver

import (
	"encoding/json"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

type server struct {
	router *muxRouter
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)

}

type handlers struct {
	report *reportHandle
}

type muxRouter struct {
	*http.ServeMux
	handler    *handlers
	middleware *Middleware
	apiUrl     *apiUrl
}

// Конфигурация роутера
func (mr *muxRouter) configureRouter() {
	url := mr.apiUrl
	mr.Handle(url.report.base, mr.voteRouter())
}

type handleResponse struct {
	logger *zap.Logger
	//errors *store.ErrorStore
}

func (s *handleResponse) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	log.Printf("type: %s; status: %d; host: %s; method: %s; backoff: %s;",
		err.Error(), code, r.Host, r.Method, time.Since(time.Now()))
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *handleResponse) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("type: %s; status: %d; host: %s; method: %s; backoff: %s;",
		"new request", code, r.Host, r.Method, time.Since(time.Now()))
}

func newServer() *server {
	logger, _ := zap.NewProduction()
	defer logger.Sync().Error()
	hr := &handleResponse{
		logger: logger,
	}
	h := &handlers{
		report: &reportHandle{
			handleResponse: hr,
		},
	}
	mwr := &Middleware{
		corsMiddleware:   corsMiddleware{},
	}
	mr := &muxRouter{
		ServeMux:   http.NewServeMux(),
		handler:    h,
		middleware: mwr,
		apiUrl:     url,
	}

	s := &server{
		router: mr,
	}
	s.router.configureRouter()
	return s
}
