package server

import (
	"regexp"
)

type Hostname string

type Pattern string

type Route struct {
	Pattern string
	Service Service
}

type Router map[Protocol]map[Hostname] []Route

type CompiledRoute struct {
	Pattern *regexp.Regexp
	Service Service
}

func (r Route) Compile () CompiledRoute {
	return CompiledRoute {
		Pattern: regexp.MustCompile(r.Pattern),
		Service: r.Service,
	}
}

type CompiledRouter map[Protocol]map[Hostname] []CompiledRoute