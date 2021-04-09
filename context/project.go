//
// This file contains the implementations relevant to a Starfire project context.
// Project contexts contain high level information about a Starfire project.
//

package context

import (
	"errors"
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

	// The path of the test plan TOML file
	// (i.e the entrypoint to the test)
	TestPlanPath string `toml:"testplan"`
}

type projectCtx struct {
	// Specifies the project root directory
	rootDir string

	//
	isInitialized bool

	// Stores project configuration information
	config Config
}

// Returns a pointer to a projectCtx object.
//
// Arguments:
//   path - The path of the project directory
//
func NewProjectCtx(path string) *projectCtx {
	ctx := projectCtx{rootDir: path}
	ctx.isInitialized = ctx.initConfig()

	return &ctx
}

// Returns whether the project is initialized.
//
// An uninitialized project context is denoted by an empty rootDir member
// variable.
//
func (ctx *projectCtx) IsInitialized() bool {
	return ctx.isInitialized
}

//
//
func (ctx *projectCtx) initConfig() bool {
	// Build settings file path
	var filepath strings.Builder
	filepath.WriteString(ctx.rootDir)
	filepath.WriteRune('/')
	filepath.WriteString(configFileName)

	// Attempt to open file
	file, err := os.OpenFile(filepath.String(), os.O_RDWR, 0755)

	// Check if file exists.
	if errors.Is(err, os.ErrNotExist) {
		log.Printf("Could not locate `%v` file in %v\n", configFileName, ctx.rootDir)
		return false
	}
	defer file.Close()

	// Parse config file
	fileInfo, _ := file.Stat()
	bytes := make([]byte, fileInfo.Size())
	file.Read(bytes)

	err = toml.Unmarshal(bytes, &ctx.config)
	if err != nil {
		log.Fatalf("Error occurred during unmarshalling: %v", err)
		return false
	}

	return true
}
