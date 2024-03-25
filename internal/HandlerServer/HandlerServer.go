package HandlerServer

import (
	dirServe "WebServerUbuntu/pkg/FunctionsString"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"os"
)

func SetHandlerHTTPRouter() *httprouter.Router {
	return httprouter.New()
}

func HandlerFunc(router *httprouter.Router) {

	router.GET("/", HomePage)
}

// HomePage главная страница мониторинга
func HomePage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	dir, _ := os.Getwd()
	dir = dirServe.GetDir(dir, 2)

	parseFile, errParse := template.ParseFiles(dir + "/web/HomePage/index.html")
	if errParse != nil {
		log.Fatal("Couldn't find the file")
	}

	if errExecuteHTML := parseFile.Execute(res, req); errExecuteHTML != nil {
		log.Fatal("Failed to execute file html")
	}
}
