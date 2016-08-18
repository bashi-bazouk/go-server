package server

import (
	"net/http"
	"strings"
	"golang.org/x/net/context"
)

type Service struct {
	GET func(http.ResponseWriter, *http.Request, context.Context)
	POST func(http.ResponseWriter, *http.Request, context.Context)
	PUT func(http.ResponseWriter, *http.Request, context.Context)
	PATCH func(http.ResponseWriter, *http.Request, context.Context)
	DELETE func(http.ResponseWriter, *http.Request, context.Context)
	HEAD func(http.ResponseWriter, *http.Request, context.Context)
	OPTIONS func(http.ResponseWriter, *http.Request, context.Context)
}

func (h Service) GetHandler(r *http.Request) func(http.ResponseWriter, *http.Request, context.Context) {
	switch strings.ToUpper(r.Method) {
	case "", "GET":
		return  h.GET
	case "POST":
		return h.POST
	case "PUT":
		return h.PUT
	case "PATCH":
		return h.PATCH
	case "DELETE":
		return h.DELETE
	case "HEAD":
		return h.HEAD
	case "OPTIONS":
		return h.OPTIONS
	default:
		return nil
	}
}

type Protocol int
const (
	HTTP Protocol = iota
	HTTPS
)



