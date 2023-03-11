package state

import (
	"time"
)

// State specifies all known enums data.
type State struct {
	BodyStyleNames         map[int32]string `json:"body_style_names,omitempty"`
	BodyStyleValues        map[string]int32 `json:"body_style_values,omitempty"`
	BodyStyleAliases       []Alias          `json:"body_style_aliases,omitempty"`
	BodyStyleOrderedValues []int32          `json:"body_style_ordered_values,omitempty"`

	DamageTypeNames         map[int32]string       `json:"damage_type_names,omitempty"`
	DamageTypeValues        map[string]int32       `json:"damage_type_values,omitempty"`
	DamageTypeAliases       []Alias                `json:"damage_type_aliases,omitempty"`
	DamageTypeOrderedValues []int32                `json:"damage_type_ordered_values,omitempty"`
	DamageTypeIndicatives   *DamageTypeIndicatives `json:"damage_type_indicatives,omitempty"`

	DriveTypeNames         map[int32]string `json:"drive_type_names,omitempty"`
	DriveTypeValues        map[string]int32 `json:"drive_type_values,omitempty"`
	DriveTypeAliases       []Alias          `json:"drive_type_aliases,omitempty"`
	DriveTypeOrderedValues []int32          `json:"drive_type_ordered_values,omitempty"`

	FuelTypeNames         map[int32]string     `json:"fuel_type_names,omitempty"`
	FuelTypeValues        map[string]int32     `json:"fuel_type_values,omitempty"`
	FuelTypeAliases       []Alias              `json:"fuel_type_aliases,omitempty"`
	FuelTypeOrderedValues []int32              `json:"fuel_type_ordered_values,omitempty"`
	FuelTypeIndicatives   *FuelTypeIndicatives `json:"fuel_type_indicatives,omitempty"`

	TransmissionTypeNames         map[int32]string             `json:"transmission_type_names,omitempty"`
	TransmissionTypeValues        map[string]int32             `json:"transmission_type_values,omitempty"`
	TransmissionTypeAliases       []Alias                      `json:"transmission_type_aliases,omitempty"`
	TransmissionTypeOrderedValues []int32                      `json:"transmission_type_ordered_values,omitempty"`
	TransmissionTypeIndicatives   *TransmissionTypeIndicatives `json:"transmission_type_indicatives,omitempty"`

	DamageLevelNames         map[int32]string       `json:"damage_level_names,omitempty"`
	DamageLevelValues        map[string]int32       `json:"damage_level_values,omitempty"`
	DamageLevelAliases       []Alias                `json:"damage_level_aliases,omitempty"`
	DamageLevelOrderedValues []int32                `json:"damage_level_ordered_values,omitempty"`
	DamageLevelIndicatives   DamageLevelIndicatives `json:"damage_level_indicatives,omitempty"`

	WheelPositionNames         map[int32]string          `json:"wheel_position_names,omitempty"`
	WheelPositionValues        map[string]int32          `json:"wheel_position_values,omitempty"`
	WheelPositionAliases       []Alias                   `json:"wheel_position_aliases,omitempty"`
	WheelPositionOrderedValues []int32                   `json:"wheel_position_ordered_values,omitempty"`
	WheelPositionIndicatives   *WheelPositionIndicatives `json:"wheel_position_indicatives,omitempty"`

	OdometerValueStateNames         map[int32]string `json:"odometer_value_state_names,omitempty"`
	OdometerValueStateValues        map[string]int32 `json:"odometer_value_state_values,omitempty"`
	OdometerValueStateAliases       []Alias          `json:"odometer_value_state_aliases,omitempty"`
	OdometerValueStateOrderedValues []int32          `json:"odometer_value_state_ordered_values,omitempty"`

	VehicleStateNames         map[int32]string         `json:"vehicle_state_names,omitempty"`
	VehicleStateValues        map[string]int32         `json:"vehicle_state_values,omitempty"`
	VehicleStateAliases       []Alias                  `json:"vehicle_state_aliases,omitempty"`
	VehicleStateOrderedValues []int32                  `json:"vehicle_state_ordered_values,omitempty"`
	VehicleStateIndicatives   *VehicleStateIndicatives `json:"vehicle_state_indicatives,omitempty"`

	ConditionNames         map[int32]string      `json:"condition_names,omitempty"`
	ConditionValues        map[string]int32      `json:"condition_values,omitempty"`
	ConditionAliases       []Alias               `json:"condition_aliases,omitempty"`
	ConditionOrderedValues []int32               `json:"condition_ordered_values,omitempty"`
	ConditionIndicatives   *ConditionIndicatives `json:"condition_indicatives,omitempty"`

	CountryNames         map[int32]string   `json:"country_names,omitempty"`
	CountryValues        map[string]int32   `json:"country_values,omitempty"`
	CountryAliases       []Alias            `json:"country_aliases,omitempty"`
	CountryOrderedValues []int32            `json:"country_ordered_values,omitempty"`
	CountryToPhoneCodes  map[int32][]string `json:"country_to_phone_codes,omitempty"`
	PhoneCodeToCountries map[string][]int32 `json:"phone_code_to_countries,omitempty"`
	CountryToAlpha2      map[int32]string   `json:"country_to_alpha2,omitempty"`
	Alpha2ToCountry      map[string]int32   `json:"alpha2_to_country,omitempty"`

	MakerNames         map[int32]string `json:"maker_names,omitempty"`
	MakerValues        map[string]int32 `json:"maker_values,omitempty"`
	MakerAliases       []Alias          `json:"maker_aliases,omitempty"`
	MakerOrderedValues []int32          `json:"maker_ordered_values,omitempty"`

	ModelNames         map[int32]string `json:"model_names,omitempty"`
	ModelValues        map[string]int32 `json:"model_values,omitempty"`
	ModelAliases       []Alias          `json:"model_aliases,omitempty"`
	ModelOrderedValues []int32          `json:"model_ordered_values,omitempty"`

	ClassNames         map[int32]string `json:"class_names,omitempty"`
	ClassValues        map[string]int32 `json:"class_values,omitempty"`
	ClassAliases       []Alias          `json:"class_aliases,omitempty"`
	ClassOrderedValues []int32          `json:"class_ordered_values,omitempty"`

	ColorNames         map[int32]string `json:"color_names,omitempty"`
	ColorValues        map[string]int32 `json:"color_values,omitempty"`
	ColorAliases       []Alias          `json:"color_aliases,omitempty"`
	ColorOrderedValues []int32          `json:"color_ordered_values,omitempty"`
	ColorToRGB         map[int32]RGB    `json:"color_to_rgb,omitempty"`
	ColorToHSL         map[int32]HSL    `json:"color_to_hsl,omitempty"`
	ColorToHexString   map[int32]string `json:"color_to_hex_string,omitempty"`

	OwnershipDocumentNames         map[int32]string `json:"ownership_document_names,omitempty"`
	OwnershipDocumentValues        map[string]int32 `json:"ownership_document_values,omitempty"`
	OwnershipDocumentAliases       []Alias          `json:"ownership_document_aliases,omitempty"`
	OwnershipDocumentOrderedValues []int32          `json:"ownership_document_ordered_values,omitempty"`

	OdometerUnitNames         map[int32]string `json:"odometer_unit_names,omitempty"`
	OdometerUnitValues        map[string]int32 `json:"odometer_unit_values,omitempty"`
	OdometerUnitAliases       []Alias          `json:"odometer_unit_aliases,omitempty"`
	OdometerUnitOrderedValues []int32          `json:"odometer_unit_ordered_values,omitempty"`

	ModelsToMakers  map[int32]int32   `json:"models_to_makers,omitempty"`
	MakersToModels  map[int32][]int32 `json:"makers_to_models,omitempty"`
	ClassesToMakers map[int32]int32   `json:"classes_to_makers,omitempty"`
	MakersToClasses map[int32][]int32 `json:"makers_to_classes,omitempty"`
	ModelsToClasses map[int32]int32   `json:"models_to_classes,omitempty"`
	ClassesToModels map[int32][]int32 `json:"classes_to_models,omitempty"`

	MakerModelBlacklist []MakerModelBlacklistRecord `json:"maker_model_blacklist,omitempty"`
	CommonStopWords     []string                    `json:"common_stop_words,omitempty"`
	MakerStopWords      []MakerStopWordsRecord      `json:"maker_stop_words,omitempty"`

	RegionNames         map[int32]string `json:"region_names,omitempty"`
	RegionValues        map[string]int32 `json:"region_values,omitempty"`
	RegionAliases       []Alias          `json:"region_style_aliases,omitempty"`
	RegionOrderedValues []int32          `json:"region_style_ordered_values,omitempty"`

	VehicleStateDamageTypeToDamageLevel map[int32]map[int32]int32 `json:"vehicle_state_damage_type_to_damage_level,omitempty"`

	Translations map[string]map[string]map[string]string `json:"translations,omitempty"`

	Timestamp time.Time `json:"created_at,required"`
}

