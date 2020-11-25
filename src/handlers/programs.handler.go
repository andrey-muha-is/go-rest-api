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

type ProgramsHandler struct {
	repository *repositories.ProgramsRepository
}

func NewProgramsHandler(repository *repositories.ProgramsRepository) *ProgramsHandler {
	return &ProgramsHandler{repository}
}

func (cs *ProgramsHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	zap.S().Debugf("%s: fetch all programs", r.Method)
	data, err := cs.repository.FindAll()

	if err != nil {
		errString := fmt.Sprintf("Error reading programs list: %s\n", err)
		fmt.Fprint(w, errString)
		return
	}

	utils.SendSuccessJson(w, data)
}

func (cs *ProgramsHandler) FindByChannelId(w http.ResponseWriter, r *http.Request) {
	channelID := chi.URLParam(r, "channelID")
	zap.S().Debugf("%s: fetch programs with channel_id=%s", r.Method, channelID)
	data, err := cs.repository.FindByChannelId(channelID)

	if err != nil {
		errString := fmt.Sprintf("Error fetching programs for channel_id=%s, err: %s\n", channelID, err)

		if err == sql.ErrNoRows {
			errString = fmt.Sprintf("No programs for channel with id: %s found\n", channelID)
		}

		fmt.Fprint(w, errString)
		return
	}

	utils.SendSuccessJson(w, data)
}
