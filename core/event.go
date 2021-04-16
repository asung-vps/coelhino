package core

import (
	"errors"
	"fmt"
	"strings"
)

type Event interface {
	// Returns the command
	GetCommand() Command
	// Returns the arguments associated with this event
	GetArgs() []string
	// Returns the event type
	GetType() EventType
	// Executes the event
	Execute() error
}

func ParseToEvent(stmt string) (Event, error) {
	stmt = removeWhiteSpace(stmt)

	// Parse command
	idx := strings.Index(stmt, ":")
	command := stmt[:idx]

	// Parse arguments
	argStr := stmt[idx+1:]
	args := strings.Split(argStr, ",")

	// Event struct used is dependent on the provided command
	switch command {
	case RUN_CMD:
		return &TransactionEvent{
			command: command,
			args:    args,
		}, nil
	default:
		msg := "unknown command: %v"
		return nil, errors.Unwrap(fmt.Errorf(msg, command))
	}
}

type TransactionEvent struct {
	command Command
	args    []string
}

func (event *TransactionEvent) GetCommand() Command {
	return event.command
}

func (event *TransactionEvent) GetArgs() []string {
	return event.args
}

func (event *TransactionEvent) GetType() EventType {
	return TRANSACTION_EVENT
}

func (event *TransactionEvent) Execute() error {
	return nil
}

func removeWhiteSpace(target string) string {
	var builder strings.Builder
	builder.Grow(len(target))
	for idx, char := range target {
		switch char {
		case ' ', '\\', '\t', '\n', '\v', '\f', '\r':
			if idx > 0 && target[idx-1] == '\\' {
				builder.WriteRune(char)
			}
		default:
			builder.WriteRune(char)
		}
	}
	return builder.String()
}
