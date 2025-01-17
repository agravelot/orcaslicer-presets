package process

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/agravelot/genrator/utils"
)

// Process is the struct for the process
type Process struct {
	NozzleSize  float64 `json:"-"`
	LayerHeight float64 `json:"-"`

	// Mandatory
	Name              string `json:"name"`
	From              string `json:"from"`
	Inherits          string `json:"inherits"`
	PrinterSettingsID string `json:"printer_settings_id"`
	Version           string `json:"version"`
	IsCustomDefined   string `json:"is_custom_defined"`
	InfoFile          string `json:"-"`

	// Optional
	SkirtLoops string `json:"skirt_loops"`
	SkirtSpeed string `json:"skirt_speed"`

	TravelSpeed             string `json:"travel_speed"`
	BrimType                string `json:"brim_type"`
	OnlyOneWallTop          string `json:"only_one_wall_top,omitempty"`
	WallLoops               string `json:"wall_loops,omitempty"`
	BottomShellLayers       string `json:"bottom_shell_layers,omitempty"`
	PrintSettingsId         string `json:"print_settings_id,omitempty"`
	SparseInfillDensity     string `json:"sparse_infill_density,omitempty"`
	SparseInfillPattern     string `json:"sparse_infill_pattern,omitempty"`
	TopShellLayers          string `json:"top_shell_layers,omitempty"`
	Resolution              string `json:"resolution,omitempty"`
	RaftContactDistance     string `json:"raft_contact_distance,omitempty"`
	RaftLayers              string `json:"raft_layers,omitempty"`
	ElefantFootCompensation string `json:"elefant_foot_compensation,omitempty"`
	// Used for scarf joint
	SeamSlopeType        string `json:"seam_slope_type,omitempty"`
	SeamSlopeConditional string `json:"seam_slope_conditional,omitempty"`
	AccelToDecelEnable   string `json:"accel_to_decel_enable,omitempty"`

	// Infill
	InfillAnchor    string `json:"infill_anchor,omitempty"`
	InfillAnchorMax string `json:"infill_anchor_max,omitempty"`

	// Support
	SupportBasePatternSpacing     string `json:"support_base_pattern_spacing,omitempty"`
	SupportBottomInterfaceSpacing string `json:"support_bottom_interface_spacing,omitempty"`
	SupportBottomZDistance        string `json:"support_bottom_z_distance,omitempty"`
	SupportInterfaceSpacing       string `json:"support_interface_spacing,omitempty"`
	SupportTopZDistance           string `json:"support_top_z_distance,omitempty"`
	SupportBasePattern            string `json:"support_base_pattern,omitempty"`
	// Preferred Branch Angle
	TreeSupportAngleSlow                string `json:"tree_support_angle_slow,omitempty"`
	TreeSupportBranchAngleOrganic       string `json:"tree_support_branch_angle_organic,omitempty"`
	TreeSupportBranchDiameterAngle      string `json:"tree_support_branch_diameter_angle,omitempty"`
	TreeSupportBranchDiameterDoubleWall string `json:"tree_support_branch_diameter_double_wall,omitempty"`
	TreeSupportBranchDiameterOrganic    string `json:"tree_support_branch_diameter_organic,omitempty"`
	TreeSupportBranchDistanceOrganic    string `json:"tree_support_branch_distance_organic,omitempty"`
	TreeSupportTipDiameter              string `json:"tree_support_tip_diameter,omitempty"`
	TreeSupportTopRate                  string `json:"tree_support_top_rate,omitempty"`

	// Layer width
	InitialLayerLineWidth        string `json:"initial_layer_line_width,omitempty"`
	InnerWallLineWidth           string `json:"inner_wall_line_width,omitempty"`
	InternalSolidInfillLineWidth string `json:"internal_solid_infill_line_width,omitempty"`
	LineWidth                    string `json:"line_width,omitempty"`
	OuterWallLineWidth           string `json:"outer_wall_line_width,omitempty"`
	SparseInfillLineWidth        string `json:"sparse_infill_line_width,omitempty"`
	SupportLineWidth             string `json:"support_line_width,omitempty"`
	TopSurfaceLineWidth          string `json:"top_surface_line_width,omitempty"`
	TravelAcceleration           string `json:"travel_acceleration,omitempty"`
	XyHoleCompensation           string `json:"xy_hole_compensation,omitempty"`
	BottomShellThickness         string `json:"bottom_shell_thickness,omitempty"`
	TopShellThickness            string `json:"top_shell_thickness,omitempty"`

	OuterWallAcceleration           string `json:"outer_wall_acceleration,omitempty"`
	OuterWallJerk                   string `json:"outer_wall_jerk,omitempty"`
	OuterWallSpeed                  string `json:"outer_wall_speed,omitempty"`
	InitialLayerAcceleration        string `json:"initial_layer_acceleration,omitempty"`
	InitialLayerJerk                string `json:"initial_layer_jerk,omitempty"`
	InnerWallAcceleration           string `json:"inner_wall_acceleration,omitempty"`
	InnerWallJerk                   string `json:"inner_wall_jerk,omitempty"`
	InnerWallSpeed                  string `json:"inner_wall_speed,omitempty"`
	InternalSolidInfillSpeed        string `json:"internal_solid_infill_speed,omitempty"`
	DefaultAcceleration             string `json:"default_acceleration,omitempty"`
	SparseInfillSpeed               string `json:"sparse_infill_speed,omitempty"`
	TopSurfaceAcceleration          string `json:"top_surface_acceleration,omitempty"`
	TopSurfaceJerk                  string `json:"top_surface_jerk,omitempty"`
	TopSurfaceSpeed                 string `json:"top_surface_speed,omitempty"`
	DefaultJerk                     string `json:"default_jerk,omitempty"`
	GapInfillSpeed                  string `json:"gap_infill_speed,omitempty"`
	InfillJerk                      string `json:"infill_jerk,omitempty"`
	TravelJerk                      string `json:"travel_jerk,omitempty"`
	SparseInfillAcceleration        string `json:"sparse_infill_acceleration,omitempty"`
	InternalSolidInfillAcceleration string `json:"internal_solid_infill_acceleration,omitempty"`
	InitialLayerSpeed               string `json:"initial_layer_speed,omitempty"`
	InitialLayerInfillSpeed         string `json:"initial_layer_infill_speed,omitempty"`

	TopSurfacePattern   string `json:"top_surface_pattern,omitempty"`
	BridgeSpeed         string `json:"bridge_speed,omitempty"`
	InternalBridgeSpeed string `json:"internal_bridge_speed,omitempty"`
	BridgeAcceleration  string `json:"bridge_acceleration,omitempty"`

	Overhang14Speed              string `json:"overhang_1_4_speed,omitempty"`
	Overhang24Speed              string `json:"overhang_2_4_speed,omitempty"`
	Overhang34Speed              string `json:"overhang_3_4_speed,omitempty"`
	Overhang44Speed              string `json:"overhang_4_4_speed,omitempty"`
	SupportInterfaceBottomLayers string `json:"support_interface_bottom_layers,omitempty"`
	SupportInterfaceTopLayers    string `json:"support_interface_top_layers,omitempty"`
	SupportAngle                 string `json:"support_angle,omitempty"`

	PostProcess    []string `json:"post_process,omitempty"`
	FilenameFormat string   `json:"filename_format,omitempty"`

	PreciseOuterWall  string `json:"precise_outer_wall,omitempty"`
	ReverseOnEven     string `json:"overhang_reverse,omitempty"`
	InfillWallOverlap string `json:"infill_wall_overlap,omitempty"`

	ExtrusionRateSmoothing string `json:"max_volumetric_extrusion_rate_slope,omitempty"`
}

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
		fmt.Sprintf("/Users/agravelot/test.sh %s", mode),
	}
}

