package controller

import (
	"github.com/edge/ignition-operator/pkg/controller/ignition"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, ignition.Add)
}
