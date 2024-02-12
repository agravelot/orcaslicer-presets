package generator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Process struct {
	// Mandatory
	Name              string `json:"name"`
	From              string `json:"from"`
	Inherits          string `json:"inherits"`
	PrinterSettingsID string `json:"printer_settings_id"`
	Version           string `json:"version"`
	IsCustomDefined   string `json:"is_custom_defined"`
	InfoFile          string `json:"-"`

	// Optional
	SkirtLoops              string `json:"skirt_loops"`
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

	// Infill
	InfillAnchor    string `json:"infill_anchor,omitempty"`
	InfillAnchorMax string `json:"infill_anchor_max,omitempty"`

	// Support
	SupportBasePatternSpacing     string `json:"support_base_pattern_spacing,omitempty"`
	SupportBottomInterfaceSpacing string `json:"support_bottom_interface_spacing,omitempty"`
	SupportBottomZDistance        string `json:"support_bottom_z_distance,omitempty"`
	SupportInterfaceSpacing       string `json:"support_interface_spacing,omitempty"`
	SupportTopZDistance           string `json:"support_top_z_distance,omitempty"`
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

	TopSurfacePattern string `json:"top_surface_pattern,omitempty"`
}

type Machine struct {
	// Mandatory
	Name              string `json:"name"`
	From              string `json:"from"`
	Inherits          string `json:"inherits"`
	PrinterSettingsID string `json:"printer_settings_id"`
	Version           string `json:"version"`
	IsCustomDefined   string `json:"is_custom_defined"`
	InfoFile          string `json:"-"`

	// Optional
	BeforeLayerChangeGcode string   `json:"before_layer_change_gcode,omitempty"`
	ChangeFilamentGcode    string   `json:"change_filament_gcode,omitempty"`
	LayerChangeGcode       string   `json:"layer_change_gcode,omitempty"`
	MachineEndGcode        string   `json:"machine_end_gcode,omitempty"`
	MachineMaxJerkX        []string `json:"machine_max_jerk_x,omitempty"`
	MachineMaxJerkY        []string `json:"machine_max_jerk_y,omitempty"`
	MachineMaxSpeedE       []string `json:"machine_max_speed_e,omitempty"`
	MachineMaxSpeedZ       []string `json:"machine_max_speed_z,omitempty"`
	MachineStartGcode      string   `json:"machine_start_gcode,omitempty"`
	NozzleType             string   `json:"nozzle_type,omitempty"`
	PrintHost              string   `json:"print_host,omitempty"`
	RetractLiftAbove       []string `json:"retract_lift_above,omitempty"`
	RetractionLength       []string `json:"retraction_length,omitempty"`
	Thumbnails             []string `json:"thumbnails,omitempty"`
	ZHop                   []string `json:"z_hop,omitempty"`
	ZHopTypes              []string `json:"z_hop_types,omitempty"`
	SupportMultiBedTypes   string   `json:"support_multi_bed_types,omitempty"`
	PrintableHeight        string   `json:"printable_height,omitempty"`

	// Retraction
	RetractionMinimumTravel []string `json:"retraction_minimum_travel,omitempty"`
	WipeDistance            []string `json:"wipe_distance,omitempty"`
	RetractBeforeWipe       []string `json:"retract_before_wipe,omitempty"`
	Wipe                    []string `json:"wipe,omitempty"`
}

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
		// TODO 0.8
	}

	types := []string{"STANDARD", "STRUCTURAL", "SPEED", "STRUCTURAL SPEED"}

	var process []Process

	for _, t := range types {
		for _, inherit := range inherits {
			nozzleSize := getNozzleSize(inherit)
			layerHeigth := getLayerHeight(inherit)
			//if nozzleSize != 0.4 {
			//	nozzleName = fmt.Sprintf("%.2f nozzle", nozzleSize)
			//}
			name := fmt.Sprintf("%s - %s - %s", "Gen", t, inherit)

			m := Process{
				From:              "User",
				Inherits:          inherit,
				IsCustomDefined:   "0",
				Name:              name,
				PrinterSettingsID: name,
				Version:           "1.9.0.2",

				// TODO dynamic update_time ?
				InfoFile: "sync_info = update\nuser_id = \nsetting_id = \nbase_id = GP004\nupdated_time = 1703950786\n",

				SkirtLoops:         "2",
				TravelSpeed:        "500",
				BrimType:           "no_brim",
				OnlyOneWallTop:     "1",
				Resolution:         "0.008",
				TravelAcceleration: "15000",
				// TODO Yes ? No ?
				ElefantFootCompensation: "0", // Will break overhang on 2 layer
				BottomShellThickness:    "0.5",
				TopShellThickness:       "0.7",
				SparseInfillPattern:     "grid",

				TopSurfacePattern: "monotonicline",

				InitialLayerSpeed:       "85",
				InitialLayerInfillSpeed: "110",
			}

			if strings.Contains(t, "STRUCTURAL") {
				m.WallLoops = fmt.Sprintf("%.0f", math.Ceil(1.6/nozzleSize))        // 1.6mm
				m.TopShellLayers = fmt.Sprintf("%.0f", math.Ceil(1/layerHeigth))    // 1mm
				m.BottomShellLayers = fmt.Sprintf("%.0f", math.Ceil(1/layerHeigth)) // 1mm
				m.BottomShellThickness = "1.0"
				m.TopShellThickness = "1.0"

				m.SparseInfillPattern = "gyroid"
				m.SparseInfillDensity = "40%"
			}

			if strings.Contains(t, "SPEED") {
				// Velocity
				m.OuterWallSpeed = "150"
				m.InnerWallSpeed = "250"
				m.TravelSpeed = "500"
				m.SparseInfillSpeed = "300"
				m.InternalSolidInfillSpeed = "300"
				m.TopSurfaceSpeed = "150"
				m.GapInfillSpeed = "150"

				// Accel
				m.DefaultAcceleration = "8000"
				m.TravelAcceleration = "20000"
				m.OuterWallAcceleration = "5000"
				m.InnerWallAcceleration = "8000"
				m.SparseInfillAcceleration = "8000"
				m.InternalSolidInfillAcceleration = "8000"
				m.InitialLayerAcceleration = "1000"
				m.TopSurfaceAcceleration = "5000"

				// Jerks
				m.DefaultJerk = "12"
				m.OuterWallJerk = "10"
				m.InnerWallJerk = "10"
				m.InfillJerk = "17"
				m.TopSurfaceJerk = "12"
				m.InitialLayerJerk = "12"
				m.TravelJerk = "17"
				m.SparseInfillPattern = "grid"
			}

			// define on nozzle size
			if nozzleSize == 0.4 {
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

				if layerHeigth <= 0.15 {
					m.TopSurfaceLineWidth = "0.4"
				}
			}

			// define on nozzle size
			if nozzleSize == 0.6 {
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

				if strings.Contains(t, "STANDARD") || strings.Contains(t, "SPEED") {
					m.SparseInfillLineWidth = "0.68"
					m.SupportLineWidth = "0.6"
				} else if strings.Contains(t, "STRUCTURAL") {
					m.SparseInfillLineWidth = "0.6"
					m.SupportLineWidth = "0.55"
				} else {
					panic("unsupported type")
				}
			}

			process = append(process, m)
		}
	}

	return process, nil
}

