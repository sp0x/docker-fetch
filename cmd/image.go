package main

import "strings"

// ImageName is a struct describing an image name like "owner/repo:tag"
type ImageName struct {
	FullName string
	Name     string
	Tag      string
}

// ParseImageInfo parses a name like "owner/repo:tag"
func ParseImageInfo(imgName string) ImageName {
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