type Alias struct {
	Alias string `json:"alias"`
	Value int32  `json:"value"`
}

type DamageTypeIndicatives struct {
	IsRejectedRepair       map[int32]struct{} `json:"is_rejected_repair"`
	IsElectrical           map[int32]struct{} `json:"is_electrical"`
	IsSide                 map[int32]struct{} `json:"is_side"`
	IsTheft                map[int32]struct{} `json:"is_theft"`
	IsStrip                map[int32]struct{} `json:"is_strip"`
	IsRepossession         map[int32]struct{} `json:"is_repossession"`
	IsFrame                map[int32]struct{} `json:"is_frame"`
	IsNormalWear           map[int32]struct{} `json:"is_normal_wear"`
	IsPartialRepair        map[int32]struct{} `json:"is_partial_repair"`
	IsFrontEnd             map[int32]struct{} `json:"is_front_end"`
	IsAllOver              map[int32]struct{} `json:"is_all_over"`
	IsMinorDentOrScratches map[int32]struct{} `json:"is_minor_dent_or_scratches"`
	IsLeftRear             map[int32]struct{} `json:"is_left_rear"`
	IsTransmission         map[int32]struct{} `json:"is_transmission"`
	IsVandalism            map[int32]struct{} `json:"is_vandalism"`
	IsRightFront           map[int32]struct{} `json:"is_right_front"`
	IsSuspension           map[int32]struct{} `json:"is_suspension"`
	IsRearEnd              map[int32]struct{} `json:"is_rear_end"`
	IsRightSide            map[int32]struct{} `json:"is_right_side"`
	IsFlood                map[int32]struct{} `json:"is_flood"`
	IsFreshWater           map[int32]struct{} `json:"is_fresh_water"`
	IsRollover             map[int32]struct{} `json:"is_rollover"`
	IsUndercarriage        map[int32]struct{} `json:"is_undercarriage"`
	IsRightRear            map[int32]struct{} `json:"is_right_rear"`
	IsLeftFront            map[int32]struct{} `json:"is_left_front"`
	IsInteriorBurn         map[int32]struct{} `json:"is_interior_burn"`
	IsDamageHistory        map[int32]struct{} `json:"is_damage_history"`
	IsRoof                 map[int32]struct{} `json:"is_roof"`
	IsFrontAndRear         map[int32]struct{} `json:"is_front_and_rear"`
	IsHail                 map[int32]struct{} `json:"is_hail"`
	IsMechanical           map[int32]struct{} `json:"is_mechanical"`
	IsSaltWater            map[int32]struct{} `json:"is_salt_water"`
	IsTotalBurn            map[int32]struct{} `json:"is_total_burn"`
	IsBurn                 map[int32]struct{} `json:"is_burn"`
	IsReplacedVin          map[int32]struct{} `json:"is_replaced_vin"`
	IsBiohazardOrChemical  map[int32]struct{} `json:"is_biohazard_or_chemical"`
	IsLeftAndRight         map[int32]struct{} `json:"is_left_and_right"`
	IsEngineBurn           map[int32]struct{} `json:"is_engine_burn"`
	IsLeftSide             map[int32]struct{} `json:"is_left_side"`
	IsNeedsRepair          map[int32]struct{} `json:"is_needs_repair"`
	IsStormDamage          map[int32]struct{} `json:"is_storm_damage"`
	IsNone                 map[int32]struct{} `json:"is_none"`
	IsExteriorBurn         map[int32]struct{} `json:"is_exterior_burn"`
	IsEngineDamage         map[int32]struct{} `json:"is_engine_damage"`
	IsMissingOrAlteredVin  map[int32]struct{} `json:"is_missing_or_altered_vin"`
	IsCashForClunkers      map[int32]struct{} `json:"is_cash_for_clunkers"`
}

