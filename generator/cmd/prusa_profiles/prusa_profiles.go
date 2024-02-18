package main

import (
	"github.com/agravelot/genrator/prusa_profiles"
	"github.com/kr/pretty"
)

func main() {
	// https://raw.githubusercontent.com/prusa3d/PrusaSlicer-settings/master/live/PrusaResearch/1.12.0-alpha0.ini
	res, err := prusa_profiles.Load("1.12.0-alpha0")

	if err != nil {
		panic(err)
	}

	pretty.Log(res)
}
