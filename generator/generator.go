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

	// Optional
	OnlyOneWallTop         string   `json:"only_one_wall_top,omitempty"`
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

type Machine struct {
	// Mandatory
	Name              string `json:"name"`
	From              string `json:"from"`
	Inherits          string `json:"inherits"`
	PrinterSettingsID string `json:"printer_settings_id"`
	Version           string `json:"version"`
	IsCustomDefined   string `json:"is_custom_defined"`

	RetractionLength []string `json:"retraction_length,omitempty"`
}

func GenerateProcess() ([]Process, error) {
	inherits := []string{"0.20mm Standard @Voron"}

	var process []Process

	for _, inherit := range inherits {
		name := fmt.Sprintf("%s - %s", "Gen", inherit)

		m := Process{
			From:              "User",
			Inherits:          inherit,
			IsCustomDefined:   "0",
			Name:              name,
			PrinterSettingsID: name,
			Version:           "1.9.0.2",
		}

		process = append(process, m)
	}

	return process, nil
}

func GenerateMachines() ([]Machine, error) {
	inherits := []string{"Voron 2.4 300 0.4 nozzle"}

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

			RetractionLength: []string{"0.4"},
			//"nozzle_type": "brass",
			//"print_host": "https://192.168.0.35",
			//  "retract_lift_above": [
			//        "0.25"
			//    ], ??
			//   "thumbnails": [
			//        "32x32",
			//        "400x300"
			//    ],
			// "z_hop": [
			//        "0.2"
			//    ],
			//    "z_hop_types": [
			//        "Auto Lift"
			//    ]
		}

		machines = append(machines, m)
	}

	return machines, nil
}
