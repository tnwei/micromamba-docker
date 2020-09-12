package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
    "strings"
    "flag"
)

type CondaEnv struct {
	Name         string   `yaml:"name"`
	Channels     []string `yaml:"channels"`
	Dependencies []string `yaml:"dependencies"`
	Prefix       string   `yaml:"prefix"`
}

func parseCondaEnv(fname string) CondaEnv {
    // Read from file
	yamlFile, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	// Init struct to hold parsed YAML
	env := CondaEnv{}

	// Parse from YAML
	err = yaml.Unmarshal(yamlFile, &env)
	if err != nil {
		panic(err)
	}
	return env
}

func main() {
	// Get user input for filename
    fnamePtr := flag.String("fname", "environment.yml", "Path to conda environment.yml file")
    optionsPtr := flag.String("options", "-y", "Specify options to pass to micromamba. Need to wrap in double quotes, e.g. \"-y --dry-run --verbose\"")
    flag.Parse()

    // Get YAML contents
	env := parseCondaEnv(*fnamePtr)

    // Convert array contents of dependencies and channels to str
    channels := strings.Join(env.Channels, " ")
    deps := strings.Join(env.Dependencies, " ")
    // args := strings.Join(optionsArgs, " ")

	// Piece together the install string
	// micromamba create <options> -p <envName> <packages> -c <channels>
    var installStr string
    // Note that env.Name is created as a dir in /
    // micromamba needs abs dirs at this point in time
    installStr = fmt.Sprintf("micromamba create %s -p /%s %s -c %s\n", *optionsPtr, env.Name, deps, channels)

	fmt.Print(installStr)
}
