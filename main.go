package main

import (
	"fmt"
	"net/http"

	"github.com/leondore/lenslocked/controllers"
	"github.com/leondore/lenslocked/templates"
	"github.com/leondore/lenslocked/views"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "home.gohtml")),
	))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "contact.gohtml")),
	))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "faq.gohtml")),
	))

	usersController := controllers.Users{}
	usersController.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"layout-page.gohtml", "signup.gohtml",
	))
	r.Get("/signup", usersController.New)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	assetsHandler := http.FileServer(http.Dir("assets"))
	r.Get("/assets/*", http.StripPrefix("/assets", assetsHandler).ServeHTTP)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
