// Package api implements a GraphQL server for Titian.
package api

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/vvanpo/titian/core"
)

func Schema() *graphql.Schema {
	return graphql.MustParseSchema()
}