type FuelTypeIndicatives struct {
	IsDiesel   map[int32]struct{} `json:"is_diesel"`
	IsElectric map[int32]struct{} `json:"is_electric"`
	IsHybrid   map[int32]struct{} `json:"is_hybrid"`
	IsGas      map[int32]struct{} `json:"is_gas"`
}

type WheelPositionIndicatives struct {
	IsRight map[int32]struct{} `json:"is_right"`
	IsLeft  map[int32]struct{} `json:"is_left"`
}

type TransmissionTypeIndicatives struct {
	IsManual map[int32]struct{} `json:"is_manual"`
}

type ConditionIndicatives struct {
	IsGood        map[int32]struct{}
	IsDamage      map[int32]struct{}
	IsMinorDamage map[int32]struct{}
}

type VehicleStateIndicatives struct {
	IsRunsAndDrives      map[int32]struct{}
	IsEngineStartProgram map[int32]struct{}
	IsEnhanced           map[int32]struct{}
	IsNew                map[int32]struct{}
}

type DamageLevelIndicatives map[int32]int32

// Getter is an abstract state provider.
type Getter interface {
	GetState() (*State, error)
}

// TODO: lets make black list - property of the model
type MakerModelBlacklistRecord struct {
	NormMaker      string `json:"norm_maker"`
	NormModelRegex string `json:"norm_model_regex"`
}

type MakerStopWordsRecord struct {
	Maker             int32    `json:"maker"`
	StopWords         []string `json:"stop_words"`
	ExcludeFromCommon []string `json:"exclude_from_common"`
}

// GetterFunc is a simple adapter for closures.
type GetterFunc func() (*State, error)

func (g GetterFunc) GetState() (*State, error) {
	return g()
}

// RGB is red-blue-green color representation.
type RGB struct {
	Red   uint8 `json:"red,required"`
	Green uint8 `json:"green,required"`
	Blue  uint8 `json:"blue,required"`
}

// HSL is hue-saturation-lightness color representation.
type HSL struct {
	Hue        float32 `json:"hue,required"`
	Saturation float32 `json:"saturation,required"`
	Lightness  float32 `json:"lightness,required"`
}
