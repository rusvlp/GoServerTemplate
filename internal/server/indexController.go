package server

import (
	"net/http"
)

func (server *APIServer) Index() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		SendHTML(writer, "./web/index.html")
		return
	}
}

func (server *APIServer) TestAPI() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		SendHTML(writer, "./web/testApi.html")
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
