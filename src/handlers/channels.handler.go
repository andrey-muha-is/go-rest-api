package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"../repositories"
	"../utils"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type ChannelsHandler struct {
	repository *repositories.ChannelsRepository
}

func NewChannelsHandler(repository *repositories.ChannelsRepository) *ChannelsHandler {
	return &ChannelsHandler{repository}
}

func (cs *ChannelsHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	zap.S().Debugf("%s: fetch all channels", r.Method)

	data, err := cs.repository.FindAll()

	if err != nil {
		errString := fmt.Sprintf("Error reading channels list: %s\n", err)
		fmt.Fprint(w, errString)
		return
	}

	utils.SendSuccessJson(w, data)
}

func (cs *ChannelsHandler) FindById(w http.ResponseWriter, r *http.Request) {
	channelID := chi.URLParam(r, "channelID")
	zap.S().Debugf("%s: fetch channel with id = %s", r.Method, channelID)
	data, err := cs.repository.FindById(channelID)

	if err != nil {
		errString := fmt.Sprintf("Error fetching channel with id=%s, err: %s\n", channelID, err)

		if err == sql.ErrNoRows {
			errString = fmt.Sprintf("Channel with id: %s not exist\n", channelID)
		}

		fmt.Fprint(w, errString)
		return
	}

	utils.SendSuccessJson(w, data)
}
