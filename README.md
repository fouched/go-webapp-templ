# go-webapp-templ
A simple project using [TEMPL](https://templ.guide) as template engine.

You can use [Air](https://github.com/air-verse/air) for live reloading with this project.

## Getting started

### Air Installation

With go 1.22 or higher:

```
go install github.com/air-verse/air@latest
```

The included **.air.toml** file reloads changes to templ templates as well, so install templ if required.

```
go install github.com/a-h/templ/cmd/templ@latest
```

### Docker
```
docker-compose up
```

### Dependencies
Added for templ, session management, advanced routing, CORS, Postgres, and faking data

```
go get github.com/a-h/templ
go get github.com/alexedwards/scs/v2
go get -u github.com/go-chi/chi/v5
go get github.com/go-chi/cors
go get github.com/jackc/pgx/v5
github.com/jaswdr/faker/v2
```

### Running the project
With above installed, at the root of project run:

```
air
```

This will start the web app and auto reload any changes as they are saved.


