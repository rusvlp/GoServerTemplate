package server

import (
	"CustomServerTemplate/pkg/util"
	"net/http"
)

func SendHTML(writer http.ResponseWriter, path string) {
	pageValue, err := util.ReadFile(path)
	if err != nil {
		writer.WriteHeader(404)
		return
	}
	writer.Header().Set("Content-Type", "text/html")
	writer.Write(pageValue)
	return
}