type NoisyRange struct {
	low  int
	high int
}

var noisyRanges = []NoisyRange{
	{
		low:  16,
		high: 31,
	},
	{
		low:  34,
		high: 44,
	},
	{
		low:  48,
		high: 63,
	},
	{
		low:  70,
		high: 83,
	},
	{
		low:  100,
		high: 120,
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
		log.Printf("speed %s is not a number", speed)
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

func withInherits(p *Process) error {
	tree := make([]map[string]any, 0)

	fmt.Printf("%s\n", p.Name)

	tt := make(map[string]any)

	// TODO error
	q, _ := json.Marshal(p)
	_ = json.Unmarshal(q, &tt)

	tree = append(tree, tt)

	if len(p.Inherits) == 0 {
		return nil
	}

	next := p.Inherits

	for {
		parent, err := getSystemProcessRaw(next)
		if err != nil {
			return fmt.Errorf("error reading system process %s: %w", p.Inherits, err)
		}

		tree = append(tree, parent)
		v, ok := parent["inherits"].(string)
		if v == "" || !ok {
			break
		}
		next = v
	}

	// for _, e := range tree {
	// 	fmt.Printf(" -> %s", e["name"])
	// }
	// fmt.Println()

	newMap := make(map[string]any)

	for i := len(tree) - 1; i >= 0; i-- {
		for k, v := range tree[i] {
			newMap[k] = v
		}
	}

	jsonStr, _ := json.Marshal(newMap)
	err := json.Unmarshal(jsonStr, p)
	if err != nil {
		log.Println(err)
		// return nil, err
	}

	return nil
}

func minSpeed(a string, b string) string {
	if strings.HasSuffix(a, "%") {
		return a
	}

	if strings.HasSuffix(b, "%") {
		return b
	}

	ai, err := strconv.Atoi(a)
	if err != nil {
		log.Printf("speed %s is not a number", a)
		return a
	}

	bi, err := strconv.Atoi(b)
	if err != nil {
		log.Printf("speed %s is not a number", b)
		return b
	}

	if ai < bi {
		return a
	}

	return b
}

// GenerateProcess generate the process
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

			m := Process{
				From:              "User",
				Inherits:          inherit,
				IsCustomDefined:   "0",
				Name:              name,
				PrinterSettingsID: name,
				Version:           "2.2.0.1",

				// TODO dynamic update_time ?
				InfoFile: "sync_info = update\nuser_id = \nsetting_id = \nbase_id = GP004\nupdated_time = 1703950786\n",

				LayerHeight:                  utils.GetLayerHeight(inherit),
				NozzleSize:                   utils.GetNozzleSize(inherit),
				SkirtLoops:                   "2",
				TravelSpeed:                  "300",
				TravelAcceleration:           "10000",
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

				InitialLayerSpeed:       "50",
				SkirtSpeed:              "50",
				InitialLayerInfillSpeed: "100",
				InternalBridgeSpeed:     "50",
				BridgeSpeed:             "35",
				Overhang14Speed:         "0",
				Overhang24Speed:         "50",
				Overhang34Speed:         "30",
				Overhang44Speed:         "10",

				PreciseOuterWall:  "1",
				ReverseOnEven:     "1",
				InfillWallOverlap: "15%",

				AccelToDecelEnable: "0",

				PostProcess:    getPostProcess(profile),
				FilenameFormat: fmt.Sprintf("{input_filename_base}_{filament_type[initial_tool]}_{print_time}_%s.gcode", profile),
			}

			err := withInherits(&m)
			if err != nil {
				log.Println(err)
				continue
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

				m.DefaultJerk = minSpeed(m.DefaultJerk, silentSCV)
				m.InfillJerk = minSpeed(m.InfillJerk, silentSCV)
				m.InitialLayerJerk = minSpeed(m.InitialLayerJerk, silentSCV)
				m.InnerWallJerk = minSpeed(m.InitialLayerJerk, silentSCV)
				m.OuterWallJerk = minSpeed(m.OuterWallJerk, silentSCV)
				m.TopSurfaceJerk = minSpeed(m.TopSurfaceJerk, silentSCV)
				m.TravelJerk = minSpeed(m.TravelJerk, silentSCV)
			}

			if !strings.Contains(profile, "SPEED") {
				m.SeamSlopeType = "all" // Scarf joint
				m.SeamSlopeConditional = "1"
			}

			if strings.Contains(profile, "STRUCTURAL") {
				m.WallLoops = fmt.Sprintf("%.0f", math.Ceil(1.6/m.NozzleSize))        // 1.6mm
				m.TopShellLayers = fmt.Sprintf("%.0f", math.Ceil(1/m.LayerHeight))    // 1mm
				m.BottomShellLayers = fmt.Sprintf("%.0f", math.Ceil(1/m.LayerHeight)) // 1mm
				m.BottomShellThickness = "1.0"
				m.TopShellThickness = "1.0"

				m.SparseInfillPattern = "gyroid"
				m.SparseInfillDensity = "40%"
			}

			if strings.Contains(profile, "SPEED") {
				// Velocity
				m.OuterWallSpeed = "200"
				m.InnerWallSpeed = "300"
				m.TravelSpeed = "500"
				m.SparseInfillSpeed = "300"
				m.InternalSolidInfillSpeed = "300"
				m.TopSurfaceSpeed = "150"
				m.GapInfillSpeed = "300"

				// Accel
				m.DefaultAcceleration = "12000"
				m.TravelAcceleration = "15000"
				m.OuterWallAcceleration = "5000"
				m.InnerWallAcceleration = "10000"
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

		process = append(process, p)
	}

	return process, nil
}
