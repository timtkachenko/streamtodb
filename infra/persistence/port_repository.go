package persistence

import (
	"github.com/jinzhu/gorm"
	"streamtodb/domain/entity"
	"streamtodb/domain/repo"
)

type PortRepo struct {
	db *gorm.DB
}

func NewPortRepository(db *gorm.DB) *PortRepo {
	return &PortRepo{db}
}

//PortRepo implements the repository.PortRepository interface
var _ repo.PortRepository = &PortRepo{}

func (r *PortRepo) SavePort(port *entity.Port) (*entity.Port, error) {
	err := r.db.Debug().Create(&port).Error
	if err != nil {
		return nil, err
	}
	return port, nil
}
func (r *PortRepo) UpdatePort(port *entity.Port) (*entity.Port, error) {
	err := r.db.Debug().Model(&port).Updates(port).Error
	if err != nil {
		return nil, err
	}
	return port, nil
}

func (r *PortRepo) GetPort(codename string) (*entity.Port, error) {
	var port entity.Port
	err := r.db.Debug().Where("codename = ?", codename).Take(&port).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &port, nil
}
