package machine

import (
	"fmt"
	"strconv"

	"github.com/agravelot/genrator/utils"
)

// Machine represents a 3D printer machine
type Machine struct {
	// Mandatory
	Name            string `json:"name"`
	From            string `json:"from"`
	Inherits        string `json:"inherits"`
	Version         string `json:"version"`
	IsCustomDefined string `json:"is_custom_defined"`
	InfoFile        string `json:"-"`

	// Optional
	BeforeLayerChangeGcode string   `json:"before_layer_change_gcode,omitempty"`
	ChangeFilamentGcode    string   `json:"change_filament_gcode,omitempty"`
	LayerChangeGcode       string   `json:"layer_change_gcode,omitempty"`
	MachineEndGcode        string   `json:"machine_end_gcode,omitempty"`
	MachineMaxJerkX        []string `json:"machine_max_jerk_x,omitempty"`
	MachineMaxJerkY        []string `json:"machine_max_jerk_y,omitempty"`
	MachineMaxJerkE        []string `json:"machine_max_jerk_e,omitempty"`
	MachineMaxJerkZ        []string `json:"machine_max_jerk_z,omitempty"`
	MachineMaxSpeedE       []string `json:"machine_max_speed_e,omitempty"`
	MachineMaxSpeedZ       []string `json:"machine_max_speed_z,omitempty"`
	MachineStartGcode      string   `json:"machine_start_gcode,omitempty"`
	NozzleType             string   `json:"nozzle_type,omitempty"`
	PrintHost              string   `json:"print_host,omitempty"`
	PrintHostWebui         string   `json:"print_host_webui,omitempty"`
	PrintHostAPIKey        string   `json:"printhost_apikey,omitempty"`
	RetractLiftAbove       []string `json:"retract_lift_above,omitempty"`
	RetractionLength       []string `json:"retraction_length,omitempty"`
	Thumbnails             []string `json:"thumbnails,omitempty"`
	ZHop                   []string `json:"z_hop,omitempty"`
	ZHopTypes              []string `json:"z_hop_types,omitempty"`
	SupportMultiBedTypes   string   `json:"support_multi_bed_types,omitempty"`
	PrintableHeight        string   `json:"printable_height,omitempty"`
	FanSpeedupTime         string   `json:"fan_speedup_time,omitempty"`

	// Extruder
	RetractionMinimumTravel []string `json:"retraction_minimum_travel,omitempty"`
	WipeDistance            []string `json:"wipe_distance,omitempty"`
	RetractBeforeWipe       []string `json:"retract_before_wipe,omitempty"`
	Wipe                    []string `json:"wipe,omitempty"`
	DeretractionSpeed       []string `json:"deretraction_speed,omitempty"`
	RetractionSpeed         []string `json:"retraction_speed,omitempty"`
	RetractLiftBelow        []string `json:"retract_lift_below,omitempty"`
	ResonanceAvoidance      string   `json:"resonance_avoidance,omitempty"`
	MaxResonanceAvoidance   string   `json:"max_resonance_avoidance,omitempty"`
	MinResonanceAvoidance   string   `json:"min_resonance_avoidance,omitempty"`
	BedExcludeArea          []string `json:"bed_exclude_area,omitempty"`
	TimeCost                string   `json:"time_cost,omitempty"`
	EnableFilamentRamming   string   `json:"enable_filament_ramming,omitempty"`
	PurgeInPrimeTower       string   `json:"purge_in_prime_tower,omitempty"`
	RetractLengthToolchange []string `json:"retract_length_toolchange,omitempty"`
	ParkingPosRetraction    string   `json:"parking_pos_retraction,omitempty"`
	ExtraLoadingMove        string   `json:"extra_loading_move,omitempty"`
	CoolingTubeLength       string   `json:"cooling_tube_length,omitempty"`
	CoolingTubeRetraction   string   `json:"cooling_tube_retraction,omitempty"`
}

