package main

import (
	"context"

	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

type backend struct {
	*framework.Backend
}

func Backend(c *logical.BackendConfig) *backend {
	var b backend

	// TODO add implementation
	b.Backend = &framework.Backend{
		BackendType: logical.TypeCredential,
		AuthRenew:   b.pathAuthRenew,
		PathsSpecial: &logical.Paths{
			Unauthenticated: []string{"login"},
		},
		Paths: []*framework.Path{
			&framework.Path{
				Pattern: "login",
				Fields: map[string]*framework.FieldSchema{
					"password": &framework.FieldSchema{
						Type: framework.TypeString,
					},
				},
				Callbacks: map[logical.Operation]framework.OperationFunc{
					logical.UpdateOperation: b.pathAuthLogin,
				},
			},
		},
	}

	return &b
}

func (b *backend) pathAuthLogin(_ context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	// TODO
	return nil, nil
}

func (b *backend) pathAuthRenew(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	// TODO
	return nil, nil
}
