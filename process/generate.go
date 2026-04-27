package process

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/agravelot/genrator/utils"
)

func GenerateProcess() ([]Process, error) {
	// TODO Matrix nozzle -> height -> data
	inherits := []string{
		// 0.4
		"0.08mm Extra Fine @Voron",
		"0.12mm Fine @Voron",
		"0.15mm Optimal @Voron",
		"0.20mm Standard @Voron",
		"0.24mm Draft @Voron",
		"0.28mm Extra Draft @Voron",
		// 0.6
		"0.18mm Fine 0.6 nozzle @Voron",
		"0.24mm Optimal 0.6 nozzle @Voron",
		"0.30mm Standard 0.6 nozzle @Voron",
		"0.36mm Draft 0.6 nozzle @Voron",
		"0.42mm Extra Draft 0.6 nozzle @Voron",
		// 0.8
		"0.24mm Fine 0.8 nozzle @Voron",
		"0.40mm Standard 0.8 nozzle @Voron",
		"0.48mm Draft 0.8 nozzle @Voron",
		"0.56mm Extra Draft 0.8 nozzle @Voron",
	}

	profiles := []string{"STANDARD", "STRUCTURAL", "STRUCTURAL SILENT", "STANDARD SILENT", "STANDARD SPEED", "STRUCTURAL SPEED"}

	var process []Process

	for _, profile := range profiles {
		for _, inherit := range inherits {
			name := fmt.Sprintf("%s - %s - %s", "Gen", profile, inherit)

			outterwallSpeed := 200

			m := Process{
				From:            "User",
				Inherits:        inherit,
				IsCustomDefined: "0",
				Name:            name,
				Version:         "2.2.0.1",

				// TODO dynamic update_time ?
				InfoFile: "sync_info = update\nuser_id = \nsetting_id = \nbase_id = GP004\nupdated_time = 1703950786\n",

				LayerHeight: utils.GetLayerHeight(inherit),
				NozzleSize:  utils.GetNozzleSize(inherit),
				SkirtLoops:  "2",

				// Accels
				TravelAcceleration:              "10000",
				TopSurfaceAcceleration:          "2000",
				InnerWallAcceleration:           "5000",
				DefaultAcceleration:             "6000",
				InitialLayerAcceleration:        "5000",
				SparseInfillAcceleration:        "10000",
				InternalSolidInfillAcceleration: "5000", // to avoid curling and hit curlings
				BridgeAcceleration:              "2000",

				BrimType:                     "no_brim",
				OnlyOneWallTop:               "1",
				Resolution:                   "0.008",
				ElefantFootCompensation:      "0.1", // /!\ Can break overhang on second layer
				BottomShellThickness:         "0.5",
				TopShellThickness:            "0.7",
				SparseInfillPattern:          "grid",
				SupportInterfaceBottomLayers: "0",
				SupportInterfaceTopLayers:    "5",
				SupportAngle:                 "45",
				SupportBasePattern:           "default",

				TopSurfacePattern: "monotonicline",

				// Speeds
				TravelSpeed:              "500",
				InitialLayerSpeed:        "50",
				SkirtSpeed:               "50",
				SparseInfillSpeed:        "250",
				InitialLayerInfillSpeed:  "100",
				InternalBridgeSpeed:      "150%",
				OuterWallSpeed:           fmt.Sprintf("%d", outterwallSpeed),
				InnerWallSpeed:           "250",
				BridgeSpeed:              "50",
				Overhang14Speed:          fmt.Sprintf("%d", int(float64(outterwallSpeed)*0.8)), // Outterwall * 0.8
				Overhang24Speed:          "50",
				Overhang34Speed:          "30",
				Overhang44Speed:          "15",
				SupportSpeed:             "120",
				SupportInterfaceSpeed:    "60",
				GapInfillSpeed:           "120",
				InternalSolidInfillSpeed: "250",

				PreciseOuterWall:           "0",
				ReverseOnEven:              "0",
				InfillWallOverlap:          "15%",
				TopBottomInfillWallOverlap: "15%",

				WipeOnLoops: "1",

				AccelToDecelEnable: "0",

				// Prime
				PrimeTowerBrimWidth:     "5",
				EnablePrimeTower:        "1",
				WipeTowerNoSparseLayers: "0",
				WipeTowerConeAngle:      "15",
				WipeTowerWallType:       "cone",
				PrimeTowerWidth:         "35",
				PrimeVolume:             "15",

				PostProcess:    getPostProcess(profile),
				FilenameFormat: fmt.Sprintf("{input_filename_base}_{filament_type[initial_tool]}_{print_time}_%s.gcode", profile),
			}

			err := withInherits(&m)
			if err != nil {
				log.Println(err)
				continue
			}

			if !strings.Contains(profile, "SPEED") {
				m.SmallPerimeterSpeed = "80%"
				m.SmallPerimeterThreshold = "6.5"

				// scraf can be tricy at time
				m.SeamSlopeType = "all" // Scarf joint
				m.SeamSlopeConditional = "1"
			}

			if strings.Contains(profile, "SILENT") {
				m.OuterWallSpeed = minSpeed(m.OuterWallSpeed, silentMaxSpeed)
				m.InnerWallSpeed = minSpeed(m.InnerWallSpeed, silentMaxSpeed)
				m.TravelSpeed = minSpeed(m.TravelSpeed, silentMaxSpeed)
				m.SparseInfillSpeed = minSpeed(m.SparseInfillSpeed, silentMaxSpeed)
				m.InternalSolidInfillSpeed = minSpeed(m.InternalSolidInfillSpeed, silentMaxSpeed)
				m.TopSurfaceSpeed = minSpeed(m.TopSurfaceSpeed, silentMaxSpeed)
				m.GapInfillSpeed = minSpeed(m.GapInfillSpeed, silentMaxSpeed)

				m.TravelAcceleration = minSpeed(m.TravelAcceleration, silentMaxAccel)
				m.BridgeAcceleration = minSpeed(m.BridgeAcceleration, silentMaxAccel)
				m.DefaultAcceleration = minSpeed(m.DefaultAcceleration, silentMaxAccel)
				m.InnerWallAcceleration = minSpeed(m.InnerWallAcceleration, silentMaxAccel)
				m.OuterWallAcceleration = minSpeed(m.OuterWallAcceleration, silentMaxAccel)
				m.InitialLayerAcceleration = minSpeed(m.InitialLayerAcceleration, silentMaxAccel)
				m.SparseInfillAcceleration = minSpeed(m.SparseInfillAcceleration, silentMaxAccel)
				m.TopSurfaceAcceleration = minSpeed(m.TopSurfaceAcceleration, silentMaxAccel)
				m.InternalSolidInfillAcceleration = minSpeed(m.InternalSolidInfillAcceleration, silentMaxAccel)

				// m.DefaultJerk = minSpeed(m.DefaultJerk, silentSCV)
				// m.InfillJerk = minSpeed(m.InfillJerk, silentSCV)
				// m.InitialLayerJerk = minSpeed(m.InitialLayerJerk, silentSCV)
				// m.InnerWallJerk = minSpeed(m.InitialLayerJerk, silentSCV)
				// m.OuterWallJerk = minSpeed(m.OuterWallJerk, silentSCV)
				// m.TopSurfaceJerk = minSpeed(m.TopSurfaceJerk, silentSCV)
				// m.TravelJerk = minSpeed(m.TravelJerk, silentSCV)
			}

			if strings.Contains(profile, "STRUCTURAL") {
				m.WallLoops = fmt.Sprintf("%.0f", math.Ceil(1.6/m.NozzleSize))        // 1.6mm
				m.TopShellLayers = fmt.Sprintf("%.0f", math.Ceil(1/m.LayerHeight))    // 1mm
				m.BottomShellLayers = fmt.Sprintf("%.0f", math.Ceil(1/m.LayerHeight)) // 1mm
				m.BottomShellThickness = "1.0"
				m.TopShellThickness = "1.0"

				m.SparseInfillPattern = "gyroid"
				m.SparseInfillDensity = "40%"
				m.PreciseOuterWall = "0"
			}

			if strings.Contains(profile, "SPEED") {
				// Velocity
				m.OuterWallSpeed = "300"
				m.InnerWallSpeed = "500"
				m.TravelSpeed = "600"
				m.SparseInfillSpeed = "400"
				m.InternalSolidInfillSpeed = "400"
				m.TopSurfaceSpeed = "150"
				m.GapInfillSpeed = "200"

				// Accel
				m.DefaultAcceleration = "20000"
				m.TravelAcceleration = "20000"
				m.OuterWallAcceleration = "5000"
				m.InnerWallAcceleration = "20000"
				m.BridgeAcceleration = "40%"
				m.SparseInfillAcceleration = "100%"
				m.InternalSolidInfillAcceleration = "100%"
				m.InitialLayerAcceleration = "500"
				m.TopSurfaceAcceleration = "2000"

				m.SparseInfillPattern = "grid"

				// Jerks
				m.DefaultJerk = "0"
				m.OuterWallJerk = "9"
				m.InnerWallJerk = "9"
				m.InfillJerk = "12"
				m.TopSurfaceJerk = "9"
				m.InitialLayerJerk = "9"
				m.TravelJerk = "12"
			}

			if strings.Contains(profile, "STRUCTURAL") && strings.Contains(profile, "SPEED") {
				m.SparseInfillPattern = "3dhoneycomb"
			}

			// define on nozzle size
			if m.NozzleSize == 0.4 {
				m.RaftContactDistance = "0.15"

				// Infill anchor
				m.InfillAnchor = "2"
				m.InfillAnchorMax = "12"

				// Width
				m.OuterWallLineWidth = "0.45"
				m.LineWidth = "0.45"
				m.InitialLayerLineWidth = "0.5"
				m.SparseInfillLineWidth = "0.45"
				m.InnerWallLineWidth = "0.45"
				m.InternalSolidInfillLineWidth = "0.45"
				m.SupportLineWidth = "0.36"
				m.TopSurfaceLineWidth = "0.42"

				if m.LayerHeight <= 0.15 {
					m.TopSurfaceLineWidth = "0.4"
				}
			}

			if m.NozzleSize == 0.8 {
				m.TreeSupportBranchDiameterDoubleWall = "0"
			}

			// define on nozzle size
			if m.NozzleSize == 0.6 {
				m.RaftContactDistance = "0.25"
				// Support
				m.SupportTopZDistance = "0.22"
				m.SupportInterfaceSpacing = "0.25"

				// Infill anchor
				m.InfillAnchor = "2.5"
				m.InfillAnchorMax = "20"

				m.TreeSupportBranchDiameterDoubleWall = "5"

				// Width
				m.OuterWallLineWidth = "0.6"
				m.LineWidth = "0.68"
				m.InitialLayerLineWidth = "0.68"
				m.InnerWallLineWidth = "0.6"
				m.InternalSolidInfillLineWidth = "0.6"
				m.TopSurfaceLineWidth = "0.5"

				if strings.Contains(profile, "STANDARD") || strings.Contains(profile, "SPEED") {
					m.SparseInfillLineWidth = "0.68"
					m.SupportLineWidth = "0.6"
				} else if strings.Contains(profile, "STRUCTURAL") {
					m.SparseInfillLineWidth = "0.6"
					m.SupportLineWidth = "0.55"
				} else {
					panic("unsupported type")
				}
			}

			process = append(process, m)
		}
	}

	for i := range process {
		// TODO Error group
		process[i].TravelSpeed, _ = avoidNoisySpeeds(process[i].TravelSpeed)
		process[i].BridgeSpeed, _ = avoidNoisySpeeds(process[i].BridgeSpeed)
		process[i].InternalBridgeSpeed, _ = avoidNoisySpeeds(process[i].InternalBridgeSpeed)
		process[i].Overhang14Speed, _ = avoidNoisySpeeds(process[i].Overhang14Speed)
		process[i].Overhang24Speed, _ = avoidNoisySpeeds(process[i].Overhang24Speed)
		process[i].Overhang34Speed, _ = avoidNoisySpeeds(process[i].Overhang34Speed)
		process[i].Overhang44Speed, _ = avoidNoisySpeeds(process[i].Overhang44Speed)
		process[i].SparseInfillSpeed, _ = avoidNoisySpeeds(process[i].SparseInfillSpeed)
		process[i].InternalSolidInfillSpeed, _ = avoidNoisySpeeds(process[i].InternalSolidInfillSpeed)
		process[i].TopSurfaceSpeed, _ = avoidNoisySpeeds(process[i].TopSurfaceSpeed)
		process[i].GapInfillSpeed, _ = avoidNoisySpeeds(process[i].GapInfillSpeed)
		process[i].InitialLayerSpeed, _ = avoidNoisySpeeds(process[i].InitialLayerSpeed)
		process[i].SkirtSpeed, _ = avoidNoisySpeeds(process[i].SkirtSpeed)
		process[i].InitialLayerInfillSpeed, _ = avoidNoisySpeeds(process[i].InitialLayerInfillSpeed)
	}

	// ERS
	for _, p := range process {
		p.Name = fmt.Sprintf("ERS - %s", p.Name)
		var vf float64 = 10
		vi, err := strconv.ParseFloat(p.OuterWallSpeed, 64)
		if err != nil {
			log.Printf("OuterWallSpeed '%s' is not a number profile=%s", p.OuterWallSpeed, p.Name)
			continue
		}
		// TODO divide?
		a, err := strconv.ParseFloat("-"+p.OuterWallAcceleration, 64)
		if err != nil {
			log.Println(p.Name)
			log.Printf("OuterWallAcceleration '%s' is not a number profile=%s", p.OuterWallAcceleration, p.Name)
			continue
		}

		t := (vf - vi) / a

		// log.Printf("(%f - %f) / %f = %f", vf, vi, a, t)

		w, err := strconv.ParseFloat(p.OuterWallLineWidth, 64)
		if err != nil {
			log.Println(p.Name)
			log.Printf("OuterWallLineWidth %s is not a number", p.OuterWallLineWidth)
			continue
		}

		h := p.LayerHeight

		// ext := (w-h)*h + math.Pi*math.Pow(h/2, 2)
		initialExtrusion := utils.EllipticalExtrusionRate(w, h, vi)
		finalExtrusion := utils.EllipticalExtrusionRate(w, h, vf)
		deltaExtrusion := finalExtrusion - initialExtrusion
		extrusionRateChange := deltaExtrusion / t

		p.ExtrusionRateSmoothing = strconv.FormatFloat(math.Floor(-extrusionRateChange*0.8), 'f', 0, 64)
		// Reduce if TTC
		p.MaxVolumetricExtrusionRateSlopeSegmentLength = "0.5"

		process = append(process, p)
	}

	return process, nil
}
