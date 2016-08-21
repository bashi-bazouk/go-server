package server

import (
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"regexp"
	"strconv"
)

type RequestHandler func(http.ResponseWriter, *http.Request, *context.Context)

type Service struct {
	GET			RequestHandler
	POST		RequestHandler
	PUT			RequestHandler
	PATCH		RequestHandler
	DELETE	RequestHandler
	HEAD		RequestHandler
	OPTIONS	RequestHandler
}

func (h Service) GetHandler(r *http.Request) func(http.ResponseWriter, *http.Request, *context.Context) {
	switch strings.ToUpper(r.Method) {
	case "", "GET":
		return h.GET
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

func (p Protocol) String () string {
	switch p {
	case HTTP: return "HTTP"
	case HTTPS: return "HTTPS"
	default: panic("Invalid Protocol")
	}
}



// Random utilities for services.

func GetHTTPPort(c *context.Context) int {
	return (*c).Value("Application").(Application).Configuration.Ports[HTTP]
}

func GetHTTPSPort(c *context.Context) int {
	return (*c).Value("Application").(Application).Configuration.Ports[HTTPS]
}

func ReadImpliedPort(r *http.Request) int {
	if r.TLS == nil {
		return 80
	} else {
		return 443
	}
}

func ReadHostnameAndPort(r *http.Request) (hostname string, port int) {
	match_subgrouped := regexp.MustCompile("([^:]+)(?::([1-9][0-9]+))?")
	submatches := match_subgrouped.FindStringSubmatch(r.Host)

	if len(submatches) == 0 {
		hostname = ""
		port = ReadImpliedPort(r)

	} else if len(submatches) == 2 {
		hostname = submatches[1]
		port = ReadImpliedPort(r)

	} else { // len(submatches) == 3
		hostname = submatches[1]
		port, _ = strconv.Atoi(submatches[2])
	}

	return hostname, port
}

func ReadHostname(r *http.Request) string {
	hostname, _ := ReadHostnameAndPort(r)
	return hostname
}

func ResolveHostname(r *http.Request, c *context.Context) string {
	// Maps "" and "localhost" to default_hostname
	hostname := ReadHostname(r)
	if hostname == "localhost" || hostname == "" {
		hostname = (*c).Value("Application").(Application).Configuration.DefaultHost
	}
	return hostname
}