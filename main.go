package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/leondore/lenslocked/controllers"
	"github.com/leondore/lenslocked/models"
	"github.com/leondore/lenslocked/templates"
	"github.com/leondore/lenslocked/views"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Open database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		log.Fatalf("could not open db connection: %s\n", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Instantiate router and set up routes
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

	usersController := controllers.Users{
		UserService: &models.UserService{DB: db},
	}
	usersController.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"layout-page.gohtml", "signup.gohtml",
	))
	usersController.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"layout-page.gohtml", "signin.gohtml",
	))
	r.Get("/signup", usersController.New)
	r.Post("/users", usersController.Create)
	r.Get("/signin", usersController.SignIn)
	r.Post("/signin", usersController.ProcessSignIn)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	assetsHandler := http.FileServer(http.Dir("assets"))
	r.Get("/assets/*", http.StripPrefix("/assets", assetsHandler).ServeHTTP)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
