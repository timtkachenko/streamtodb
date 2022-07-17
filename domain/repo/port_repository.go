package repo

import (
	"streamtodb/domain/entity"
)

type PortRepository interface {
	SavePort(*entity.Port) (*entity.Port, error)
	GetPort(codename string) (*entity.Port, error)
	UpdatePort(port *entity.Port) (*entity.Port, error)
}
