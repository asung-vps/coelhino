package core

import (
	"os"
	"strings"
	"testing"
)

// Tests that the config file was properly deserialized to the Config struct
//
func TestConfigInitialization(t *testing.T) {
	// Build settings file path
	workingDir, _ := os.Getwd()
	var builder strings.Builder
	builder.WriteString(workingDir)
	builder.WriteString("/testdata/config_init")
	rootDir := builder.String()

	// Create new project context
	context := Context{}
	context.initConfig(rootDir)

	// Test `name` field was properly parsed
	expectedName := "StarfireTest"
	actualName := context.config.Name
	if actualName != expectedName {
		msg := "Config name does not match.\nExpected: %v\nGot: %v"
		t.Errorf(msg, expectedName, actualName)
	}

	// Test `orchestration-file` field was properly parsed
	builder.Reset()
	builder.WriteString(context.rootDir)
	builder.WriteString("/source/orch.toml")
	expectedFilename := builder.String()
	actualFilename := context.config.OrchestrationFile
	if actualFilename != expectedFilename {
		msg := "Orchestration file name does not match.\nExpected: %v\nGot: %v"
		t.Errorf(msg, expectedFilename, actualFilename)
	}
}
