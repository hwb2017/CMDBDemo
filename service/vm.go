package service

func (s *Service) ListBasicView() (interface{}, error) {
	return s.dao.ListVMBasicView()
}