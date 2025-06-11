package config

import (
	"database/sql"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-webapp-templ/internal/repo"
	"github.com/gorilla/schema"
	"html/template"
	"log"
)

type App struct {
	Addr          string
	DSN           string
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Session       *scs.SessionManager
	DB            *sql.DB
	TemplateCache map[string]*template.Template
	Decoder       *schema.Decoder
	Repo          Repo
}

type Repo struct {
	CustomerRepo repo.CustomerRepoInterface
}
