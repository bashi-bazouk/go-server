package configuration

import (
	. "server"
	. "services"
)

var ApplicationRouter = Router {
	HTTP: map[Hostname] []Route {
		"brianledger.net": []Route {
			{ "/(.*)", UpgradeToHTTPS, },
		},
		"softarc.net": []Route {
			{ "/(.*)", UpgradeToHTTPS, },
		},
		"client.run": []Route {
			{ "/(.*)", UpgradeToHTTPS, },
		},
		"devonshireyaw.com": []Route {
			{ "/(.*)", ServeStatic, },
		},
		"nwpaincenter.com": []Route {
			{ "/(.*)", ServeStatic, },
		},
	},
	HTTPS: map[Hostname] []Route {
		"brianledger.net": []Route {
			{ "/client.js", ServeClient("brianledger_public"), },
			{ "/cdn/(.*)", ServeStatic, },
			{ "/(.*)", ServeStatic, },
		},
		"softarc.net": []Route {
			{ "/(.*)", ServeStatic, },
		},
		"client.run": []Route {
			{ "/(.*)", ServeStatic, },
		},
	},
}

