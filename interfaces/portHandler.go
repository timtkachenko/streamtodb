package interfaces

import (
	"streamtodb/domain/entity"
	"streamtodb/domain/service"
)

//PortHandler struct defines the dependencies that will be used
type PortHandler struct {
	service service.PortServiceInterface
}

//PortHandler constructor
func NewPortHandler(s service.PortServiceInterface) *PortHandler {
	return &PortHandler{s}
}

func (s *PortHandler) SavePort(port entity.Port) error {
	if err := port.Validate(); err != nil {
		return err
	}
	p, err := s.service.GetPort(port.Codename)
	if err != nil {
		return err
	}
	if p != nil {
		port.ID = p.ID
		_, err := s.service.UpdatePort(&port)
		if err != nil {
			return err
		}
	} else {
		_, err := s.service.SavePort(&port)
		if err != nil {
			return err
		}
	}
	return nil
}
