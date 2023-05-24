package api

import (
	"Library/service"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type routes struct {
	root    *mux.Router
	apiRoot *mux.Router
}

type api struct {
	routes     *routes
	libService service.LibService
	logger     *zap.Logger
}

func Init(root *mux.Router,
	libService service.LibService,
	logger *zap.Logger,
) {
	r := routes{
		root:    root,
		apiRoot: root.PathPrefix("/api").Subrouter(),
	}
	api := api{
		routes:     &r,
		libService: libService,
	}
	api.initService()
}

func (api *api) initService() {
	api.routes.apiRoot.HandleFunc("/v1/sign-up-user", api.SignUpUser).Methods("POST")
}
