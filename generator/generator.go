package generator

import (
	"fmt"
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
	SkirtLoops          string `json:"skirt_loops"`
	TravelSpeed         string `json:"travel_speed"`
	BrimType            string `json:"brim_type"`
	OnlyOneWallTop      string `json:"only_one_wall_top,omitempty"`
	WallLoops           string `json:"wall_loops"`
	BottomShellLayers   string `json:"bottom_shell_layers"`
	PrintSettingsId     string `json:"print_settings_id"`
	SparseInfillDensity string `json:"sparse_infill_density"`
	SparseInfillPattern string `json:"sparse_infill_pattern"`
	TopShellLayers      string `json:"top_shell_layers"`
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

	types := []string{"STRUCTURAL", "SPEED"}

	var process []Process

	for _, t := range types {
		for _, inherit := range inherits {
			//nozzleSize := getNozzleSize(inherit)
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

				SkirtLoops:     "2",
				TravelSpeed:    "450",
				BrimType:       "no_brim",
				OnlyOneWallTop: "1",
			}

			if t == "STRUCTURAL" {
				m.WallLoops = "4"
				m.TopShellLayers = "5"
				m.BottomShellLayers = "5"
				m.SparseInfillPattern = "gyroid"
				m.SparseInfillDensity = "40%"
			}

			process = append(process, m)
		}
	}

	return process, nil
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

		m := Machine{
			From:              "User",
			Inherits:          inherit,
			Name:              name,
			IsCustomDefined:   "0",
			Version:           "1.9.0.2",
			PrinterSettingsID: name,
			// TODO dynamic update_time ?
			InfoFile: "sync_info = update\nuser_id = \nsetting_id = \nbase_id = GM001\nupdated_time = 1682282966\n",

			RetractionLength:    []string{"0.4"},
			ZHop:                []string{"0.2"},
			ZHopTypes:           []string{"Auto Lift"},
			Thumbnails:          []string{"32x32", "400x400"},
			RetractLiftAbove:    []string{"0.25"},
			NozzleType:          "brass",
			PrintHost:           "https://192.168.0.35",
			ChangeFilamentGcode: "M600",

			MachineStartGcode:      "SET_PRINT_STATS_INFO TOTAL_LAYER=[total_layer_count]\n\nPRINT_START EXTRUDER=[nozzle_temperature_initial_layer] BED=[bed_temperature_initial_layer_single] CHAMBER=[chamber_temperature] PRINT_MIN={first_layer_print_min[0]},{first_layer_print_min[1]} PRINT_MAX={first_layer_print_max[0]},{first_layer_print_max[1]} NOZZLE_DIAMETER={nozzle_diameter[0]}",
			MachineEndGcode:        "PRINT_END\n; total layers count = [total_layer_count]",
			BeforeLayerChangeGcode: ";BEFORE_LAYER_CHANGE\n;[layer_z]\nG92 E0\nON_LAYER_CHANGE\n",
			LayerChangeGcode:       ";AFTER_LAYER_CHANGE\n;[layer_z]\nAFTER_LAYER_CHANGE\nSET_PRINT_STATS_INFO CURRENT_LAYER={layer_num + 1}",
		}

		machines = append(machines, m)
	}

	return machines, nil
}
