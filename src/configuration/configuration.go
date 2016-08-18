package configuration

import (
	"utilities"
	. "server"
)

// Environment = Development | Production
type Environment int
const (
	DEVELOPMENT Environment = iota
	PRODUCTION
)


var EnvironmentSettings = map[Environment]ApplicationSettings{

	DEVELOPMENT: {
		Keys: KeySettings {
			SelfSign: true,
			CertFile: "dev-cert.pem",
			KeyFile: "dev-key.pem",
			CSR: utilities.CertificateSigningRequest {
				Host:      "brianledger.net,softarc.net",
				ValidFrom: "Jan 8 00:00:00 1988",
				ValidFor:  1 << 20,
				IsCA:      true,
				RSABits:   4096,
			},
		},
		Ports: PortSettings {
			HTTP: 8080,
			HTTPS: 4430,
		},
		DefaultHost: "brianledger.net",
	},

	PRODUCTION: {
		Keys: KeySettings {
			SelfSign: false,
			CertFile: "cert.pem",
			KeyFile: "key.pem",
		},
		Ports: PortSettings {
			HTTP: 80,
			HTTPS: 443,
		},
		DefaultHost: "brianledger.net",
	},

}