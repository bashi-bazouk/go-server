package settings

import (
	"time"
	"services"
)

// Environment = Development | Production
type Environment int
const (
	DEVELOPMENT Environment = iota
	PRODUCTION
)



type Application struct {
	environment	Environment
	routes			services.Routes
}


// Settings

type KeySettings struct {
		 SelfSign bool
		 CertFile string
		 KeyFile  string
		 CSR      CertificateSigningRequest }

type PortSettings map[services.HTTPProtocol]int

type Settings struct {
	Keys KeySettings
	Ports PortSettings
}

// Miscellaneous

type CertificateSigningRequest struct {
	Host       string
	ValidFrom  string
	ValidFor   time.Duration
	IsCA       bool
	RSABits    int
	ECDSACurve string
}
