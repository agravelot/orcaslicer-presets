package prusa_profiles

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"gopkg.in/ini.v1"
)

type Result struct {
}

func Load(version string) (Result, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/prusa3d/PrusaSlicer-settings/master/live/PrusaResearch/%s.ini", version)
	res, err := http.Get(url)
	if err != nil {
		return Result{}, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return Result{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Result{}, fmt.Errorf("failed to read response body: %w", err)
	}

	inidata, err := ini.Load(bodyBytes)

	if err != nil {
		return Result{}, fmt.Errorf("failed to load ini file: %w", err)
	}

	for _, s := range inidata.Sections() {
		if !strings.Contains(s.Name(), "printer:") {
			continue
		}
		if s.Name() != "printer:Original Prusa XL 0.4 nozzle" {
			continue
		}
		_ = s.Key("inherits")
	}

	return Result{}, nil
}
