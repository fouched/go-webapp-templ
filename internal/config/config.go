package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-webapp-templ/internal/repo"
	"html/template"
	"log"
)

type App struct {
	DSN           string
	Addr          string
	Repo          Repo
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Session       *scs.SessionManager
	TemplateCache map[string]*template.Template
}

type Repo struct {
	CustomerRepo repo.CustomerRepoInterface
}
