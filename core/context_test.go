package core

import (
	"os"
	"strings"
	"testing"
)

// Tests a config file was properly deserialized
//
func TestConfigInitialization(t *testing.T) {
	// Build settings file path
	workingDir, _ := os.Getwd()
	var builder strings.Builder
	builder.WriteString(workingDir)
	builder.WriteString("/testdata/general")
	rootDir := builder.String()

	// Create new project context
	context := Context{}
	context.initConfig(rootDir)

	// Test `name` field was properly parsed
	t.Run("`name` field", func(t *testing.T) {
		expected := "StarfireTest"
		actual := context.config.Name
		if expected != actual {
			msg := "Config name does not match.\nExpected: %v\nGot: %v"
			t.Errorf(msg, expected, actual)
		}
	})

	// Test `orchestration-file` field was properly parsed
	t.Run("`orchestration_file` field", func(t *testing.T) {
		builder.Reset()
		builder.WriteString(context.rootDir)
		builder.WriteString("/source/orch.toml")
		expected := builder.String()
		actual := context.config.OrchestrationFile
		if expected != actual {
			msg := "Orchestration file name does not match.\nExpected: %v\nGot: %v"
			t.Errorf(msg, expected, actual)
		}
	})
}
