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

	tpl, err := views.ParseFS(templates.FS, "layout-page.gohtml", "home.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "layout-page.gohtml", "contact.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "layout-page.gohtml", "faq.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	assetsHandler := http.FileServer(http.Dir("assets"))
	r.Get("/assets/*", http.StripPrefix("/assets", assetsHandler).ServeHTTP)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
