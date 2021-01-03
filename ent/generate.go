// +build tools

package ent

import (
	_ "github.com/facebook/ent/cmd/ent" // generating binary
)

//go:generate go run github.com/facebook/ent/cmd/ent generate ./schema
//go:generate go run github.com/facebook/ent/cmd/ent describe ./schema
