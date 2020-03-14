package gclient

// ShowContainers is the consumer utility
func ShowContainers(s Supplier) ([]string, error) {
	result := []string{}
	allC, err := s.ListContainers()

	if err != nil {
		return nil, err
	}

	for _, v := range allC {
		result = append(result, v)
	}

	return result, nil
}

// DownloadImage is the main image pull consumer
func DownloadImage(s Supplier, name string) bool {
	err := s.PullImage(name)

	if err != nil {
		return false
	}

	return true
}

// ShowImagesIDs return a slice with all images
func ShowImagesIDs(s Supplier) ([]string, error) {
	IDs := []string{}

	imageList, err := s.ListImages()

	for _, v := range imageList {
		IDs = append(IDs, v.ID)
	}

	if err != nil {
		return nil, err
	}

	return IDs, nil
}

// CreateContainer return
func CreateContainer(s Supplier) error {
	err := s.BuildContainer()

	if err != nil {
		return err
	}

	return nil
}
