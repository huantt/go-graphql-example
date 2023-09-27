package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/huantt/go-graphql-sample/adapter"
	"github.com/huantt/go-graphql-sample/config"
	"github.com/huantt/go-graphql-sample/graphql/loader"
	"github.com/huantt/go-graphql-sample/graphql/resolver"
	"github.com/huantt/go-graphql-sample/graphql/schema"
	"github.com/huantt/go-graphql-sample/graphql/tracer"
	"github.com/huantt/go-graphql-sample/middleware"
	"github.com/huantt/go-graphql-sample/pkg/log"
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
	graphqlSchema := graphql.MustParseSchema(
		schemaString,
		rootResolver,
		graphql.MaxParallelism(cfg.Graphql.MaxParallelism),
		graphql.MaxDepth(cfg.Graphql.MaxDepth),
		graphql.Tracer(tracer.NewCustomTracer()),
	)
	loaders := loader.Init(jsonPlaceholder)

	service := gin.Default()
	service.Use(middleware.AddRequestId)
	service.Use(middleware.CORS(cfg.AllowedOrigins))
	service.POST("/graphql", middleware.NewHandler(graphqlSchema, loaders))

	if err := service.Run(fmt.Sprintf("0.0.0.0:%d", cfg.Port)); err != nil {
		log.Fatal(err)
	}
}
