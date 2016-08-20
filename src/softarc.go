package main


import (
	. "server"
	. "configuration"
)


func main() {
	NewApplication(EnvironmentSettings[DEVELOPMENT], ApplicationRouter).Start()
}