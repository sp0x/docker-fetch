package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type config struct {
	Verbose bool
	Dist    string
}

func initConfig() {
	globalFlags.Usage = usage
	verbose := false
	outputDir := ""
	globalFlags.BoolVar(&verbose, "v", false, "Enable verbose output")
	globalFlags.StringVar(&outputDir, "d", outputDir, "Set the output directory. If not set, the repository is cloned in a new directory inside the current working directory.")
	args := os.Args[1:]
	err := globalFlags.Parse(args)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, fmt.Errorf("arguments error: %s", err))
		os.Exit(1)
	}
	globalConfig = &config{Verbose: verbose, Dist: outputDir}
	if !verbose {
		log.SetLevel(log.WarnLevel)
	}
}
