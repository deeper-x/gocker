package gclient

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Supplier is the client interface
type Supplier interface {
	BuildClient()
	ShowContainers() ([]string, error)
}

// Manager is the client core struct
type Manager struct {
	Cli *client.Client
}

// NewManager returns the Manager struct
func NewManager() Manager {
	return Manager{}
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

// ShowContainers is the consumer utility
func ShowContainers(m Supplier) bool {
	allC, err := m.ShowContainers()

	if err != nil {
		return false
	}

	fmt.Println(allC)

	return true
}
