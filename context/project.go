//
// This file contains the implementations relevant to a Coelhino project context.
// Project contexts contain high level information about a Coelhino project.
//

package context

type projectCtx struct {
	// Specifies the project root directory
	rootDir string
}

func NewProjectCtx(projectPath string) *projectCtx {

}

// Returns whether the project is initialized.
//
// An uninitialized project context is denoted by an empty rootDir member
// variable.
func (ctx *projectCtx) IsInitialized() bool {
	return ctx.rootDir != ""
}
