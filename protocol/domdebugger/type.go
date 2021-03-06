package domdebugger

import (
	"github.com/diiyw/cuto/protocol/runtime"
	"github.com/diiyw/cuto/protocol/dom"
)

// DOM breakpoint type.
type DOMBreakpointType string

// Object event listener.
type EventListener  struct {

	// `EventListener`'s type.
	Type	string	`json:"type"`

	// `EventListener`'s useCapture.
	UseCapture	bool	`json:"useCapture"`

	// `EventListener`'s passive flag.
	Passive	bool	`json:"passive"`

	// `EventListener`'s once flag.
	Once	bool	`json:"once"`

	// Script id of the handler code.
	ScriptId	runtime.ScriptId	`json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber	int	`json:"lineNumber"`

	// Column number in the script (0-based).
	ColumnNumber	int	`json:"columnNumber"`

	// Event handler function value.
	Handler	runtime.RemoteObject	`json:"handler,omitempty"`

	// Event original handler function value.
	OriginalHandler	runtime.RemoteObject	`json:"originalHandler,omitempty"`

	// Node the listener is added to (if any).
	BackendNodeId	dom.BackendNodeId	`json:"backendNodeId,omitempty"`
}
