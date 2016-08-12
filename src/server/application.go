package server

import (
	"net/http"
	"time"
	"log"
	"os"
	"strconv"
	"fmt"
)

var (
	DevelopmentApplication = Application {
		environment: DEVELOPMENT,
		routes: map[],
	}
	ProductionApplication = Application {
		environment: PRODUCTION,
		routes: ApplicationRoutes,
	}
)

func (app Application) Settings () Settings {
	return EnvironmentSettings[app.environment]
}


func (app Application) EnsureCertificates () {
	var keySettings = app.Settings().Keys

	var _, maybeCertError = os.Stat(keySettings.CertFile)
	var _, maybeKeyError = os.Stat(keySettings.KeyFile)

	if os.IsNotExist(maybeCertError) || os.IsNotExist(maybeKeyError) {
		// Certificate or Key is missing.
		println("Missing one of (%s, %s).", keySettings.CertFile, keySettings.KeyFile)
		if keySettings.SelfSign {
			println("Auto-generating...")
			Sign(keySettings.CSR, keySettings.CertFile, keySettings.KeyFile)
			println("Done.")
		} else {
			log.Fatal("Missing Certificates.")
		}
	} else {
		println("Found certificates.")
	}

}


func (app Application) handlers () map[HTTPProtocol]http.Handler {
	var handlers = map[HTTPProtocol]http.Handler { }
	var ports = app.Settings().Ports
	for _, protocol := range [2]HTTPProtocol { HTTP, HTTPS } {
		var port = ports[protocol]
		var server_multiplexer = http.NewServeMux()
		for host, hostRoutes := range app.routes[protocol] {
			for pattern, handler := range hostRoutes {
				var host_pattern = ""
				if (protocol == HTTP && port != 80) || (protocol == HTTPS && port != 443){
					host_pattern = fmt.Sprintf("%s:%s%s", string(host), strconv.Itoa(port), string(pattern))
				} else {
					host_pattern = fmt.Sprintf("%s%s", string(host), string(pattern))
				}


				server_multiplexer.HandleFunc(host_pattern, handler)
			}
		}
		handlers[protocol] = server_multiplexer
	}
	return handlers
}


func (app Application) Start () {
	app.EnsureCertificates()

	var settings = app.Settings()
	var handlers = app.handlers()

	var httpServer = &http.Server{
			Addr:           ":" + strconv.Itoa(settings.Ports[HTTP]),
			Handler:        handlers[HTTP],
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

	log.Print("Starting HTTP Server at ", settings.Ports[HTTP])
	go httpServer.ListenAndServe()


	var httpsServer = &http.Server{
			Addr:           ":" + strconv.Itoa(settings.Ports[HTTPS]),
			Handler:        handlers[HTTPS],
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

	log.Print("Starting HTTPS Server at ", settings.Ports[HTTPS])
	log.Fatal(httpsServer.ListenAndServeTLS(settings.Keys.CertFile, settings.Keys.KeyFile))

}