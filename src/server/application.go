package server


import (
	"net/http"
	"time"
	"log"
	"os"
	"strconv"
	"utilities"
	"golang.org/x/net/context"
	"regexp"
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
	compiledRouter CompiledRouter
	compiledContext context.Context
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
	if app.compiledContext == nil {
		app.compile()
	}

	var protocol Protocol
	if r.TLS != nil {
		protocol = HTTPS
	} else {
		protocol = HTTP
	}

	println("Check URL", r.URL.String(), "::", r.URL.Host)
	var hostname string
	if r.URL.Host == "" {
		hostname = app.Configuration.DefaultHost
	} else {
		hostname, _  = utilities.SplitHost(r.URL.Host, -1)
	}

	println("Hostname is", hostname)


	var longest_matched_subgroups []string
	var longest_match_length int
	var longest_match_service *Service = nil

	for _, route := range app.compiledRouter[protocol][Hostname(hostname)] {
		subgroups := route.Pattern.FindStringSubmatch(r.URL.Path)
		if subgroups != nil && len(subgroups[0]) > longest_match_length {
			longest_matched_subgroups = subgroups
			longest_match_length = len(subgroups[0])
			longest_match_service = &route.Service
		}
	}

	println("?", r.URL.String())
	if longest_match_service != nil {
		handler := longest_match_service.GetHandler(r)
		if handler != nil {
			context := context.WithValue(app.compiledContext, "groups", longest_matched_subgroups)
			handler(w, r, context)
		} else {
			http.Error(w, "INVALID METHOD!!!", 404)
			println("!", r.URL, "INVALID METHOD!!!")
		}
	} else {
		http.Error(w, "NOT FOUND!!!", 404)
		println("!", r.URL.String(), "NOT FOUND!!!")
	}
}


func (app Application) Start () {
	app.compile()
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


// Appendix

func (app Application) compile() {
	app.compiledContext = context.WithValue(context.Background(), "Application", app)

	app.compiledRouter = make(CompiledRouter)
	for protocol, hosts := range app.Router {
		if app.compiledRouter[protocol] == nil {
			app.compiledRouter[protocol] = make(map[Hostname] []Route)
		}
		for host, patterns := range hosts {
			for pattern, service := range patterns {
				app.compiledRouter[protocol][host] = append(app.compiledRouter[protocol][host], Route {
					Pattern: regexp.MustCompile(string(pattern)),
					Service: service,
				})
			}
		}
	}
}