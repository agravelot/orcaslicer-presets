package main

import (
	"encoding/json"
	generator "github.com/agravelot/genrator"
	"github.com/kr/pretty"
	"os"
)

func main() {
	processes, err := generator.GenerateProcess()

	if err != nil {
		panic(err)
	}

	for _, process := range processes {
		err = writeProcess(process)
		if err != nil {
			panic(err)
		}
	}

	// Machine
	machines, err := generator.GenerateMachines()

	if err != nil {
		panic(err)
	}

	for _, machine := range machines {
		err = writeMachine(machine)

		if err != nil {
			panic(err)
		}
	}
}

func writeMachine(machine generator.Machine) error {
	jsonMachine, err := json.MarshalIndent(machine, "", "	")
	if err != nil {
		return err
	}

	pretty.Log(string(jsonMachine))

	err = os.WriteFile("machine/"+machine.Name+".json", jsonMachine, 0644)

	if err != nil {
		return err
	}

	err = os.WriteFile("machine/"+machine.Name+".info", []byte(machine.InfoFile), 0644)

	return err
}

func writeProcess(process generator.Process) error {
	jsonProcess, err := json.MarshalIndent(process, "", "	")
	if err != nil {
		return err
	}

	pretty.Log(string(jsonProcess))

	err = os.WriteFile("process/"+process.Name+".json", jsonProcess, 0644)

	if err != nil {
		return err
	}

	err = os.WriteFile("process/"+process.Name+".info", []byte(process.InfoFile), 0644)

	return err
}
