package service

func (s *Service) ListVMBasicView() (interface{}, error) {
	return s.dao.ListVMBasicView()
}