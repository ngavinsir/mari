package main

//go:generate sqlboiler --wipe psql

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/ngavinsir/mari/graph"
	"github.com/ngavinsir/mari/graph/generated"
)

func main() {
	db, err := sql.Open("postgres", `dbname=mari host=localhost user=postgres password=postgres`)
	dieIf(err)

	router := chi.NewRouter()

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: &graph.Resolver{
				DB: db,
			},
		}),
	)

	router.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	router.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
