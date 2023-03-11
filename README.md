# buying-power-svc
Buying power is a module responsible for holding and managing users' deposits.
It interacts with payments providers using `payments-api` service.
To participate in the auction, the user has to deposit some funds as collateral for his bids.
Deposits happy path:
* User requests a list of deposits he can perform (we can have several nominations of deposits like 300 USD, 500 USD, and 1000 USD). Available options also depend on the number of funds we hold for the user. Example: Max amount is 1500 USD in total. The user has two active deposits of 500 USD, so we return 300 USD and 500 USD.
* User creates a deposit that redirects him to the payment provider (handled by the front-end and `payment-api`).
* Deposit state is updated in `internal/services/ingest/tbc`. Buying-power-svc listens to incoming events and updates the state of the deposit accordingly
* As soon as the payments provider confirms the deposit, the user can attach the deposit to the active lot. Also corresponding `lead` is created and marked accordingly in the admin panel. It's not possible to attach a deposit to any lot. For example: if live bidding for the lot starts far away in the future before the deposit is automatically unlocked from the user's credit card, it won't be possible to attach. Successful attachment of deposit allows bidding up to a certain amount defined in config.
* If the lot's state changes to lost, we automatically detach the deposit from the lot, so it can be attached again to a new one.
* If the user's bid won the lot and the broker confirms it, `buying-power-svc` charges deposit state from preauthorization to actual charge.

## Live bidding
Buying power contains functionality for live bidding (a limited period when users place bids in real-time. At the end, the winner is selected). It was built on top of Copart's functionality. Unfortunately, as Copart does not provide public API, bid synchronization reliability is insufficient. Thus this feature has been suspended. But we can relatively
easily modify it to implement live bidding for internal inventory.


## Requirements & Building
Refer to `Dockerfile` and `.gitlab-ci.yml`

## Configuration:
See config.yaml as an example of a valid configuration.

## Docs
Documentation source code is available at `docs`. Rendered version is [here](https://eauction.gitlab.io/buying-power-svc/)

# buying-power-svc
