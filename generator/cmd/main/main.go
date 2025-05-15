package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/agravelot/genrator/machine"
	"github.com/agravelot/genrator/process"
)

// const exportPath = ""
const exportPath = "../default/"

func main() {
	processes, err := process.GenerateProcess()
	if err != nil {
		panic(err)
	}

	for _, p := range processes {
		err = writeProcess(p)
		if err != nil {
			panic(err)
		}
	}

	// Machine
	machines, err := machine.GenerateMachines()
	if err != nil {
		panic(err)
	}

	for _, m := range machines {
		err = writeMachine(m)
		if err != nil {
			panic(err)
		}
	}
}

func writeMachine(machine machine.Machine) error {
	jsonMachine, err := json.MarshalIndent(machine, "", "	")
	if err != nil {
		return fmt.Errorf("error marshalling machine: %w", err)
	}

	err = os.WriteFile(exportPath+"machine/"+machine.Name+".json", jsonMachine, 0644)
	if err != nil {
		return fmt.Errorf("error writing machine: %w", err)
	}

	err = os.WriteFile(exportPath+"machine/"+machine.Name+".info", []byte(machine.InfoFile), 0644)
	if err != nil {
		return fmt.Errorf("error writing machine info: %w", err)
	}

	return nil
}

func writeProcess(process process.Process) error {
	jsonProcess, err := json.MarshalIndent(process, "", "	")
	if err != nil {
		return fmt.Errorf("error marshalling process: %w", err)
	}

	err = os.WriteFile(exportPath+"process/"+process.Name+".json", jsonProcess, 0644)
	if err != nil {
		return fmt.Errorf("error writing process: %w", err)
	}

	err = os.WriteFile(exportPath+"process/"+process.Name+".info", []byte(process.InfoFile), 0644)
	if err != nil {
		return fmt.Errorf("error writing process info: %w", err)
	}

	return nil
}
