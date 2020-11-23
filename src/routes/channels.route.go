package routes

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
)

func ChannelsRoutes(r chi.Router) {
	r.Get("/", getChannels)
	r.Get("/{channelID}/programs", getChannelPrograms)
}

func getChannels(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get list of all channels")
}

func getChannelPrograms(w http.ResponseWriter, r *http.Request) {
	channelID := chi.URLParam(r, "channelID")

	fmt.Fprintf(w, "Get list of programs for channel %s", channelID)
}