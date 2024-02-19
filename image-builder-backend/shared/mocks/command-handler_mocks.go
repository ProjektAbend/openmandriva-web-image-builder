package mocks

type MockCommandHandler struct{}

func (c *MockCommandHandler) RunCommand(_ string, _ ...string) error {
	return nil
}
