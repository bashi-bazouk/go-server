package configuration

import (
	. "server"
	. "services"
)

var ApplicationRouter = Router {
	HTTP: map[Hostname]map[Pattern] Service {
		"brianledger.net": map[Pattern] Service {
			"/(.*)": ServeStatic("src/static"),
		},
		"softarc.net": map[Pattern] Service {
			"/(.*)": ServeStatic("src/static"),
		},
		"client.run": map[Pattern] Service {
			"/(.*)": ServeStatic("src/static"),
		},
		"devonshireyaw.com": map[Pattern] Service {
			"/(.*)": ServeStatic("src/static"),
		},
		"nwpaincenter.com": map[Pattern] Service {
			"/(.*)": ServeStatic("src/static"),
		},
	},
	HTTPS: map[Hostname]map[Pattern] Service {
		"brianledger.net": map[Pattern] Service {
			"/(.*)": ServeStatic("src/static"),
		},
		"softarc.net": map[Pattern] Service {
			"/(.*)": ServeStatic("src/static"),
		},
		"client.run": map[Pattern] Service {
			"/(.*)": ServeStatic("src/static"),
		},
		"devonshireyaw.com": map[Pattern] Service {
			"/(.*)": ServeStatic("src/static"),
		},
		"nwpaincenter.com": map[Pattern] Service {
			"/(.*)": ServeStatic("src/static"),
		},
	},
}

