package gclient

import (
	"log"

	"github.com/docker/docker/client"
)

// MockManager mock Repo object
type MockManager struct {
	Cli *client.Client
}

// NewMockManager returns MockManager instance
func NewMockManager() MockManager {
	return MockManager{}
}

// BuildClient mock
func (m *MockManager) BuildClient() {
	log.Println("Build client successfull")
}

// ShowContainers list system containers
func (m *MockManager) ShowContainers() ([]string, error) {
	return []string{}, nil
}
