package services

import (
	. "server"
	. "net/http"
	"golang.org/x/net/context"
	"utilities"
	"path"
	"os"
)

func ServeStatic(staticFolder string) Service {
	doServeStatic := func (w ResponseWriter, r *Request, c context.Context) {
		hostname, _ := utilities.SplitHost(r.URL.Host, -1)
		requestPath := c.Value("groups").([]string)[1]
		workingDirectory, _ := os.Getwd()
		fullPath := path.Join(workingDirectory, staticFolder, hostname, requestPath)
		ServeFile(w, r, fullPath)
	}
	return Service {
		GET: doServeStatic,
		HEAD: doServeStatic,
	}
}