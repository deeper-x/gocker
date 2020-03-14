package gclient

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	// man := NewManager()

	// man.BuildClient()

	// man.ShowContainers()
}

// TestShowContainers test list container
func TestShowContainers(t *testing.T) {
	mockObj := NewMockManager()
	itIs := ShowContainers(&mockObj)

	log.Println("itIs:", itIs)

	if !itIs {
		t.Error("ListContainer error")
	}
}
