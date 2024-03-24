package HandlerServer

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func SetHandlerHTTPRouter() *httprouter.Router {
	return httprouter.New()
}

func HandlerFunc(router *httprouter.Router) {

	router.GET("/", Hello)
}

func Hello(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	fmt.Fprintf(res, "fff")
}
