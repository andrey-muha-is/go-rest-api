package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/channels", channelsRoutes)

	http.ListenAndServe(":3000", r)
}

func channelsRoutes(r chi.Router) {
	r.Get("/", GetChannels)
	r.Get("/{channelID}/programs", GetChannelPrograms)
}

func GetChannels(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get list of all channels")
}

func GetChannelPrograms(w http.ResponseWriter, r *http.Request) {
	channelID := chi.URLParam(r, "channelID")

	fmt.Fprintf(w, "Get list of programs for channel %s", channelID)
}
