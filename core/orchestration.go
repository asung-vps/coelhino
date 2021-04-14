package core

import (
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

// Orchestration struct contains information on the orchestration file specified
// in the `starfire.toml` configuration file.
//
type Orchestration struct {
	// A written overview of the testing plan
	Description string `toml:"description"`

	// A list of events denoting the order of execution
	ExecutionOrder []string `toml:"execution_order"`
	eventList      []Event
}

func (orch *Orchestration) init(filepath string) error {
	orchFile, err := os.OpenFile(filepath, os.O_RDWR, 0755)
	if err != nil {
		log.Printf("Error occurred opening orchestration file: %v\n", err)
		return err
	}
	defer orchFile.Close()

	// Parse orchestration file
	fileInfo, _ := orchFile.Stat()
	bytes := make([]byte, fileInfo.Size())
	orchFile.Read(bytes)
	err = toml.Unmarshal(bytes, orch)
	if err != nil {
		msg := "Error occurred while unmarshalling to Orchestration struct: %v"
		log.Fatalf(msg, err)
		return err
	}

	// Create event list

	return nil
}
