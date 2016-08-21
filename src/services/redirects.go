package services

import (
	. "server"
	"net/http"
	"golang.org/x/net/context"
	"strconv"
	"path"
	"utilities"
)

func HandleUpgradeToHTTPS(w http.ResponseWriter, r *http.Request, c *context.Context) {
	url := r.URL
	url.Scheme = "https"

	httpsPort, _ := (*c).Value("Application").(Application).Configuration.Ports[HTTPS]
	if httpsPort != 443 {
		url.Host = ReadHostname(r) + ":" + strconv.Itoa(httpsPort)
	}

	http.Redirect(w, r, url.String(), 301)
}


var UpgradeToHTTPS = Service {
	GET: HandleUpgradeToHTTPS,
	POST: HandleUpgradeToHTTPS,
	PUT: HandleUpgradeToHTTPS,
	PATCH: HandleUpgradeToHTTPS,
	DELETE: HandleUpgradeToHTTPS,
	HEAD: HandleUpgradeToHTTPS,
	OPTIONS: HandleUpgradeToHTTPS,
}


func Link(pathFromRoot string) Service {
	fullPath := path.Join(utilities.Root, pathFromRoot)
	return Service {
		GET: func(w http.ResponseWriter, r *http.Request, c *context.Context) {
			http.ServeFile(w, r, fullPath)
		},
	}
}


func ServeClient (clientName string) Service {
	println("gonna serve", path.Join(utilities.Root, "src/clients", clientName + ".js"))
	return Link(path.Join("src/clients", clientName + ".js"))
}