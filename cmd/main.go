package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"strings"
)
import "flag"

var globalFlags = flag.NewFlagSet("", flag.ExitOnError)
var globalConfig *Config

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	initConfig()
}

func main() {
	if globalFlags.NArg() == 0 {
		globalFlags.Usage()
		os.Exit(1)
	}
	imgInfo := parseImageInfo(globalFlags.Arg(0))
	githubUrl := formatGithubUrl(imgInfo)
	outputDir := formatOutputDir(imgInfo, globalConfig.Dist)
	log.Infof("Cloning %v", githubUrl)
	_, err := git.PlainClone(outputDir, false, &git.CloneOptions{
		URL: githubUrl,
	})
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, fmt.Errorf("error while cloning: %s", err))
		os.Exit(1)
		return
	}
	log.Infof("Cloned %s", imgInfo.FullName)

}

func formatOutputDir(n ImageName, baseDir string) string {
	if baseDir == "" {
		baseDir = n.Name
	}
	return baseDir
}

func formatGithubUrl(n ImageName) string {
	repo := strings.Trim(strings.TrimLeft(n.FullName, "/"), "\t \n\r")
	url := fmt.Sprintf("https://github.com/%s", repo)
	return url
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	globalFlags.PrintDefaults()
}
