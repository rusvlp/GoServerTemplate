package server

import "net/http"

func (server *APIServer) Index() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Тут нужно вернуть главную страницу
	}
}

func (server *APIServer) SendHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		/* Проверка запроса на корректный метод. Пока хз, надо ли это
		if request.Method != http.MethodGet {
			writer.WriteHeader(405)
			return
		} */

		//writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Content-Type", "text/html")
		writer.Write([]byte(testPage))
		return
	}
}

func (server *APIServer) TestAPI() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")
	}
}

var testPage = "<html>" +
	"<head></head>" +
	"<body>" +
	"<h1>Hello</h1><br>" +
	"<h2>World</h2>" +
	"</body>" +
	"</html>"
