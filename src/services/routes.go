package services

import (
	"net/http"
	"utilities"
)

type WebAddress struct {
	Protocol	utilities.HTTPProtocol
	Host 			string
	Port 			int
	Pattern 	string
}

// Routes = HTTPProtocol -> host -> pattern -> ServiceContext -> unit
type host string
type pattern string

type Routes map[utilities.HTTPProtocol]map[host]map[pattern]http.HandlerFunc


var ApplicationRoutes = Routes {
	utilities.HTTP: {
		"devonshireyaw.com": {
			"/": Trace(http.FileServer(http.Dir("/Users/brianpl/Desktop/softarc-go/src/serve/devonshireyaw.com/")).ServeHTTP),
		},
	},
	utilities.HTTPS: {
		"brianledger.net": {
			"/": Trace(http.FileServer(http.Dir("/Users/brianpl/Desktop/softarc-go/src/serve/brianledger.net/")).ServeHTTP),
		},
		"softarc.net": {
			"/": http.FileServer(http.Dir("/Users/brianpl/Desktop/softarc-go/src/serve/softarc.net/")).ServeHTTP,
		},
	},
}