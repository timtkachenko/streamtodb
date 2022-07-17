package service

import (
	"streamtodb/domain/entity"
	"streamtodb/domain/repo"
)

type PortServiceInterface interface {
	SavePort(port *entity.Port) (*entity.Port, error)
	UpdatePort(port *entity.Port) (*entity.Port, error)
	GetPort(codename string) (*entity.Port, error)
}
type PortService struct {
	repo.PortRepository
}

func NewPortService(portRepo repo.PortRepository) PortServiceInterface {
	return &PortService{portRepo}
}

func (p *PortService) SavePort(port *entity.Port) (*entity.Port, error) {
	return p.PortRepository.SavePort(port)
}

func (p *PortService) UpdatePort(port *entity.Port) (*entity.Port, error) {
	return p.PortRepository.UpdatePort(port)
}

func (p *PortService) GetPort(codename string) (*entity.Port, error) {
	return p.PortRepository.GetPort(codename)
}

var _ PortServiceInterface = &PortService{}
