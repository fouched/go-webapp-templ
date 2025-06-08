package config

import (
	"database/sql"
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

type App struct {
	DSN           string
	Addr          string
	DB            *sql.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Session       *scs.SessionManager
	TemplateCache map[string]*template.Template
}
