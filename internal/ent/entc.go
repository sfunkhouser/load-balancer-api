//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/events"

	"go.infratographer.com/load-balancer-api/x/pubsubinfo"
)

func main() {
	xExt, err := entx.NewExtension(
		entx.WithFederation(),
		entx.WithJSONScalar(),
	)
	if err != nil {
		log.Fatalf("creating entx extension: %v", err)
	}

	pubsubExt, err := pubsubinfo.NewExtension()
	if err != nil {
		log.Fatalf("creating pubsubinfo extension: %v", err)
	}

	gqlExt, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("schema/ent.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaHook(xExt.GQLSchemaHooks()...),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(
			xExt,
			gqlExt,
			pubsubExt,
		),
		entc.TemplateDir("./internal/ent/templates"),
		entc.FeatureNames("intercept"),
		entc.Dependency(
			entc.DependencyType(&events.Publisher{}),
		),
	}

	if err := entc.Generate("./internal/ent/schema", &gen.Config{
		Target:   "./internal/ent/generated",
		Package:  "go.infratographer.com/load-balancer-api/internal/ent/generated",
		Header:   entx.CopyrightHeader,
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
