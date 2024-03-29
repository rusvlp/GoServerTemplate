package server

import (
	"CustomServerTemplate/pkg/util"
	json "encoding/json"
	"net/http"
)

type HandleWrapper struct {
	writer http.ResponseWriter
	err    error
}

func (hw *HandleWrapper) fromJson(v any, r *http.Request) {
	if hw.err != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	hw.err = decoder.Decode(&v)

	if hw.err != nil {
		hw.writer.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (hw *HandleWrapper) toJson(v any) []byte {
	if hw.err != nil {
		return nil
	}

	json, err := json.Marshal(v)
	if err != nil {
		hw.err = err
		hw.writer.WriteHeader(500)
		return nil
	}

	return json
}

func (hw *HandleWrapper) jsonResponse(v any) {
	hw.writer.Header().Set("Content-Type", "application/json")
	hw.writer.Write(hw.toJson(v))
}

func (hw *HandleWrapper) htmlResponse(path string) {
	pageValue, err := util.ReadFile(path)
	if err != nil {
		hw.writer.WriteHeader(404)
		return
	}
	hw.writer.Header().Set("Content-Type", "text/html")
	hw.writer.Write(pageValue)
	return
}

//TODO: Парсер из формы под любую структуру

/* func (hw *HandleWrapper) parseFromForm(v any, r *http.Request) {
	fields : = util.GetFieldsOfStruct(v)

} */
