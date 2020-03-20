package main

type ImageName struct {
	FullName string
	Name     string
	Tag      string
}

func parseImageInfo(imgName string) ImageName {
	parts := strings.Split(imgName, ":")
	fullName := parts[0]
	tag := "latest"
	if len(parts) > 1 {
		tag = parts[1]
	}
	parts = strings.Split(fullName, "/")
	repoName := parts[0]
	if len(parts) > 1 {
		repoName = parts[1]
	}
	return ImageName{FullName: fullName, Name: repoName, Tag: tag}
}
