# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
## 2.4.0
### Added
* AND and OR operators for rules
## 2.3.0
### Added
* Admin address check
## 2.2.0

### Added
* broker platform check
* type-agnostic allow.Account

### Deprecated
* allow.Admin/PlatformID in favor of allow.Admin/Platform

## 2.1.0
### Changed
* parse jwt from cookies if `Authorization` header is empty

## 2.0.0
### Added
* `AccountPaused`, `ImportSuggestionsAllowed` and `IsImporterBlacklistManager` parameters to claims.
* `jwt` package was moved from `go-sdk`.
* `PassedGateway` request predicate.
* Platform account check.

### Changed
* IsUser function pass business accounts too

## 1.2.0
## Changed
* go-sdk is of version 1.0.0 now (instead of master branch constraint)

## 1.1.3
### Changed
* Support platform id constraint by admin account rule

## 1.1.2
### Changed
* Support custom account id in business account rule

## 1.1.1
### Changed
* eAuction/go version to 1.0.0

## 1.1.0
### Added
* Business account rule

## 1.0.0

### Added

* SkipCheck
* comfig implementation
* User address constraint
* Broker rule
* Add RuleFunc adapter
