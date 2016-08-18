package server

import (
	"regexp"
)

type Hostname string

type Pattern string

type Router map[Protocol]map[Hostname]map[Pattern] Service

type Route struct {
	Pattern *regexp.Regexp
	Service Service
}
type CompiledRouter map[Protocol]map[Hostname] []Route