package server

import (
	"CustomServerTemplate/pkg/util"
	"net/http"
)

func (server *APIServer) Index() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Тут нужно вернуть главную страницу
	}
}

func (server *APIServer) SendHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		pageValue, err := util.ReadFile("./web/index.html")
		if err != nil {
			writer.WriteHeader(404)
			return
		}
		writer.Header().Set("Content-Type", "text/html")
		writer.Write(pageValue)
		return
	}
}

func (server *APIServer) TestAPI() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		pageValue, err := util.ReadFile("./web/testApi.html")
		if err != nil {
			writer.WriteHeader(404)
			return
		}
		writer.Header().Set("Content-Type", "text/html")
		writer.Write(pageValue)
		return

	}
}

var testPage = "<html>" +
	"<head></head>" +
	"<body>" +
	"<h1>Hello</h1><br>" +
	"<h2>World</h2>" +
	"</body>" +
	"</html>"
