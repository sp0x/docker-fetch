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

	globalFlags.Usage = usage
	verbose := false
	outputDir := ""
	globalFlags.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	globalFlags.StringVar(&outputDir, "dist", outputDir, "Set the output directory")
	args := os.Args[1:]
	err := globalFlags.Parse(args)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, fmt.Errorf("arguments error: %s", err))
		os.Exit(1)
	}
	globalConfig = &Config{Verbose: verbose, Dist: outputDir}
	if !verbose {
		log.SetLevel(log.WarnLevel)
	}
}

func main() {
	if globalFlags.NArg() == 0 {
		globalFlags.Usage()
		os.Exit(1)
	}
	imageName := globalFlags.Arg(0)
	parts := strings.Split(imageName, ":")
	imageName = parts[0]
	parts = strings.Split(imageName, "/")
	repoName := parts[0]
	if len(parts) > 1 {
		repoName = parts[1]
	}
	githubUrl := formatGithubUrl(imageName)
	outputDir := formatOutputDir(repoName, globalConfig.Dist)
	log.Infof("Cloning %v", githubUrl)
	_, err := git.PlainClone(outputDir, false, &git.CloneOptions{
		URL: githubUrl,
	})
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, fmt.Errorf("error while cloning: %s", err))
		os.Exit(1)
		return
	}
	log.Infof("Cloned %s", imageName)

}

func formatOutputDir(imageName string, baseDir string) string {
	if baseDir == "" {
		baseDir = imageName
	}
	return baseDir
}

func formatGithubUrl(imageName string) string {
	parts := strings.Split(imageName, ":")
	image := parts[0]
	image = strings.Trim(strings.TrimLeft(image, "/"), "\t \n\r")
	url := fmt.Sprintf("https://github.com/%s", image)
	return url
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	globalFlags.PrintDefaults()
}
