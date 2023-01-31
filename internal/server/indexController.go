package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func (server *APIServer) ParamTest() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		id, err := strconv.Atoi(request.URL.Query().Get("id"))
		if err != nil {
			writer.WriteHeader(404)
			return
		}

		fmt.Fprintf(writer, "Requested ID is: %d", id)
	}
}

func (server *APIServer) PathParamTest() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.Atoi(mux.Vars(request)["id"])
		if err != nil {
			writer.WriteHeader(404)
			return
		}
		fmt.Fprintf(writer, "Requested ID is: %d", id)
	}
}

var testPage = "<html>" +
	"<head></head>" +
	"<body>" +
	"<h1>Hello</h1><br>" +
	"<h2>World</h2>" +
	"</body>" +
	"</html>"
