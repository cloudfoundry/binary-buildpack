package supply

import (
	"github.com/cloudfoundry/libbuildpack"
)

type Manifest interface {
}
type Stager interface {
}
type Supplier struct {
	Stager   Stager
	Manifest Manifest
	Log      *libbuildpack.Logger
}

func New(stager Stager, manifest Manifest, logger *libbuildpack.Logger) *Supplier {
	return &Supplier{
		Stager:   stager,
		Manifest: manifest,
		Log:      logger,
	}
}

func (s *Supplier) Run() error {
	return nil
}
