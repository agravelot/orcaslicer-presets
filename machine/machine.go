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
	RetractionMinimumTravel  []string `json:"retraction_minimum_travel,omitempty"`
	WipeDistance             []string `json:"wipe_distance,omitempty"`
	TravelSlope              []string `json:"travel_slope,omitempty"`
	RetractBeforeWipe        []string `json:"retract_before_wipe,omitempty"`
	Wipe                     []string `json:"wipe,omitempty"`
	RetractWhenChangingLayer []string `json:"retract_when_changing_layer,omitempty"`
	DeretractionSpeed        []string `json:"deretraction_speed,omitempty"`
	RetractionSpeed          []string `json:"retraction_speed,omitempty"`
	RetractLiftBelow         []string `json:"retract_lift_below,omitempty"`
	ResonanceAvoidance       string   `json:"resonance_avoidance,omitempty"`
	MaxResonanceAvoidance    string   `json:"max_resonance_avoidance,omitempty"`
	MinResonanceAvoidance    string   `json:"min_resonance_avoidance,omitempty"`
	BedExcludeArea           []string `json:"bed_exclude_area,omitempty"`
	TimeCost                 string   `json:"time_cost,omitempty"`
	EnableFilamentRamming    string   `json:"enable_filament_ramming,omitempty"`
	PurgeInPrimeTower        string   `json:"purge_in_prime_tower,omitempty"`
	RetractLengthToolchange  []string `json:"retract_length_toolchange,omitempty"`
	ParkingPosRetraction     string   `json:"parking_pos_retraction,omitempty"`
	ExtraLoadingMove         string   `json:"extra_loading_move,omitempty"`
	CoolingTubeLength        string   `json:"cooling_tube_length,omitempty"`
	CoolingTubeRetraction    string   `json:"cooling_tube_retraction,omitempty"`
	// Additional Orca profile keys (stable inferred types)
	ActiveFeederMotorName                         []string `json:"active_feeder_motor_name,omitempty"`
	AdaptiveBedMeshMargin                         string   `json:"adaptive_bed_mesh_margin,omitempty"`
	ApplyTopSurfaceCompensation                   string   `json:"apply_top_surface_compensation,omitempty"`
	AutoDisableFilterOnOverheat                   string   `json:"auto_disable_filter_on_overheat,omitempty"`
	AuxiliaryFan                                  string   `json:"auxiliary_fan,omitempty"`
	BblUsePrinthost                               string   `json:"bbl_use_printhost,omitempty"`
	BedCustomModel                                string   `json:"bed_custom_model,omitempty"`
	BedCustomTexture                              string   `json:"bed_custom_texture,omitempty"`
	BedModel                                      string   `json:"bed_model,omitempty"`
	BedShape                                      string   `json:"bed_shape,omitempty"`
	BedTemperatureFormula                         string   `json:"bed_temperature_formula,omitempty"`
	BedTexture                                    string   `json:"bed_texture,omitempty"`
	BestObjectPos                                 string   `json:"best_object_pos,omitempty"`
	BottomTextureEndName                          string   `json:"bottom_texture_end_name,omitempty"`
	BottomTextureRect                             string   `json:"bottom_texture_rect,omitempty"`
	BoxId                                         string   `json:"box_id,omitempty"`
	ChangeExtrusionRoleGcode                      string   `json:"change_extrusion_role_gcode,omitempty"`
	CoolingFilterEnabled                          string   `json:"cooling_filter_enabled,omitempty"`
	CrealityFlushTime                             string   `json:"creality_flush_time,omitempty"`
	DefaultBedType                                string   `json:"default_bed_type,omitempty"`
	DefaultMaterials                              string   `json:"default_materials,omitempty"`
	DefaultNozzleVolumeType                       []string `json:"default_nozzle_volume_type,omitempty"`
	DefaultPrintProfile                           string   `json:"default_print_profile,omitempty"`
	DeretractSpeedExtruderChange                  []string `json:"deretract_speed_extruder_change,omitempty"`
	DetractionSpeed                               string   `json:"detraction_speed,omitempty"`
	DisableM73                                    string   `json:"disable_m73,omitempty"`
	EnablePowerLossRecovery                       string   `json:"enable_power_loss_recovery,omitempty"`
	EnablePreHeating                              string   `json:"enable_pre_heating,omitempty"`
	ExtruderClearanceDistToRod                    string   `json:"extruder_clearance_dist_to_rod,omitempty"`
	ExtruderClearanceMaxRadius                    string   `json:"extruder_clearance_max_radius,omitempty"`
	ExtruderMaxNozzleCount                        []string `json:"extruder_max_nozzle_count,omitempty"`
	ExtruderPrintableArea                         []string `json:"extruder_printable_area,omitempty"`
	ExtruderPrintableHeight                       []string `json:"extruder_printable_height,omitempty"`
	ExtruderType                                  []string `json:"extruder_type,omitempty"`
	ExtruderVariantList                           []string `json:"extruder_variant_list,omitempty"`
	Family                                        string   `json:"family,omitempty"`
	FanDirection                                  string   `json:"fan_direction,omitempty"`
	FanKickstart                                  string   `json:"fan_kickstart,omitempty"`
	FilamentDevAmsDryingAmsLimitations            []string `json:"filament_dev_ams_drying_ams_limitations,omitempty"`
	FilamentDevAmsDryingHeatDistortionTemperature []string `json:"filament_dev_ams_drying_heat_distortion_temperature,omitempty"`
	FilamentDevAmsDryingTemperature               []string `json:"filament_dev_ams_drying_temperature,omitempty"`
	FilamentDevAmsDryingTime                      []string `json:"filament_dev_ams_drying_time,omitempty"`
	FilamentDevChamberDryingBedTemperature        []string `json:"filament_dev_chamber_drying_bed_temperature,omitempty"`
	FilamentDevChamberDryingTime                  []string `json:"filament_dev_chamber_drying_time,omitempty"`
	FilamentDevDryingCoolingTemperature           []string `json:"filament_dev_drying_cooling_temperature,omitempty"`
	FilamentDevDryingSofteningTemperature         []string `json:"filament_dev_drying_softening_temperature,omitempty"`
	FileStartGcode                                string   `json:"file_start_gcode,omitempty"`
	GcodeFlavor                                   string   `json:"gcode_flavor,omitempty"`
	GrabLength                                    []string `json:"grab_length,omitempty"`
	GroupAlgoWithTime                             string   `json:"group_algo_with_time,omitempty"`
	HeadWrapDetectZone                            []string `json:"head_wrap_detect_zone,omitempty"`
	HostType                                      string   `json:"host_type,omitempty"`
	HotendCoolingRate                             []string `json:"hotend_cooling_rate,omitempty"`
	HotendHeatingRate                             []string `json:"hotend_heating_rate,omitempty"`
	HotendModel                                   string   `json:"hotend_model,omitempty"`
	ImageBedType                                  string   `json:"image_bed_type,omitempty"`
	Instantiation                                 string   `json:"instantiation,omitempty"`
	IsArtillery                                   string   `json:"is_artillery,omitempty"`
	IsSupport3mf                                  string   `json:"is_support_3mf,omitempty"`
	IsSupportAirCondition                         string   `json:"is_support_air_condition,omitempty"`
	IsSupportMqtt                                 string   `json:"is_support_mqtt,omitempty"`
	IsSupportMultiBox                             string   `json:"is_support_multi_box,omitempty"`
	IsSupportTimelapse                            string   `json:"is_support_timelapse,omitempty"`
	LongRetractionsWhenCut                        []string `json:"long_retractions_when_cut,omitempty"`
	MachineLedLightExist                          string   `json:"machine_LED_light_exist,omitempty"`
	MachineHotendChangeTime                       string   `json:"machine_hotend_change_time,omitempty"`
	MachineLoadFilamentTime                       string   `json:"machine_load_filament_time,omitempty"`
	MachineMaxJunctionDeviation                   []string `json:"machine_max_junction_deviation,omitempty"`
	MachinePauseGcode                             string   `json:"machine_pause_gcode,omitempty"`
	MachinePlatformMotionEnable                   string   `json:"machine_platform_motion_enable,omitempty"`
	MachinePrepareCompensationTime                string   `json:"machine_prepare_compensation_time,omitempty"`
	MachineSwitchExtruderTime                     string   `json:"machine_switch_extruder_time,omitempty"`
	MachineTech                                   string   `json:"machine_tech,omitempty"`
	MachineUnloadFilamentTime                     string   `json:"machine_unload_filament_time,omitempty"`
	ManualFilamentChange                          string   `json:"manual_filament_change,omitempty"`
	MasterExtruderId                              string   `json:"master_extruder_id,omitempty"`
	MaxResonanceAvoidanceSpeed                    string   `json:"max_resonance_avoidance_speed,omitempty"`
	MinResonanceAvoidanceSpeed                    string   `json:"min_resonance_avoidance_speed,omitempty"`
	ModelId                                       string   `json:"model_id,omitempty"`
	MultiZone                                     string   `json:"multi_zone,omitempty"`
	MultiZoneNumber                               string   `json:"multi_zone_number,omitempty"`
	NotSupportBedType                             string   `json:"not_support_bed_type,omitempty"`
	NozzleFlushDataset                            []string `json:"nozzle_flush_dataset,omitempty"`
	NozzleHeight                                  string   `json:"nozzle_height,omitempty"`
	NozzleHrc                                     string   `json:"nozzle_hrc,omitempty"`
	PauseGcode                                    string   `json:"pause_gcode,omitempty"`
	PelletModdedPrinter                           string   `json:"pellet_modded_printer,omitempty"`
	PhysicalExtruderMap                           []string `json:"physical_extruder_map,omitempty"`
	PreferredOrientation                          string   `json:"preferred_orientation,omitempty"`
	PrimeTowerPositionType                        string   `json:"prime_tower_position_type,omitempty"`
	PrinterAgent                                  string   `json:"printer_agent,omitempty"`
	PrinterExtruderId                             []string `json:"printer_extruder_id,omitempty"`
	PrinterExtruderVariant                        []string `json:"printer_extruder_variant,omitempty"`
	PrinterModel                                  string   `json:"printer_model,omitempty"`
	PrinterSettingsId                             string   `json:"printer_settings_id,omitempty"`
	PrinterStructure                              string   `json:"printer_structure,omitempty"`
	PrinterTechnology                             string   `json:"printer_technology,omitempty"`
	PrinterVariant                                string   `json:"printer_variant,omitempty"`
	PrinthostAuthorizationType                    string   `json:"printhost_authorization_type,omitempty"`
	PrinthostCafile                               string   `json:"printhost_cafile,omitempty"`
	PrinthostPassword                             string   `json:"printhost_password,omitempty"`
	PrinthostPort                                 string   `json:"printhost_port,omitempty"`
	PrinthostSslIgnoreRevoke                      string   `json:"printhost_ssl_ignore_revoke,omitempty"`
	PrinthostUser                                 string   `json:"printhost_user,omitempty"`
	PrintingByObjectGcode                         string   `json:"printing_by_object_gcode,omitempty"`
	RammingPressureAdvanceValue                   string   `json:"ramming_pressure_advance_value,omitempty"`
	RemainingTimes                                string   `json:"remaining_times,omitempty"`
	RenamedFrom                                   string   `json:"renamed_from,omitempty"`
	RetractOnTopLayer                             []string `json:"retract_on_top_layer,omitempty"`
	RetractionDistancesWhenCut                    []string `json:"retraction_distances_when_cut,omitempty"`
	RightIconOffsetBed                            string   `json:"right_icon_offset_bed,omitempty"`
	ScanFirstLayer                                string   `json:"scan_first_layer,omitempty"`
	ScanFolder                                    string   `json:"scan_folder,omitempty"`
	SettingId                                     string   `json:"setting_id,omitempty"`
	SettingsId                                    string   `json:"settings_id,omitempty"`
	SilentMode                                    string   `json:"silent_mode,omitempty"`
	SupportBoxTempControl                         string   `json:"support_box_temp_control,omitempty"`
	SupportChamberTempControl                     string   `json:"support_chamber_temp_control,omitempty"`
	SupportCoolingFilter                          string   `json:"support_cooling_filter,omitempty"`
	SupportObjectSkipFlush                        string   `json:"support_object_skip_flush,omitempty"`
	TemplateCustomGcode                           string   `json:"template_custom_gcode,omitempty"`
	ThumbnailSize                                 []string `json:"thumbnail_size,omitempty"`
	ThumbnailsFormat                              string   `json:"thumbnails_format,omitempty"`
	ThumbnailsInternal                            string   `json:"thumbnails_internal,omitempty"`
	ThumbnailsInternalSwitch                      string   `json:"thumbnails_internal_switch,omitempty"`
	TimeLapseGcode                                string   `json:"time_lapse_gcode,omitempty"`
	ToolChangeTempratureWait                      string   `json:"tool_change_temprature_wait,omitempty"`
	ToolchangeGcode                               string   `json:"toolchange_gcode,omitempty"`
	Type                                          string   `json:"type,omitempty"`
	UpwardCompatibleMachine                       []string `json:"upward_compatible_machine,omitempty"`
	Url                                           string   `json:"url,omitempty"`
	UseActivePelletFeeding                        string   `json:"use_active_pellet_feeding,omitempty"`
	UseDoubleExtruderDefaultTexture               string   `json:"use_double_extruder_default_texture,omitempty"`
	UseExtruderRotationVolume                     string   `json:"use_extruder_rotation_volume,omitempty"`
	UseFirmwareRetraction                         string   `json:"use_firmware_retraction,omitempty"`
	UseRectGrid                                   string   `json:"use_rect_grid,omitempty"`
	UseRelativeEDistances                         string   `json:"use_relative_e_distances,omitempty"`
	WipeTowerType                                 string   `json:"wipe_tower_type,omitempty"`
	WrappingDetectionGcode                        string   `json:"wrapping_detection_gcode,omitempty"`
	WrappingExcludeArea                           []string `json:"wrapping_exclude_area,omitempty"`
	ZHopWhenPrime                                 []string `json:"z_hop_when_prime,omitempty"`
	ZLiftType                                     string   `json:"z_lift_type,omitempty"`
	ZOffset                                       string   `json:"z_offset,omitempty"`

	// Additional Orca profile keys (mixed shapes, kept flexible)
	BedMeshMax                       any `json:"bed_mesh_max,omitempty"`
	BedMeshMin                       any `json:"bed_mesh_min,omitempty"`
	BedMeshProbeDistance             any `json:"bed_mesh_probe_distance,omitempty"`
	DefaultFilamentProfile           any `json:"default_filament_profile,omitempty"`
	EmitMachineLimitsToGcode         any `json:"emit_machine_limits_to_gcode,omitempty"`
	EnableLongRetractionWhenCut      any `json:"enable_long_retraction_when_cut,omitempty"`
	ExtruderClearanceHeightToLid     any `json:"extruder_clearance_height_to_lid,omitempty"`
	ExtruderClearanceHeightToRod     any `json:"extruder_clearance_height_to_rod,omitempty"`
	ExtruderClearanceRadius          any `json:"extruder_clearance_radius,omitempty"`
	ExtruderColour                   any `json:"extruder_colour,omitempty"`
	ExtruderOffset                   any `json:"extruder_offset,omitempty"`
	ExtrudersCount                   any `json:"extruders_count,omitempty"`
	FanSpeedupOverhangs              any `json:"fan_speedup_overhangs,omitempty"`
	HighCurrentOnFilamentSwap        any `json:"high_current_on_filament_swap,omitempty"`
	MachineMaxAccelerationE          any `json:"machine_max_acceleration_e,omitempty"`
	MachineMaxAccelerationExtruding  any `json:"machine_max_acceleration_extruding,omitempty"`
	MachineMaxAccelerationRetracting any `json:"machine_max_acceleration_retracting,omitempty"`
	MachineMaxAccelerationTravel     any `json:"machine_max_acceleration_travel,omitempty"`
	MachineMaxAccelerationX          any `json:"machine_max_acceleration_x,omitempty"`
	MachineMaxAccelerationY          any `json:"machine_max_acceleration_y,omitempty"`
	MachineMaxAccelerationZ          any `json:"machine_max_acceleration_z,omitempty"`
	MachineMaxSpeedX                 any `json:"machine_max_speed_x,omitempty"`
	MachineMaxSpeedY                 any `json:"machine_max_speed_y,omitempty"`
	MachineMinExtrudingRate          any `json:"machine_min_extruding_rate,omitempty"`
	MachineMinTravelRate             any `json:"machine_min_travel_rate,omitempty"`
	MachineToolChangeTime            any `json:"machine_tool_change_time,omitempty"`
	MaxLayerHeight                   any `json:"max_layer_height,omitempty"`
	MinLayerHeight                   any `json:"min_layer_height,omitempty"`
	NozzleDiameter                   any `json:"nozzle_diameter,omitempty"`
	NozzleVolume                     any `json:"nozzle_volume,omitempty"`
	PrintableArea                    any `json:"printable_area,omitempty"`
	PrinterNotes                     any `json:"printer_notes,omitempty"`
	RetractLiftEnforce               any `json:"retract_lift_enforce,omitempty"`
	RetractRestartExtra              any `json:"retract_restart_extra,omitempty"`
	RetractRestartExtraToolchange    any `json:"retract_restart_extra_toolchange,omitempty"`
	SingleExtruderMultiMaterial      any `json:"single_extruder_multi_material,omitempty"`
	SupportAirFiltration             any `json:"support_air_filtration,omitempty"`
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

		nozzleSize := utils.GetNozzleSize(inherit)

		m := Machine{
			From:            "User",
			Inherits:        inherit,
			Name:            name,
			IsCustomDefined: "0",
			Version:         "1.9.0.2",
			// TODO dynamic update_time ?
			InfoFile: "sync_info = update\nuser_id = \nsetting_id = \nbase_id = GM001\nupdated_time = 1682282966\n",

			// Retaction, zhops and wipes
			RetractionLength:         []string{fmt.Sprintf("%f", nozzleSize+0.3)},
			RetractionSpeed:          []string{"45"},
			DeretractionSpeed:        []string{"25"},
			ZHop:                     []string{"0.4"}, // low travel slope (angle) allows higher ZHop
			TravelSlope:              []string{"1"},
			RetractLiftAbove:         []string{"0"},
			RetractionMinimumTravel:  []string{"1"},
			RetractWhenChangingLayer: []string{"1"},
			Wipe:                     []string{"0"}, // disable wipe, use wipe on loops (inwards) instead
			WipeDistance:             []string{"1"},
			RetractBeforeWipe:        []string{"30%"},

			NozzleType:      "brass",
			Thumbnails:      []string{"32x32/PNG", "400x300/PNG"},
			PrintHost:       "https://moonraker.agravelot.eu",
			PrintHostWebui:  "https://fluidd.agravelot.eu",
			PrintHostAPIKey: utils.GetApiKeyFromEnv("VORON_API_KEY"),

			ChangeFilamentGcode:  "M600",
			SupportMultiBedTypes: "1",
			PrintableHeight:      "255",

			// FanSpeedupTime: "0.5",

			MachineMaxSpeedE: []string{"30", "25"},
			MachineMaxSpeedZ: []string{"20", "12"},

			// TODO Extract it in file?
			MachineStartGcode:      "SET_PRINT_STATS_INFO TOTAL_LAYER=[total_layer_count]\n\nPRINT_START EXTRUDER=[nozzle_temperature_initial_layer] BED=[bed_temperature_initial_layer_single] CHAMBER=[chamber_temperature] PRINT_MIN={first_layer_print_min[0]},{first_layer_print_min[1]} PRINT_MAX={first_layer_print_max[0]},{first_layer_print_max[1]} NOZZLE_DIAMETER={nozzle_diameter[initial_extruder]} MATERIAL=[filament_type[initial_extruder]]",
			MachineEndGcode:        "PRINT_END\n; total layers count = [total_layer_count]",
			BeforeLayerChangeGcode: ";BEFORE_LAYER_CHANGE\n;[layer_z]\nG92 E0\nON_LAYER_CHANGE\n",
			LayerChangeGcode:       ";AFTER_LAYER_CHANGE\n;[layer_z]\nAFTER_LAYER_CHANGE\nSET_PRINT_STATS_INFO CURRENT_LAYER={layer_num + 1}",
			MachineMaxJerkX:        []string{"12", "12"}, // 20
			MachineMaxJerkY:        []string{"12", "12"}, // 20
			MachineMaxJerkZ:        []string{"3", "3"},
			MachineMaxJerkE:        []string{"2.5", "2.5"},

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
		m.Name = m.Name + " AFC"
		m.ChangeFilamentGcode = `T[next_extruder] PURGE_LENGTH=[flush_length] BYPASS_AFC_BRUSH=1\n;FLUSH_START\n;EXTERNAL_PURGE {flush_length}\n;FLUSH_END`
		m.EnableFilamentRamming = "0"
		m.MachineStartGcode = m.MachineStartGcode + " TOOL={initial_tool}" // purge with GBP
		m.PurgeInPrimeTower = "0"                                          // purge with GBP
		m.RetractLengthToolchange = []string{"0"}
		m.ParkingPosRetraction = "0"
		m.ExtraLoadingMove = "0"
		m.CoolingTubeLength = "0"
		m.CoolingTubeRetraction = "0"
		machines = append(machines, m)
	}

	return machines, nil
}
