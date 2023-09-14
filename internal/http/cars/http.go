package cars

import (
	"encoding/json"
	"errors"
	"github.com/iamvbr/vehicle-service/internal/models"
	"github.com/iamvbr/vehicle-service/internal/services"
	"io"
	"net/http"
	"strings"
)

type handler struct {
	svc services.Car
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Create(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodGet:
		if strings.HasSuffix(r.URL.Path, "cars") {
			h.Index(w, r)
		} else {
			h.Read(w, r)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func New(svc services.Car) *handler {
	return &handler{svc: svc}
}

var internalServerError = []byte(`{"message":"internal server error"}`)

// TODO: this is not the right way to handle errors, and should come from a common place
func formatError(err error) []byte {
	v, _ := json.Marshal(map[string]string{"message": err.Error()})
	return v
}
func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	cars, err := h.svc.Find(r.Context(), r.URL.Query().Get("category"))
	// TODO: We can have common response handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(internalServerError)
		return
	}

	response, _ := json.Marshal(cars)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *handler) Read(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	car, err := h.svc.Get(r.Context(), id)
	if err != nil {
		//TODO: it might not always be 400
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatError(err))
		return
	}

	response, _ := json.Marshal(car)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatError(err))
		return
	}

	var m models.Car
	if err := json.Unmarshal(b, &m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatError(errors.New("invalid input data. Please check the request body")))
		return
	}

	car, err := h.svc.Create(r.Context(), &m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatError(err))
		return
	}

	response, _ := json.Marshal(car)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatError(err))
		return
	}

	var m models.Car
	if err := json.Unmarshal(b, &m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatError(errors.New("invalid input data. Please check the request body")))
		return
	}

	car, err := h.svc.Update(r.Context(), id, &m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatError(err))
		return
	}

	response, _ := json.Marshal(car)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
