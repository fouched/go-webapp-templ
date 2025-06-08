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

### Step 2: Create Docker containers (-d in background)
```
docker-compose up -d
```

### Step 3: Run database migrations
Use pop (soda) - a nice tool that is part of the Buffalo framework that can be run on its own. 

See: https://gobuffalo.io/documentation/database/pop/

Install from https://github.com/gobuffalo/pop/releases and add to path.

Run the migration
```
soda migrate
```

### Step 4 : Run the project
With above installed, at the root of project run:
```
air
```
This will start the web app and auto reload any changes as they are saved.

### Step 5 : Optional - seeding some data
To create some data uncomment the line in main.go and run the project once.

### Step 6 : Change repository implementations
Repositories are injected via config.App which is configured when the application initializes.
The application come with a native SQL implementation, and an upper/db one. 
Update main.go to switch implementations - this can logically also done via environment based configuration if need be.  

## Dependencies
Added for templ, session management, advanced routing, CORS, Postgres, and faking data

```
go get github.com/a-h/templ
go get github.com/alexedwards/scs/v2
go get -u github.com/go-chi/chi/v5
go get github.com/go-chi/cors
go get github.com/jackc/pgx/v5
go get github.com/upper/db/v4/adapter/postgresql
go get github.com/jaswdr/faker/v2
```


