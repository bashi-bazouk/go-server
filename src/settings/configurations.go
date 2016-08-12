package settings



var EnvironmentSettings = map[Environment]Settings {

	DEVELOPMENT: {
		Keys: KeySettings {
			SelfSign: true,
			CertFile: "dev-cert.pem",
			KeyFile: "dev-key.pem",
			CSR: CertificateSigningRequest {
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
	},

}


