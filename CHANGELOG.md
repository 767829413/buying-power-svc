# Changelog

## Unreleased

## 1.34.0

### Changed
- Deployed to the new cluster.

## 1.33.10
### Changed
- Use new payment-service connector.

### Fixed
- Ignore transactions with status `CREATED` when confirming deposits.

## 1.33.3
### Fixed
- Fixed copart client reload cron pattern.

## 1.33.2
### Changed
- Reload copart clients once a day to prevent dead sessions and spread the load on proxies more evenly.
## 1.33.0
- Handle payments namespacing.
- Allow to reset participant's state back to deposit pending for admin/broker
## 1.32.0
### Changed
- Updated copar-parser-svc connector
## 1.31.0
### Changed
- Switch to puppeteer for live bidding
## 1.30.1
### Fixed
- Do not return lot's events for lot that has been moved to another line
## 1.30.0
### Added
- Handle pre bid
## 1.29.0
### Added
- Use individual proxy for each copart account
## 1.28.0
### Added
- Save auction join details
## 1.27.0
 - Allow to configure deposit limits for autorecharger.
## 1.26.0
### Fixed
- Handle several transactions per one deposit, where only one transaction is successful
## 1.25.0
### Added
- Ensure all events have sale ID
### Changed
- Updated copart parser.
## 1.24.7
### Added
- Allow to negotiate compression for ws
## 1.24.5
### Added
- Allow to configure deposit limits for autorecharger.
## 1.24.2
### Changed
- Updated copart parser.
## 1.24.1
### Changed
- Changed MotorlandAZ production credentials.
## 1.24.0
### Changed
- Return artificial deposits before expiration taking into account lot processing time
## 1.23.0
### Changed
- Allow admin or user's broker to update bid in any state of the participant or lot
- Allow admin or user's broker to create participation for user
## 1.22.0
### Changed
- Switched to new version of copart-parser
## 1.21.2
### Fixed
- Take into account for deposit options customer price
## 1.21.0
### Added
- Allow to specify if limit is suitable for buy now
## 1.20.1
### Fixed
- Lots closer issue with null external id
## 1.20.0
### Added
- Allow brokers to manage participants
## 1.19.2
### Fixed
- Create movement `locked` on creation of withdrawal request
## 1.19.1
### Fixed
- Do not close lots that has participations
## 1.19.0
### Added
- lot closing service: 3 hours after lot end time lot marked as finished
## 1.18.1
### Fixed
- Balance recharge
## 1.18.0
### Added
- Send outbided with to participant event
## 1.17.0
### Added
- Keep alive for live bidders
## 1.16.0
### Added
- Allow to specify limits for each platform separately
## 1.15.0
### Changed
- Async handling of ping requests
## 1.14.0
### Added
- Do not mark as lost participants for sold lots
- Mark as lost participants in deposit pending state when lot has ended
- Mark all participants of the platform for lot as lost
### Fixed
- Race condition on several copart accounts controlling same lane
## 1.13.0
### Added
- App level ping
## 1.12.0
### Added
- Live bidding improvements
### Fixed
- Limit is insufficient error for lot.Currency != deposit.Currency

## 1.11.0
### Added
- Deposit amount can be specified as fraction of lot buy now amount; in that case shelf fee is the deposit value.
- Deposit returning "at lost" and "on won" money charging strategies
  (card or balance for return; shelf fee or all deposits for charging)
  may be specified in the config.
- Auto approve withdrawal option.
- Handle refunded state.

## 1.10.0
### Added
- Create artificial deposits for DEALER account types if buying_power_recharger is enabled
## 1.9.1
### Fix
- Frontend skip `STATE_DEAL_COMPLETED` so shelf fees were not charged;
  added `STATE_IN_TRANSIT` to states for fee charging.

## 1.9.0
### Added
- Allow to bid via live bidding
## 1.8.0
### Changed
- Switched to new copart-bidder API
## 1.7.0
### Added
- Participants management API
## 1.6.1
### Fixed
- Kafka ingestion loop for winner details.

## 1.6.0
### Added
- Winners upsert API
- Publisher for winners
- svc-definition.yaml

## 1.5.0
### Added
- ingest live bidding details
- Use shelf fees specific for requesting and lot platform

## 1.4.0
### Added
- live bidding for copart
- live bidding simulator
### Fixed
- External ID setting

## 1.3.0
### Changed
- Do not attach paid deposits to the lot, if one of the deposits user requested to attach failed

## 1.1.0
### Changed
- Allow to specify customer price to account during calculation of deposit options