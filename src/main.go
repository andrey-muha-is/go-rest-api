package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	
	"./routes"
)

func main() {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/channels", routes.ChannelsRoutes)

	http.ListenAndServe(":3000", r)
}