func getLayerHeight(inheritString string) float64 {
	// Extract 4 first digit of string and convert it to float32
	// 0.08
	layerHeight, err := strconv.ParseFloat(inheritString[:4], 32)
	if err != nil {
		panic(err)
	}
	return layerHeight
}

func getNozzleSize(inheritString string) float64 {

	if strings.Contains(inheritString, " 0.6 ") {
		return 0.6
	}

	if strings.Contains(inheritString, " 0.8 ") {
		return 0.8
	}

	return 0.4
}

func GenerateMachines() ([]Machine, error) {
	inherits := []string{
		"Voron 2.4 300 0.4 nozzle",
		"Voron 2.4 300 0.6 nozzle",
		"Voron 2.4 300 0.8 nozzle",
	}

	var machines []Machine

	for _, inherit := range inherits {
		name := fmt.Sprintf("%s - %s", "Gen", inherit)

		nozzleSize := getNozzleSize(inherit)

		m := Machine{
			From:              "User",
			Inherits:          inherit,
			Name:              name,
			IsCustomDefined:   "0",
			Version:           "1.9.0.2",
			PrinterSettingsID: name,
			// TODO dynamic update_time ?
			InfoFile: "sync_info = update\nuser_id = \nsetting_id = \nbase_id = GM001\nupdated_time = 1682282966\n",

			RetractionLength:     []string{"0.7"},
			ZHop:                 []string{"0.2"},
			ZHopTypes:            []string{"Auto Lift"},
			Thumbnails:           []string{"32x32", "400x300"},
			RetractLiftAbove:     []string{"0"},
			NozzleType:           "brass",
			PrintHost:            "https://192.168.0.35",
			ChangeFilamentGcode:  "M600",
			SupportMultiBedTypes: "1",
			PrintableHeight:      "255",

			// TODO Extract it in file?
			MachineStartGcode:      "SET_PRINT_STATS_INFO TOTAL_LAYER=[total_layer_count]\n\nPRINT_START EXTRUDER=[nozzle_temperature_initial_layer] BED=[bed_temperature_initial_layer_single] CHAMBER=[chamber_temperature] PRINT_MIN={first_layer_print_min[0]},{first_layer_print_min[1]} PRINT_MAX={first_layer_print_max[0]},{first_layer_print_max[1]} NOZZLE_DIAMETER={nozzle_diameter[0]}",
			MachineEndGcode:        "PRINT_END\n; total layers count = [total_layer_count]",
			BeforeLayerChangeGcode: ";BEFORE_LAYER_CHANGE\n;[layer_z]\nG92 E0\nON_LAYER_CHANGE\n",
			LayerChangeGcode:       ";AFTER_LAYER_CHANGE\n;[layer_z]\nAFTER_LAYER_CHANGE\nSET_PRINT_STATS_INFO CURRENT_LAYER={layer_num + 1}",
			MachineMaxJerkX:        []string{"20", "12"}, // 20
			MachineMaxJerkY:        []string{"20", "12"}, // 20
		}

		if nozzleSize <= 0.4 {
			m.Wipe = []string{"0"}
		} else {
			m.Wipe = []string{"1"}
			m.RetractBeforeWipe = []string{"80%"}
			// TODO Add wipe distance?
		}

		machines = append(machines, m)
	}

	return machines, nil
}
