package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/agravelot/genrator/machine"
	"github.com/agravelot/genrator/process"
)

// const exportPath = ""
var exportPath = ""

var generatedFolder = "../user/default"

// func init() {
// 	generatedFolder = filepath.Join(".", "generated")
// 	if runtime.GOOS == "linux" {
// 		exportPath = "~/.config/OrcaSlicer"
// 		return
// 	}
//
// 	panic("please add os path")
// }

func main() {
	err := os.MkdirAll(generatedFolder, 0755)
	if err != nil {
		panic(err)
	}

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

	err = os.MkdirAll(filepath.Join(generatedFolder, "machine"), 0755)
	if err != nil {
		return fmt.Errorf("error creating machine folder: %w", err)
	}

	err = os.WriteFile(filepath.Join(generatedFolder, "machine", machine.Name+".json"), jsonMachine, 0644)
	if err != nil {
		return fmt.Errorf("error writing machine: %w", err)
	}

	err = os.WriteFile(filepath.Join(generatedFolder, "machine", machine.Name+".info"), []byte(machine.InfoFile), 0644)
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

	err = os.MkdirAll(filepath.Join(generatedFolder, "process"), 0755)
	if err != nil {
		return fmt.Errorf("error creating process folder: %w", err)
	}

	err = os.WriteFile(filepath.Join(generatedFolder, "process", process.Name+".json"), jsonProcess, 0644)
	if err != nil {
		return fmt.Errorf("error writing process: %w", err)
	}

	err = os.WriteFile(filepath.Join(generatedFolder, "process", process.Name+".info"), []byte(process.InfoFile), 0644)
	if err != nil {
		return fmt.Errorf("error writing process info: %w", err)
	}

	return nil
}
