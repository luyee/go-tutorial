package apps

import (
	"github.com/gorilla/mux"
	"github.com/zhyq0826/go-tutorial/devops/apps/app"
	"github.com/zhyq0826/go-tutorial/devops/apps/computer"
	"github.com/zhyq0826/go-tutorial/devops/apps/consul"
	"github.com/zhyq0826/go-tutorial/devops/apps/domain"
	"github.com/zhyq0826/go-tutorial/devops/apps/service"
	"github.com/zhyq0826/go-tutorial/devops/apps/home"
)

// InitRouter init all app route
func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router = domain.InitRoutes(router)
	router = computer.InitRoutes(router)
	router = service.InitRoutes(router)
	router = consul.InitRoutes(router)
	router = app.InitRoutes(router)
	router = home.InitRoutes(router)
	return router
}
