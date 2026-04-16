package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteProfile(baseDir string, kind string, name string, info string, profile any) error {
	jsonProfile, err := json.MarshalIndent(profile, "", "\t")
	if err != nil {
		return fmt.Errorf("error marshalling %s: %w", kind, err)
	}

	kindDir := filepath.Join(baseDir, kind)
	if err := os.MkdirAll(kindDir, 0755); err != nil {
		return fmt.Errorf("error creating %s folder: %w", kind, err)
	}

	jsonPath := filepath.Join(kindDir, name+".json")
	if err := os.WriteFile(jsonPath, jsonProfile, 0644); err != nil {
		return fmt.Errorf("error writing %s: %w", kind, err)
	}

	infoPath := filepath.Join(kindDir, name+".info")
	if err := os.WriteFile(infoPath, []byte(info), 0644); err != nil {
		return fmt.Errorf("error writing %s info: %w", kind, err)
	}

	return nil
}
