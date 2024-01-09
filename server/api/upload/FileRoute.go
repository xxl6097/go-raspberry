package upload

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-http/server/inter"
	"github.com/xxl6097/go-http/server/route"
	"net/http"
)

type FileRoute struct {
	controller *FileController
}

func NewRoute(ctl *FileController) inter.IRoute {
	opt := &FileRoute{
		controller: ctl,
	}
	return opt
}

func (this *FileRoute) Setup(router *mux.Router) {
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/file/upload",
		Fun:    this.controller.Upload,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodGet,
		Path:   "/file/index",
		Fun:    this.controller.Upload,
		NoAuth: true,
	})

	route.RouterUtil.AddFileServer(router, route.ApiModel{
		Path:           "/file/files/",
		NoAuthByPrefix: true,
	})
	//static file handler.
	//http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))
	//router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("./"))))
}
