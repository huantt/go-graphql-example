package main

import (
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/huantt/go-graphql-sample/adapter"
	config "github.com/huantt/go-graphql-sample/config"
	"github.com/huantt/go-graphql-sample/graphql/resolver"
	"github.com/huantt/go-graphql-sample/graphql/schema"
	"github.com/huantt/go-graphql-sample/loader"
	"github.com/huantt/go-graphql-sample/pkg/log"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	cfg.Log.Build()

	schemaString, err := schema.String()
	if err != nil {
		log.Fatalf("reading embedded schema contents: %schemaString", err)
	}
	jsonPlaceholder := adapter.NewJsonPlaceHolder()
	rootResolver, err := resolver.NewRootResolver(jsonPlaceholder)
	if err != nil {
		log.Fatalf("creating rootResolver resolver: %schemaString", err)
	}
	graphqlSchema := graphql.MustParseSchema(schemaString, rootResolver)
	loaders := loader.Init(jsonPlaceholder)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: newHandler(graphqlSchema, loaders),
	}
	log.Infof("Listening on port: %d", cfg.Port)
	if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
