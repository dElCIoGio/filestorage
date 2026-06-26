package signing

import (
	"github.com/dElCIoGio/filestorage/internal/modules/signing/contracts"
)

type Module struct {
	API *contracts.API
}

func NewModule(api contracts.API) *Module {

	return &Module{&api}
}
