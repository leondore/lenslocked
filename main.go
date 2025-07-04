package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/leondore/lenslocked/controllers"
	"github.com/leondore/lenslocked/migrations"
	"github.com/leondore/lenslocked/models"
	"github.com/leondore/lenslocked/templates"
	"github.com/leondore/lenslocked/views"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Set up the database
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		log.Fatalf("could not open db connection: %s\n", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		log.Fatalf("could not migrate: %s\n", err.Error())
		os.Exit(1)
	}

	// Instantiate services
	userService := models.UserService{DB: db}
	sessionService := models.SessionService{DB: db}

	// Set up middleware
	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}

	csrfKey := []byte("GCEv4FtNr6sGzxymtX7fDrPXhAj7ntG6")
	csrfMw := csrf.Protect(csrfKey, csrf.Secure(false))

	// Set up controllers
	usersController := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersController.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"layout-page.gohtml", "signup.gohtml",
	))
	usersController.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"layout-page.gohtml", "signin.gohtml",
	))

	assetsHandler := http.FileServer(http.Dir("assets"))

	// Instantiate router and set up routes
	r := chi.NewRouter()
	r.Use(csrfMw)
	r.Use(umw.SetUser)

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "home.gohtml")),
	))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "contact.gohtml")),
	))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "faq.gohtml")),
	))

	r.Get("/signup", usersController.New)
	r.Post("/users", usersController.Create)
	r.Get("/signin", usersController.SignIn)
	r.Post("/signin", usersController.ProcessSignIn)
	r.Post("/signout", usersController.ProcessSignOut)
	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", usersController.CurrentUser)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	r.Get("/assets/*", http.StripPrefix("/assets", assetsHandler).ServeHTTP)

	// Start server
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
