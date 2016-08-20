package configuration

import (
	. "server"
	. "services"
)

var ApplicationRouter = Router {
	HTTP: map[Hostname] []Route {
		"brianledger.net": []Route {
			{ "/(.*)", ServeStatic("src/static"), },
		},
		"softarc.net": []Route {
			{ "/(.*)", ServeStatic("src/static"), },
		},
		"client.run": []Route {
			{ "/(.*)", ServeStatic("src/static"), },
		},
		"devonshireyaw.com": []Route {
			{ "/(.*)", ServeStatic("src/static"), },
		},
		"nwpaincenter.com": []Route {
			{ "/(.*)", ServeStatic("src/static"), },
		},
	},
	HTTPS: map[Hostname] []Route {
		"brianledger.net": []Route {
			{ "/(.*)", ServeStatic("src/static"), },
		},
		"softarc.net": []Route {
			{ "/(.*)", ServeStatic("src/static"), },
		},
		"client.run": []Route {
			{ "/(.*)", ServeStatic("src/static"), },
		},
		"devonshireyaw.com": []Route {
			{ "/(.*)", ServeStatic("src/static"), },
		},
		"nwpaincenter.com": []Route {
			{ "/(.*)", ServeStatic("src/static"), },
		},
	},
}

