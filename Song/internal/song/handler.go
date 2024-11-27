package song

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/sytayav/Song/internal/apperror"
	"github.com/sytayav/Song/internal/handlers"
	"github.com/sytayav/Song/pkg/logging"
	"net/http"
)

const (
	songsURL = "/songs"
	songURL  = "/songs/:uuid"
)

type handler struct {
	logger     *logging.Logger
	repository Repository
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, songsURL, apperror.Middleware(h.GetList))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	all, err := h.repository.FindAll(context.TODO())
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}
