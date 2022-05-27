package apiserver

import "github.com/julienschmidt/httprouter"

type Router interface {
	router() *httprouter.Router
}

type urlReport struct {
	base      string
	send        string
}

type apiUrl struct {
	report *urlReport
}

var (
	url = &apiUrl{
		report: &urlReport{
			base:      "/api/report/",
			send:        "/api/report/send",
		},
	}
)
