package dss

import (
	"encoding/json"
	"net/http"
)

type J struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func writeMessage(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	var j J
	if code > 300 {
		j = J{
			Code:   code,
			Errors: data,
		}
	} else {
		j = J{
			Code: code,
			Data: data,
		}
	}
	json.NewEncoder(w).Encode(j)
}

func writeError(w http.ResponseWriter, code int, err error) {
	writeMessage(w, code, err.Error())
}

func (h *Handler) Moora(w http.ResponseWriter, req *http.Request) {
	r := MooraSpec{}
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		return
	}
	n, err := r.Normalization()
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	res := r.Optimization(n)
	writeMessage(w, http.StatusCreated, res)
}
