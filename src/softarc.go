package main


import (
	"server"
	. "configuration"
)

var Application = server.Application {
	Configuration: EnvironmentSettings[DEVELOPMENT],
	Router: ApplicationRouter,
}

func main() {
	Application.Start()
}