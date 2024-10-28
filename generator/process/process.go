package process

import (
	"fmt"
	"math"
	"strings"

	"github.com/agravelot/genrator/utils"
)

// Process is the struct for the process
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
	SeamSlopeType      string `json:"seam_slope_type,omitempty"`
	AccelToDecelEnable string `json:"accel_to_decel_enable,omitempty"`

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
		// TODO 0.8
	}

	types := []string{"STANDARD", "STRUCTURAL", "SPEED", "STRUCTURAL SPEED"}

	var process []Process

	for _, t := range types {
		for _, inherit := range inherits {
			nozzleSize := utils.GetNozzleSize(inherit)
			layerHeight := utils.GetLayerHeight(inherit)
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
				Version:           "2.2.0.1",

				// TODO dynamic update_time ?
				InfoFile: "sync_info = update\nuser_id = \nsetting_id = \nbase_id = GP004\nupdated_time = 1703950786\n",

				SkirtLoops:                   "2",
				TravelSpeed:                  "300",
				BrimType:                     "no_brim",
				OnlyOneWallTop:               "1",
				Resolution:                   "0.008",
				TravelAcceleration:           "10000",
				ElefantFootCompensation:      "0.1", // /!\ Can break overhang on second layer
				BottomShellThickness:         "0.5",
				TopShellThickness:            "0.7",
				SparseInfillPattern:          "grid",
				SupportInterfaceBottomLayers: "0",
				SupportInterfaceTopLayers:    "5",
				SupportAngle:                 "45", // Reduce first layer motor noise

				TopSurfacePattern: "monotonicline",

				InitialLayerSpeed:       "46",
				SkirtSpeed:              "46",
				InitialLayerInfillSpeed: "100",
				InternalBridgeSpeed:     "46", // Avoid 50 speed noise spike
				BridgeSpeed:             "100",
				Overhang14Speed:         "0",
				Overhang24Speed:         "45", // Avoid 50 speed noise spike
				Overhang34Speed:         "30",
				Overhang44Speed:         "10",

				// Scarf joint
				SeamSlopeType: "all",

				AccelToDecelEnable: "0",
			}

			if strings.Contains(t, "STRUCTURAL") {
				m.WallLoops = fmt.Sprintf("%.0f", math.Ceil(1.6/nozzleSize))        // 1.6mm
				m.TopShellLayers = fmt.Sprintf("%.0f", math.Ceil(1/layerHeight))    // 1mm
				m.BottomShellLayers = fmt.Sprintf("%.0f", math.Ceil(1/layerHeight)) // 1mm
				m.BottomShellThickness = "1.0"
				m.TopShellThickness = "1.0"

				m.SparseInfillPattern = "gyroid"
				m.SparseInfillDensity = "40%"
			}

			if strings.Contains(t, "SPEED") {
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

			if strings.Contains(t, "STRUCTURAL") && strings.Contains(t, "SPEED") {
				m.SparseInfillPattern = "3dhoneycomb"
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

				if layerHeight <= 0.15 {
					m.TopSurfaceLineWidth = "0.4"
				}
			}

			if nozzleSize == 0.8 {
				m.TreeSupportBranchDiameterDoubleWall = "0"
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
