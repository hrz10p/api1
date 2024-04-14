package service

import (
	"ass1/internal/data"
)

type Service struct {
	data.ModelsI
}

func New(d data.ModelsI) *Service {
	return &Service{
		d,
	}
}

func (s *Service) GetModuleInfo(id int) (data.ModuleInfo, error) {
	return s.RetrieveModuleInfo(id)
}

func (s *Service) UpdateModuleInfo(d data.ModuleInfo) (data.ModuleInfo, error) {
	var r data.ModuleInfo
	if err := s.ModelsI.UpdateModuleInfo(d); err != nil {
		return r, err
	}
	return s.ModelsI.RetrieveModuleInfo(d.ID)
}
func (s *Service) CreateModuleInfo(d data.ModuleInfo) (data.ModuleInfo, error) {
	var r data.ModuleInfo
	id, err := s.ModelsI.InsertModuleInfo(d)
	if err != nil {
		return r, err
	}
	return s.ModelsI.RetrieveModuleInfo(id)
}

func (s *Service) DeleteModuleInfo(id int) error {
	return s.ModelsI.DeleteModuleInfo(id)
}
