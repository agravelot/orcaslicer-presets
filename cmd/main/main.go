package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/agravelot/genrator/filament"
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
		err = output.WriteProfile(generatedFolder, "machine", m.Name, m.InfoFile, m)
		if err != nil {
			panic(err)
		}
	}

	filaments, err := filament.GenerateFilaments()
	if err != nil {
		panic(err)
	}

	for _, f := range filaments {
		err = output.WriteProfile(generatedFolder, "filament", f.Name, f.InfoFile, f)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf(
		"Generated %d process, %d machine, %d filament profiles in %s\n",
		len(processes),
		len(machines),
		len(filaments),
		generatedFolder,
	)
}