func GenerateMachines() ([]Machine, error) {
	inherits := []string{
		"Voron 2.4 300 0.4 nozzle",
		"Voron 2.4 300 0.6 nozzle",
		"Voron 2.4 300 0.8 nozzle",
	}

	var machines []Machine

	// TODO add variants, HF, Hardnozzle, etc
	for _, inherit := range inherits {
		name := fmt.Sprintf("%s - %s", "Gen", inherit)

		// nozzleSize := getNozzleSize(inherit)

		m := Machine{
			From:            "User",
			Inherits:        inherit,
			Name:            name,
			IsCustomDefined: "0",
			Version:         "1.9.0.2",
			// TODO dynamic update_time ?
			InfoFile: "sync_info = update\nuser_id = \nsetting_id = \nbase_id = GM001\nupdated_time = 1682282966\n",

			RetractionLength:  []string{"0.6"},
			RetractionSpeed:   []string{"60"},
			DeretractionSpeed: []string{"40"},
			ZHop:              []string{"0.2"}, // TODO Maybe *2 or = layer height
			Thumbnails:        []string{"32x32/PNG", "400x300/PNG"},
			RetractLiftAbove:  []string{"0"},
			NozzleType:        "brass",
			PrintHost:         "https://moonraker.agravelot.eu",
			PrintHostWebui:    "https://fluidd.agravelot.eu",
			PrintHostAPIKey:   utils.GetApiKeyFromEnv("VORON_API_KEY"),

			ChangeFilamentGcode:     "M600",
			SupportMultiBedTypes:    "1",
			PrintableHeight:         "255",
			RetractionMinimumTravel: []string{"1.5"},
			Wipe:                    []string{"2"},
			RetractBeforeWipe:       []string{"0%"},
			FanSpeedupTime:          "0.8",

			MachineMaxSpeedE: []string{"30", "25"},
			MachineMaxSpeedZ: []string{"20", "12"},

			// TODO Extract it in file?
			MachineStartGcode:      "SET_PRINT_STATS_INFO TOTAL_LAYER=[total_layer_count]\n\nPRINT_START EXTRUDER=[nozzle_temperature_initial_layer] BED=[bed_temperature_initial_layer_single] CHAMBER=[chamber_temperature] PRINT_MIN={first_layer_print_min[0]},{first_layer_print_min[1]} PRINT_MAX={first_layer_print_max[0]},{first_layer_print_max[1]} NOZZLE_DIAMETER={nozzle_diameter[0]}",
			MachineEndGcode:        "PRINT_END\n; total layers count = [total_layer_count]",
			BeforeLayerChangeGcode: ";BEFORE_LAYER_CHANGE\n;[layer_z]\nG92 E0\nON_LAYER_CHANGE\n",
			LayerChangeGcode:       ";AFTER_LAYER_CHANGE\n;[layer_z]\nAFTER_LAYER_CHANGE\nSET_PRINT_STATS_INFO CURRENT_LAYER={layer_num + 1}",
			MachineMaxJerkX:        []string{"20", "12"}, // 20
			MachineMaxJerkY:        []string{"20", "12"}, // 20
			MachineMaxJerkZ:        []string{"3", "0.4"},
			MachineMaxJerkE:        []string{"10", "10"},

			ResonanceAvoidance:    "1",
			MaxResonanceAvoidance: "170",
			MinResonanceAvoidance: "70",

			BedExcludeArea: []string{
				"5x300", "5x0", "0x0",
				"0x300",
				"0x290",
				"41x290",
				"41x300",
				"278x300",
				"278x290",
				"300x290",
				"300x300",
			},
			TimeCost:  "0.05",
			ZHopTypes: []string{"Slope Lift"},
		}

		printableHeight, _ := strconv.Atoi(m.PrintableHeight)
		zhop, _ := strconv.Atoi(m.ZHop[0])

		m.RetractLiftBelow = []string{strconv.Itoa(printableHeight - zhop*2)}

		machines = append(machines, m)
	}

	// Generate AFC variants
	for _, m := range machines {
		machines = append(machines, Machine{
			Name:                    m.Name + " AFC",
			Version:                 "1.9.0.2",
			From:                    "User",
			Inherits:                m.Name,
			ChangeFilamentGcode:     `T[next_extruder] PURGE_LENGTH=[flush_length]\n;FLUSH_START\n;EXTERNAL_PURGE {flush_length}\n;FLUSH_END`,
			IsCustomDefined:         "0",
			EnableFilamentRamming:   "0",
			MachineStartGcode:       m.MachineStartGcode + " TOOL={initial_tool}",
			PurgeInPrimeTower:       "1",
			RetractLengthToolchange: []string{"0"},
			ParkingPosRetraction:    "0",
			ExtraLoadingMove:        "0",
			CoolingTubeLength:       "0",
			CoolingTubeRetraction:   "0",
		})
	}

	return machines, nil
}
