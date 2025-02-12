package api

import (
	"encoding/json"
	"github.com/front-go/gateway/internal/model"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	// Сюда подключать будем интересующие службы
	authC AuthClient
}

func NewHandler(authC AuthClient) *Handler {
	return &Handler{
		authC: authC,
	}
}

type RequestData struct {
	Passport Passport `json:"passport"`
	Address  Address  `json:"address"`
	FullName FullName `json:"fullname"`
}

type ResponseData struct {
	Result string `json:"result"`
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
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
			Result: `Пользователь` + " " + req.FullName.Surname + " " + "зарегистрирован",
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

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var req model.UserSignup
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	isOk, err := h.authC.DoSignup(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	log.Println(isOk)
	w.WriteHeader(http.StatusCreated)
}

func AttachHandlers(r chi.Router, handler *Handler) {
	r.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Post("/signup", handler.Signup)
	})
}
