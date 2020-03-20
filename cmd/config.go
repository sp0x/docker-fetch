package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Verbose bool
	Dist    string
}

func initConfig() {
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
