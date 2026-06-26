package service

type FakeService struct {
}

func (f FakeService) ValidateURL(url string) error {
	return nil
}

func NewFakeService() *FakeService {
	return &FakeService{}
}
