package prusa_profiles

import (
	"encoding/json"
	"fmt"
	"github.com/kr/pretty"
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

	//pretty.Log(res.Body)

	var test any

	err = json.NewDecoder(res.Body).Decode(&test)

	pretty.Log(res.StatusCode)

	bodyBytes, err := io.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//bodyString := string(bodyBytes)

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
		pretty.Log(s.Name())
		//for _, key := range s.Keys() {
		//	pretty.Log(key.Name())
		//}

		pretty.Log(s.Key("inherits"))
	}

	//pretty.Log(inidata.SectionStrings())
	//pretty.Log(inidata.Sections())

	return Result{}, nil
}
