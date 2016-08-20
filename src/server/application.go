package server


import (
	"net/http"
	"time"
	"log"
	"os"
	"strconv"
	"utilities"
	"golang.org/x/net/context"
)


// Application Settings

type KeySettings struct {
		 SelfSign bool
		 CertFile string
		 KeyFile  string
		 CSR      utilities.CertificateSigningRequest }

type PortSettings map[Protocol]int

type ApplicationSettings struct {
	Keys				KeySettings
	Ports				PortSettings
	DefaultHost	string
}


// Application

type Application struct {
	Configuration ApplicationSettings
	Router Router
	compiledRouter *CompiledRouter
	compiledContext *context.Context
}


func NewApplication(configuration ApplicationSettings, router Router) (app Application) {
	app.Configuration = configuration
	app.Router = router

	compiledContext := context.WithValue(context.Background(), "Application", app)
	app.compiledContext = &compiledContext

	var compiledRouter = make(CompiledRouter)
	for protocol, hosts := range router {
		var routesByHost = make(map[Hostname] []CompiledRoute)
		for host, routes := range hosts {
			for _, route := range routes {
				routesByHost[host] = append(routesByHost[host], route.Compile())
			}
		}
		compiledRouter[protocol] = routesByHost
	}

	app.compiledRouter = &compiledRouter

	return app
}


func (app Application) EnsureCertificates () {
	var keySettings = app.Configuration.Keys

	var _, maybeCertError = os.Stat(keySettings.CertFile)
	var _, maybeKeyError = os.Stat(keySettings.KeyFile)

	if os.IsNotExist(maybeCertError) || os.IsNotExist(maybeKeyError) {
		// Certificate or Key is missing.
		println("Missing one of (%s, %s).", keySettings.CertFile, keySettings.KeyFile)
		if keySettings.SelfSign {
			println("Auto-generating...")
			utilities.Sign(keySettings.CSR, keySettings.CertFile, keySettings.KeyFile)
			println("Done.")
		} else {
			log.Fatal("Missing Certificates.")
		}
	} else {
		println("Found certificates.")
	}
}


func (app Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var protocol Protocol
	if r.TLS != nil {
		protocol = HTTPS
	} else {
		protocol = HTTP
	}

	println("Check URL", r.URL.String(), "::", r.URL.Host)
	var hostname string
	if r.Host == "localhost" {
		hostname = app.Configuration.DefaultHost
	} else {
		hostname, _  = utilities.SplitHost(r.URL.Host, -1)
	}

	println("Hostname is", hostname)

	println("?", r.URL.String())
	println("Got something...", (*app.compiledRouter)[protocol])
	for _, route := range (*app.compiledRouter)[protocol][Hostname(hostname)] {
		subgroups := route.Pattern.FindStringSubmatch(r.URL.Path)
		if subgroups != nil {
			context := context.WithValue(*app.compiledContext, "groups", subgroups)
			handler := route.Service.GetHandler(r)
			println("!", "Found a handler!")
			handler(w, r, context)
			println("Done")
			return
		}
	}

	println("!", r.URL.String(), "404 Not Found")
	http.Error(w, "404 Not Found", 404)

}


func (app Application) Start () {
	app.EnsureCertificates()

	var configuration = app.Configuration

	var httpServer = &http.Server{
			Addr:           ":" + strconv.Itoa(configuration.Ports[HTTP]),
			Handler:        app,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

	log.Print("Starting HTTP Server at ", configuration.Ports[HTTP])
	go httpServer.ListenAndServe()


	var httpsServer = &http.Server{
			Addr:           ":" + strconv.Itoa(configuration.Ports[HTTPS]),
			Handler:        app,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

	log.Print("Starting HTTPS Server at ", configuration.Ports[HTTPS])
	log.Fatal(httpsServer.ListenAndServeTLS(configuration.Keys.CertFile, configuration.Keys.KeyFile))

}