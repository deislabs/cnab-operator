package cnab

import (
	"github.com/deislabs/duffle/pkg/utils/crud"
)

// make sure ConfigMapStore implements the Duffle CRUD store
var _ crud.Store = &ConfigMapStore{}
