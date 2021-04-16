package core

const (
	CONDITIONAL_EVENT = "conditional"
	TRANSACTION_EVENT = "transaction"
	SETTINGS_EVENT    = "settings"
)

type EventType = string

// Command keywords are used to describe what action should be taken given
// an event.
const (
	APPLY_CMD   = "apply"
	ELSE_CMD    = "else"
	ELSE_IF_CMD = "else-if"
	IF_CMD      = "if"
	IF_THEN_CMD = "if-then"
	RUN_CMD     = "run"
	THEN_CMD    = "then"
)

type Command = string
