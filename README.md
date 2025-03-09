# go-webapp-templ
A simple project using [TEMPL](https://templ.guide) as template engine.

You can use [Air](https://github.com/air-verse/air) for live reloading with this project.

## Getting started

### Step 1: Air Installation

With go 1.22 or higher:
```
go install github.com/air-verse/air@latest
```

The included **.air.toml** file reloads changes to templ templates as well, so install templ if required.
```
go install github.com/a-h/templ/cmd/templ@latest
```

### Step 2: Create Docker containers
```
docker-compose up
```

### Step 3 : Run the project
With above installed, at the root of project run:
```
air
```
This will start the web app and auto reload any changes as they are saved.

### Step 4 : Optional - seeding some data
To create some data uncomment the line in main.go and run the project once.

## Dependencies
Added for templ, session management, advanced routing, CORS, Postgres, and faking data

```
go get github.com/a-h/templ
go get github.com/alexedwards/scs/v2
go get -u github.com/go-chi/chi/v5
go get github.com/go-chi/cors
go get github.com/jackc/pgx/v5
go get github.com/jaswdr/faker/v2
```


