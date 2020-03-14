package gclient

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Supplier is the client interface
type Supplier interface {
	BuildClient()
	ListContainers() ([]string, error)
	PullImage(string) error
	ListImages() ([]types.ImageSummary, error)
	BuildContainer() error
}

// Manager is the client core struct
type Manager struct {
	Cli *client.Client
}

// NewManager returns the Manager struct
func NewManager() *Manager {
	return &Manager{}
}

// BuildClient create new Docker client instance
func (m *Manager) BuildClient() {
	c, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	m.Cli = c
}

// ListContainers list system containers
func (m *Manager) ListContainers() ([]string, error) {
	res := []string{}
	allConts, err := m.Cli.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		return res, err
	}

	for _, v := range allConts {
		res = append(res, v.ID)
	}

	return res, nil
}

// PullImage is the image pull wrapper
func (m *Manager) PullImage(name string) error {
	r, err := m.Cli.ImagePull(context.Background(), name, types.ImagePullOptions{})

	io.Copy(os.Stdout, r)

	return err
}

// ListImages returns slice of pulled images
func (m *Manager) ListImages() ([]types.ImageSummary, error) {
	images, err := m.Cli.ImageList(context.Background(), types.ImageListOptions{})

	if err != nil {
		return nil, err
	}

	return images, nil
}

// BuildContainer create local container from pulled image
func (m *Manager) BuildContainer() error {
	body, err := m.Cli.ContainerCreate(context.Background(), &container.Config{
		Image: "alpine",
	}, nil, nil, "")

	if err != nil {
		return err
	}

	log.Println(body)

	return nil
}
