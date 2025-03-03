package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// Path is the resolver for the path field.
func (r *queryResolver) Path(ctx context.Context, subject string, target string, maxPathLength int, usingOnly []model.Edge) ([]model.Node, error) {
	return r.Backend.Path(ctx, subject, target, maxPathLength, usingOnly)
}

// Neighbors is the resolver for the neighbors field.
func (r *queryResolver) Neighbors(ctx context.Context, node string, usingOnly []model.Edge) ([]model.Node, error) {
	return r.Backend.Neighbors(ctx, node, usingOnly)
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, node string) (model.Node, error) {
	return r.Backend.Node(ctx, node)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, nodes []string) ([]model.Node, error) {
	return r.Backend.Nodes(ctx, nodes)
}
