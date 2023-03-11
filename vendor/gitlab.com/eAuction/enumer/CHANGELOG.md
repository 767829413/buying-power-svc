# Changelog

## [Unreleased]
## [2.11.0]
### Added
- `FindDamageLevel` method that maps vehicle state and damage type into damage level.
- Test for enums translation and constant self-parsing.

### Fixed
- Model test data (mismatch after translation update).
- Self-parsing of enums.

### Added
## 2.10.0
- Japanese translation.

## [2.9.1]
### Fixed
- Models and classes naming

## [2.9.0]
### Added
- Extended API of makers, models and classes.

### Fixed
- Maker-models parsing of their English translations.

#### New enums
- Color
- Ownership document (UNSTABLE: data transmission was disabled)
- Odometer units

## [2.8.1]
### Fixed
- Order of condition

## [2.8.0]
### Added
- Country enum: alpha3-based constants with connections to alpha2 and phone codes.
- Porshe classes

## [2.7.6]
### Fixed
- models sorting

## [2.7.5]
### Added
- Porsche models

## [2.7.2]
### Fixed
- duplication of fiat 500 model
### Changed
- BMW series translation

## [2.7.1]
### Fixed
- Panic on concurrent normalization
## [2.6.1]

### Fixed
- Some enum aliases didn't contain alias for its own constant name.


## [2.7.0]

### Added
- JSON marshal and unmarshal method for enums.

### Fixed
- Empty string parsing for enums.

## [2.6.0]

### Added

- Russian and Georgian translations for condition enum.

### Fixed

- '...FromString' was normalizing strings before checking for their presence in '...Values' map, where keys - not normalized names of enum entries.

## [2.5.0]

### Added

- Indicatives for vehicle state enum.

## [2.4.4]

### Fixed

- Missing alias for "MINOR DENT OR SCRATCHES" damage type

## [2.4.2]
### Fixed
- Fixed lexus models

## [2.4.1]
### Added
- Missing damage type translations for ka

## [2.4.0]

## Changed

- Error "failed to find canonical form" is now typed and exported as ErrUnknownModel 

## [2.3.0]
## Added

- Indicatives for damage types;
- Indicatives for wheel positions;

## [2.2.0]
### Added
- Regions

## [2.1.0]

### Added
- Translations in state;
- Translations filter;
- Translation methods for enums;
- Function for getting maker by model.

## [2.0.3]

### Added

- Filters by fields for `/state` endpoint

## [2.0.2]

### Added

- Indicatives for `Condition` type.

### Fixed

- Invalid regex for ford f250 model (maker "FORD" with model "250" wasn't matched)
- Overriding whole path when adding `/state` segment for loading state

## 2.0.1

### Fixed

- Deployment issue

## 2.0.0
- Endpoint `/state` for loading state;
- Public method `Normalize` for normalizing string values in `pkg/normalization` package;
- `FromString` methods for enums that returns enum value by string (also checks for aliases internally); 
- `OdometerValueState` and `VehicleState` enums;
- `Find` for maker-model matching;
- Parameter `filter[enum]` for state endpoint;
- Translation fall-back to English (generated);
- Model enums may duplicate but only changes in aliases (regexes) are allowed.

### Changed

- Now state is loaded by http instead of taking it from static.
- CSVs now contain name aliases in 3rd columns
- All normalized names are now consistent (constant-cased)
- `IsKnown{{Type}}` and `{{Type}}FromString` public methods are now automatically generated;
- Models aliases are sorted in descending order.

### Removed

- `Wheel_type` as it is a duplicate of `Wheel_position`.

## 1.0.1

* fix compilation error, duh