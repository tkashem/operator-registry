package sqlite

import (
	"io"

	"github.com/operator-framework/operator-registry/pkg/registry"
)

type Manifest struct {
	Name     string
	Packages []Package
}

type Package struct {
	Name    string
	Bundles []Bundle
}

type Bundle struct {
	ClusterServiceVersion Resource
	Resources             []Resource
}

type Resource struct {
	Name string
	data io.Reader
}

type loader struct {
	store registry.Load
}
