package utils

import (
	"log"
	"math"
	"os"
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

// EllipticalExtrusionRate computate precise extrusion in mm^3/s
func EllipticalExtrusionRate(lineWidth float64, layerHeight float64, printSpeed float64) float64 {
	// Calculate the cross-sectional area of the oval
	semiMajorAxis := lineWidth / 2.0   // Half of the line width
	semiMinorAxis := layerHeight / 2.0 // Half of the layer height

	// Area of an ellipse: π * a * b
	crossSectionalArea := math.Pi * semiMajorAxis * semiMinorAxis

	// Calculate the volumetric extrusion rate
	return crossSectionalArea * printSpeed
}

func GetApiKeyFromEnv(s string) string {
	env := os.Getenv(s)
	if env == "" {
		log.Fatalf("Environment variable %s is not set", s)
	}
	return env
}
