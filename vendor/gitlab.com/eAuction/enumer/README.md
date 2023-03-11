# enumer

## Project structure

```
gitlab.com/eAuction/enumer
├── cmd
│   ├── enumer # API server
│   └── generate # CSV-Go generator
├── csv # source CSV files for enums and translations
├── internal
│   └── services
│       └── api # enums API implementation
├── pkg
│   ├── maker_model # maker-model finder
│   ├── normalization # normalization package
│   └── state
│       ├── main.go # enums data structure definitions
│       ├── remote # remote state implementation
│       └── static # local static store implementation
└── *.go # public enumer package for go consumers
```

## Updating enums

* edit CSVs (ending in `_enum` or `_l10n`, **NOT** `original/maker_models.csv`; 
    well, you can regenerate using original csv-s, but be careful, binary compatibility may not be preserved)
* `go generate ./...` turns CSVs into `internal/state/static/generated.go`
* `go run cmd/api/main.go` serves updated enums for consumers

## Consuming enums

```go
package main

import (
    "context"
    "gitlab.com/eAuction/enumer/pkg/state/remote"
"net/url"

    "gitlab.com/eAuction/enumer"
)

func main() {
    u, _ := url.Parse("http://enumer")
    // initialize package to use remote source
    g, _ := remote.NewGetter(*u, false)
    enumer.InjectStateGetter(g)

    enumer.IsKnownFuelType(1)
    _, _ = enumer.FindMakerModel("acura", "2.2blabla")
}
```

## legacy enums diff

damage level unknown normalized string is now "other" instead of ""