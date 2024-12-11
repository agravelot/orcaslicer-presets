package process

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

const sysDir = "/Applications/OrcaSlicer.app/Contents/Resources/profiles/Voron"

var (
	systemProcesses    = make(map[string]Process)
	systemProcessesRaw = make(map[string]map[string]any)
)

func listSystemProcesses() ([]string, error) {
	files, err := os.ReadDir(path.Join(sysDir, "/process"))
	if err != nil {
		return []string{}, fmt.Errorf("error reading process dire: %w", err)
	}

	processes := make([]string, 0)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		processes = append(processes, file.Name())
	}

	return processes, nil
}

func readSystemProcess(filename string) (Process, error) {
	file, err := os.Open(path.Join(sysDir, "process", filename+".json"))
	if err != nil {
		return Process{}, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Process{}, fmt.Errorf("error reading file: %w", err)
	}

	var p Process

	err = json.Unmarshal(byteValue, &p)
	if err != nil {
		return Process{}, fmt.Errorf("error decoding json: %w", err)
	}

	return p, nil
}

func readSystemProcessRaw(filename string) (map[string]any, error) {
	file, err := os.Open(path.Join(sysDir, "process", filename+".json"))
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	p := make(map[string]any)

	err = json.Unmarshal(byteValue, &p)
	if err != nil {
		return nil, fmt.Errorf("error decoding json: %w", err)
	}

	return p, nil
}

func getSystemProcessRaw(name string) (map[string]any, error) {
	if p, ok := systemProcessesRaw[name]; ok {
		return p, nil
	}

	p, err := readSystemProcessRaw(name)
	if err != nil {
		return nil, fmt.Errorf("error reading system process %s: %w", name, err)
	}

	systemProcessesRaw[name] = p

	return p, nil
}

func getSystemProcess(name string) (Process, error) {
	if p, ok := systemProcesses[name]; ok {
		return p, nil
	}

	p, err := readSystemProcess(name)
	if err != nil {
		return Process{}, fmt.Errorf("error reading system process %s: %w", name, err)
	}

	systemProcesses[name] = p

	return p, nil
}

// func buildSystemProcessesMap() (map[string]Process, error) {
// 	m := make(map[string]Process)
//
// 	files, err := listSystemProcesses()
// 	if err != nil {
// 		return m, fmt.Errorf("error listing system processes: %w", err)
// 	}
//
// 	for _, file := range files {
// 		p, err := readSystemProcess(file)
// 		if err != nil {
// 			return m, fmt.Errorf("error reading system process %s: %w", file, err)
// 		}
// 		m[file] = p
// 	}
//
// 	return m, nil
// }
