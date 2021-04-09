package context

import (
	"os"
	"strings"
	"testing"
)

func TestNewProjectCtx(t *testing.T) {
	// Build settings file path
	workingDir, _ := os.Getwd()
	var path strings.Builder
	path.WriteString(workingDir)
	path.WriteString("/testdata")

	ctx := NewProjectCtx(path.String())

	if !ctx.IsInitialized() {
		msg := "Project initialization failed."
		t.Errorf(msg)
	}
}
