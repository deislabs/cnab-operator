package controller

import (
	"github.com/deislabs/cnab-operator/pkg/controller/bundle"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, bundle.Add)
}
