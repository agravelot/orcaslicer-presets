package process

import (
	"fmt"
	"strconv"
	"strings"
)

func getMode(t string) string {
	if strings.Contains(t, "SILENT") {
		return "SILENT"
	}

	if strings.Contains(t, "SPEED") {
		return "PERFORMANCE"
	}

	return "NORMAL"
}

func getPostProcess(t string) []string {
	mode := getMode(t)

	if len(mode) == 0 {
		return []string{}
	}

	return []string{
		// TODO Change path
		// fmt.Sprintf("/Users/agravelot/test.sh %s", mode),
	}
}

type NoisyRange struct {
	low  int
	high int
}

var noisyRanges = []NoisyRange{
	{
		low:  25,
		high: 40,
	},
	{
		low:  95,
		high: 175,
	},
}

func findNearest(speed int, rang NoisyRange) string {
	d1 := speed - rang.low
	d2 := rang.low - speed

	if d1 == d2 || d1 > d2 {
		return strconv.Itoa(rang.low - 1)
	}

	return strconv.Itoa(rang.high + 1)
}

// avoidNoisySpeeds take into account registered noisy speed to avoid and pick closedt match
func avoidNoisySpeeds(speed string) (string, error) {
	if strings.HasSuffix(speed, "%") {
		return speed, nil
	}

	speedInt, err := strconv.Atoi(speed)
	if err != nil {
		return speed, fmt.Errorf("speed %s is not a number", speed)
	}

	for _, r := range noisyRanges {
		if speedInt < r.low {
			break
		}
		if speedInt > r.high {
			continue
		}

		// log.Printf("speed %d noisy, find nearest", speedInt)
		return findNearest(speedInt, r), nil
	}

	return speed, nil
}

const (
	silentMaxSpeed       = "200"
	silentMaxAccel       = "8000"
	silentSCV            = "5"
	silentMinCruiseRatio = "0.4"
)

func minSpeed(a string, b string) string {
	if strings.HasSuffix(a, "%") {
		return a
	}

	if strings.HasSuffix(b, "%") {
		return b
	}

	ai, err := strconv.Atoi(a)
	if err != nil {
		return a
	}

	bi, err := strconv.Atoi(b)
	if err != nil {
		return b
	}

	if ai < bi {
		return a
	}

	return b
}

// GenerateProcess generate the process
