package utils

import (
	"strconv"
	"strings"
)

func GetLayerHeight(inheritString string) float64 {
	layerHeight, err := strconv.ParseFloat(inheritString[:4], 32)
	if err != nil {
		panic(err)
	}
	return layerHeight
}

func GetNozzleSize(inheritString string) float64 {
	if strings.Contains(inheritString, " 0.6 ") {
		return 0.6
	}

	if strings.Contains(inheritString, " 0.8 ") {
		return 0.8
	}

	return 0.4
}
