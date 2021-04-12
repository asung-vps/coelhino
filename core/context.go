//
// This file contains the implementations relevant to a Starfire project context.
// Project contexts contain high level information about a Starfire project.
//

package core

import (
	"log"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
)

const (
	configFileName = "starfire.toml"
)

// Config struct holds high level information pertinent to the project.
//
// Design notes:
//  - Contain minimal information for simplicity
//  - No default values
type Config struct {
	// The name of the project
	Name string `toml:"name"`

	// The path of the orchestration TOML file
	OrchestrationFile string `toml:"orchestration-file"`
}

type Context struct {
	// Specifies the project root directory
	rootDir string

	// Stores project configuration information
	config Config
}

// Initializes the project context.
//
// A properly initialized context has a `starfire.toml` file that can be
// successfully unmarshalled.
//
func (context *Context) initConfig(rootDir string) error {
	context.rootDir = rootDir

	// Build settings file path
	var configPath strings.Builder
	configPath.WriteString(context.rootDir)
	configPath.WriteRune('/')
	configPath.WriteString(configFileName)

	// Check config file exists.
	configFile, err := os.OpenFile(configPath.String(), os.O_RDWR, 0755)
	if err != nil {
		log.Printf("Error occurred opening config file: %v\n", err)
		return err
	}
	defer configFile.Close()

	// Parse config file
	fileInfo, _ := configFile.Stat()
	bytes := make([]byte, fileInfo.Size())
	configFile.Read(bytes)
	err = toml.Unmarshal(bytes, &context.config)
	if err != nil {
		log.Fatalf("Error occurred during unmarshalling: %v", err)
		return err
	}

	// Set absolute path of orchestration file
	var orchPath strings.Builder
	orchPath.WriteString(context.rootDir)
	orchPath.WriteRune('/')
	orchPath.WriteString(context.config.OrchestrationFile)
	context.config.OrchestrationFile = orchPath.String()

	// Check orchestration file exists
	orchFile, err := os.OpenFile(orchPath.String(), os.O_RDWR, 0755)
	if err != nil {
		log.Printf("Error occurred when opening orchestration file: %v\n", err)
		return err
	}
	orchFile.Close()

	// If parsed successfully, parse orchestration file

	return nil
}
