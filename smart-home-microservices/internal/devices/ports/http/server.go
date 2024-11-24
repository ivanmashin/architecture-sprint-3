package http

import (
	"encoding/json"
	"net/http"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/config"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/domain"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/ports/http/middleware"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/usecase"
)

func NewHTTPServer(cfg config.Config, app *usecase.App) *http.Server {
	h := handler{app: app}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /homes", h.GetHomes)
	mux.HandleFunc("GET /homes/{homeID}/devices", h.GetDevicesInHome)

	mux.HandleFunc("POST /devices", h.CreateDevice)
	mux.HandleFunc("GET /devices/{deviceID}", h.GetDevice)
	mux.HandleFunc("PUT /devices/{deviceID}", h.UpdateDevice)
	mux.HandleFunc("DELETE /devices/{deviceID}", h.DeleteDevice)
	mux.HandleFunc("POST /devices/{deviceID}/toggle", h.ToggleDevice)

	withAuthorization := middleware.NewAuthorization(cfg)
	return &http.Server{
		Addr:    cfg.HTTPServerAddress,
		Handler: middleware.WithLogger(withAuthorization(mux)),
	}
}

type handler struct {
	converter Converter
	app       *usecase.App
}

func (h *handler) GetHomes(w http.ResponseWriter, r *http.Request) {
	result, err := h.app.GetHomes.Handle(r.Context(), middleware.MustExtractUserID(r.Context()), struct{}{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(h.converter.HomesFromDomain(result))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetDevicesInHome(w http.ResponseWriter, r *http.Request) {
	homeID := r.PathValue("homeID")
	result, err := h.app.GetDevicesInHome.Handle(r.Context(), middleware.MustExtractUserID(r.Context()),
		domain.ID(homeID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(h.converter.DevicesFromDomain(result))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetDevice(w http.ResponseWriter, r *http.Request) {
	deviceID := r.PathValue("deviceID")
	result, err := h.app.GetDeviceByID.Handle(r.Context(), middleware.MustExtractUserID(r.Context()),
		domain.ID(deviceID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(h.converter.DeviceFromDomain(result))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	homeID := r.PathValue("homeID")
	request := new(Device)
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.app.CreateDevice.Handle(r.Context(), middleware.MustExtractUserID(r.Context()),
		h.converter.DeviceToDomain(domain.ID(homeID), request))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) ToggleDevice(w http.ResponseWriter, r *http.Request) {
	deviceID := r.PathValue("deviceID")
	request := new(bool)
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.app.ToggleDevice.Handle(r.Context(), middleware.MustExtractUserID(r.Context()), usecase.ToggleCommand{
		DeviceID: domain.ID(deviceID),
		On:       *request,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	deviceID := r.PathValue("deviceID")
	request := new(DeviceUpdate)
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.app.UpdateDevice.Handle(r.Context(), middleware.MustExtractUserID(r.Context()), usecase.UpdateDeviceCommand{
		DeviceID:   domain.ID(deviceID),
		DeviceName: request.Name,
		On:         request.On,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	deviceID := r.PathValue("deviceID")
	err := h.app.DeleteDevice.Handle(r.Context(), middleware.MustExtractUserID(r.Context()), domain.ID(deviceID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
