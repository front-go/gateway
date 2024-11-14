package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	srv SrvI
}

func NewHandler(srv SrvI) *Handler {
	return &Handler{
		srv: srv,
	}
}

type RequestData struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ResponseData struct {
	Result string `json:"result"`
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	//h.srv.Summator(inModel)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("hello world"))
	case http.MethodPost:
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		var req RequestData
		err = json.Unmarshal(bytes, &req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		log.Println(req.Address)
		res := ResponseData{
			Result: req.Address + req.Name,
		}
		data, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}

}
