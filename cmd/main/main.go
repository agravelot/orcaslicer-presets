package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/agravelot/genrator/internal/output"
	"github.com/agravelot/genrator/machine"
	"github.com/agravelot/genrator/process"
)

var generatedFolder = "./user/default"

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
		err = output.WriteProfile(generatedFolder, "process", p.Name, p.InfoFile, p)
		if err != nil {
			panic(err)
		}
	}

	machines, err := machine.GenerateMachines()
	if err != nil {
		panic(err)
	}

	for _, m := range machines {
		println("Writing machine:", m.Name)
		err = output.WriteProfile(generatedFolder, "machine", m.Name, m.InfoFile, m)
		if err != nil {
			panic(err)
		}
	}
}
