# Golang NO-ORM SQL approach. 

This example shows how we can use golang with RDBS, in this case Postgres, to generate fully type-safe idiomatic Go code from SQL. We'll try [sqlc](https://github.com/sqlc-dev/sqlc) in this demo.

> *This is not ORM solution.*


## SQLC

From: [documentation](https://docs.sqlc.dev/en/latest/index.html)

> Sqlc generates fully type-safe idiomatic Go code from SQL. Here’s how it works:
>
> - You write SQL queries
> - You run sqlc to generate Go code that presents type-safe interfaces to those queries
> - You write application code that calls the methods sqlc generated

### First steps:
- Create a schema file (DDL) at `src/services/imdb/data_store/schema.sql`. Note: We are not creating the database in this script, but only defining tables and relations. The creation of the database and roles occurs when the image is run for the first time.
- Write your queries in the folder `src/services/imdb/data_store/queries` and organize them by domain, such as Actor, Movie, etc.
- Run command:

`make generate`
- Check folder `src/services/imdb/data_store/store` to inspect generated files. 

### Project Architecture. 

#### Repository Interface Layer

To make the repository layer more granular, we have an additional layer above the generated code, which exposes interfaces organized around domains.

> Check `src/services/imdb/data_store/repos/api.go` to see general idea.

#### Data Source

The data source component is best described as a wrapper around `github.com/jackc/pgx/v5`, facilitating a centralized location for managing connection and transaction objects. This object is created only once during application initialization and is provided to all services, which then use it to access the data store.

> Check `src/services/imdb/data_source`

#### Services

We have only one service here, which describes the IMDb example database with Actors, Movies, Directors, and Awards. 

> For the purpose of this demo, services are implemented partially, serving as a showcase.

In this example, services implement usually just basic CRUD operations. Therefore, they typically pass parameters to the data store interface and handle errors.

Upon every HTTP request, the data source, services, and repositories are initialized and wired. Connections or transactions are opened, and the service layer is ready for operation.

> To check service API: `src/services/imdb/api.go`

> To check example of service implementation: `src/services/imdb/actor_service`

> To check wiring: `src/services/imdb/wire.go`

#### Http handlers. 

HTTP handlers are implemented using the core Go 1.22 net/http library. However, there is one helper component in `src/handlers/common`. The purpose of this component is to facilitate the easy provision of API options to handler methods.

*Example*:

```
var (
	getActorsHandler = common.NewHttpHandler(GetActors)
	saveActorHandler = common.NewHttpHandler(NewActor, common.WithTx(true))
)

func AddActorHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /actors", getActorsHandler.HandlerFunc)
	mux.HandleFunc("POST /actor", saveActorHandler.HandlerFunc)
}

func GetActors(api *imdb.ImdbAPI, w http.ResponseWriter, r *http.Request) error {
   ....
}

func NewActor(api *imdb.ImdbAPI, w http.ResponseWriter, r *http.Request) error {
	....
}

```

## Run example 
1. Create image. From ./dev-ops in terminal run `docker build -t imdb-db:latest -f Dockerfile .`
2. From root run `docker-compose up`
3. From root in other terminal run `make run`

- To see all actors from postman run: GET `http:8090/actors` (off course for the first time it will be empty)

- To create new actor POST `http:8090/actor` from postman, with body:

```
{
    "Name": "Ralph Fiennes",
    "Birthyear": 1962
}
```

- With returned UUID you can now try something like this: `localhost:8090/actor/e70b0c8f-c247-4bcd-b060-72412d76ffef`



> Check handlers (`src/handlers`) for currently implemented API end points. 