package gclient

import (
	"testing"
)

// TestShowContainers test list container
func TestShowContainers(t *testing.T) {
	mockObj := NewMockManager()
	_, err := ShowContainers(mockObj)

	if err != nil {
		t.Errorf("ListContainer error: %v", err)
	}
}

// TestDownloadImage test the image pulling behaviour
func TestDownloadImage(t *testing.T) {
	mockMgr := NewMockManager()
	downloaded := DownloadImage(mockMgr, "dummy-name")

	if !downloaded {
		t.Errorf("DownloadImage error")
	}
}

func TestShowImages(t *testing.T) {
	mockMgr := NewMockManager()

	_, err := ShowImagesIDs(mockMgr)

	if err != nil {
		t.Errorf("Error listing images: %v", err)
	}
}

func TestBuildContainer(t *testing.T) {
	mockMgr := NewMockManager()

	err := CreateContainer(mockMgr)

	if err != nil {
		t.Errorf("Error building container %v", err)
	}
}
