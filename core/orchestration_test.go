package core

import (
	"os"
	"strings"
	"testing"
)

// Tests an orchestration file was properly deserialized to an Orchestration
// struct
//
func TestOrchestrationInit(t *testing.T) {
	// Build orchestration file path
	workingDir, _ := os.Getwd()
	var builder strings.Builder
	builder.WriteString(workingDir)
	builder.WriteString("/testdata/general/source/orch.toml")
	filepath := builder.String()

	// Initialize orchestration file
	orch := Orchestration{}
	orch.init(filepath)

	// Test `description` field was properly parsed
	t.Run("`description` field", func(t *testing.T) {
		expected := "This is an orchestration file"
		actual := orch.Description
		if expected != actual {
			msg := "Description does not match.\nExpected: %v\nGot: %v"
			t.Errorf(msg, expected, actual)
		}
	})

	// Test `execution_order` field was properly parsed.
	//
	// Assuming if first and last element were parsed correctly, then the whole
	// list was parsed correctly.
	t.Run("`execution_order` field", func(t *testing.T) {
		expected := "run:transaction.toml"
		actual := orch.ExecutionOrder[0]
		if expected != actual {
			msg := "Mismatch for orch.ExecutionOrder[0].\nExpected: %v\nGot: %v"
			t.Errorf(msg, expected, actual)
		}
		len := len(orch.ExecutionOrder) - 1
		expected = "run:transactions/transaction2.toml"
		actual = orch.ExecutionOrder[len]
		if expected != actual {
			msg := "Mismatch for orch.ExecutionOrder[%v].\nExpected: %v\nGot: %v"
			t.Errorf(msg, len, expected, actual)
		}
	})

	// Test multi-line string was properly parsed in orch file.
	t.Run("multi-line string parsing", func(t *testing.T) {
		expected := "then: transaction1.toml,transaction2.toml,transaction3.toml"
		actual := orch.ExecutionOrder[2]
		if expected != actual {
			msg := "Mismatch for orch.ExecutionOrder[2].\nExpected: %v\nGot: %v"
			t.Errorf(msg, expected, actual)
		}
	})
}
