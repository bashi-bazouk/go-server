package settings

import (
	"services"
)


var (
	DevelopmentApplication = Application {
		environment: DEVELOPMENT,
		routes: services.ApplicationRoutes,
	}
	ProductionApplication = Application {
		environment: PRODUCTION,
		routes: services.ApplicationRoutes,
	}
)