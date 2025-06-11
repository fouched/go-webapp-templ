package main

import (
	"database/sql"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-webapp-templ/internal/config"
	"github.com/fouched/go-webapp-templ/internal/driver"
	"github.com/fouched/go-webapp-templ/internal/handlers"
	"github.com/fouched/go-webapp-templ/internal/render"
	"github.com/fouched/go-webapp-templ/internal/repo"
	"github.com/gorilla/schema"
	"github.com/jaswdr/faker/v2"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// session - must be available in main package for middleware
var session *scs.SessionManager

func main() {
	app, err := newApp()
	if err != nil {
		log.Fatal(err)
	}
	// close database connection pool after application has stopped
	defer app.DB.Close()

	// seed the database ?
	//seed(app.DB)

	// Create handlers instance with dependency
	h := handlers.NewHandlers(app)

	mux := routes(h)
	srv := &http.Server{
		Addr:    app.Addr,
		Handler: mux,
	}

	app.InfoLog.Println("Starting server on", app.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func newApp() (*config.App, error) {
	addr := ":9080"
	dsn := "host=localhost port=5432 user=postgres password=password dbname=webapp-templ sslmode=disable"

	infoLog, errorLog := initLogger()
	db, err := initDatabase(dsn, errorLog)
	if err != nil {
		return nil, err
	}

	session = initSession()
	decoder := initDecoder()
	templates := initTemplates()

	app := &config.App{
		Addr:          addr,
		DSN:           dsn,
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		DB:            db,
		Session:       session,
		TemplateCache: templates,
		Decoder:       decoder,
		Repo: config.Repo{
			CustomerRepo: repo.NewCustomerRepo(db),
			//CustomerRepo: repo.NewCustomerRepoUpperDB(db.SQL), // if you want to change implementations
		},
	}

	// set up template rendering
	render.NewRenderer(app)

	return app, nil
}

func initLogger() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}

func initDatabase(dsn string, errorLog *log.Logger) (*sql.DB, error) {
	dbPool, err := driver.ConnectSQL(dsn)
	if err != nil {
		errorLog.Println("Failed to connect to DB:", err)
		return nil, err
	}
	return dbPool.SQL, nil
}

func initSession() *scs.SessionManager {
	s := scs.New()
	s.Lifetime = 24 * time.Hour
	s.Cookie.Persist = true
	s.Cookie.SameSite = http.SameSiteLaxMode

	//we can use persistent storage iso cookies for session data, this allows us to
	//restart the server without users losing the login / session information
	//https://github.com/alexedwards/scs has various options available e.g.
	//session.Store = pgxstore.New(db)
	return s
}

func initDecoder() *schema.Decoder {
	return schema.NewDecoder()
}

func initTemplates() map[string]*template.Template {
	// Can later be populated as need be
	return make(map[string]*template.Template)
}

func seed(db *sql.DB) {
	fmt.Println("Start Seeding database")
	fake := faker.New()

	for i := 0; i < 200; i++ {
		stmt := `INSERT INTO customer (
                      customer_name, tel, email, address_1, address_2, address_3, post_code, created_at, updated_at)
    			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

		company := strings.Split(fake.Company().Name(), ",")[0]
		company = strings.Split(company, "-")[0]

		_, _ = db.Exec(stmt,
			company,
			fake.Phone().E164Number(),
			"info@"+strings.ReplaceAll(company, " ", "")+".com",
			fake.Address().BuildingNumber()+" "+fake.Address().StreetName(),
			fake.Address().City(),
			fake.Address().State(),
			fake.Address().PostCode(),
			time.Now(),
			time.Now())
	}

	fmt.Println("End Seeding database")
}
