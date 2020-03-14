package gclient

import (
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// MockManager mock Repo object
type MockManager struct {
	Cli *client.Client
}

// NewMockManager returns MockManager instance
func NewMockManager() *MockManager {
	return &MockManager{}
}

// BuildClient mock
func (m *MockManager) BuildClient() {
	log.Println("Build client successfull")
}

// ListContainers list system containers
func (m *MockManager) ListContainers() ([]string, error) {
	return []string{}, nil
}

// PullImage mocks the docker image pull
func (m *MockManager) PullImage(name string) error {
	log.Printf("Pulling mock image %v", name)

	return nil
}

// ListImages returns slice of pulled images
func (m *MockManager) ListImages() ([]types.ImageSummary, error) {
	images := []types.ImageSummary{}

	return images, nil
}

// BuildContainer create local container from pulled image
func (m *MockManager) BuildContainer() error {
	return nil
}
