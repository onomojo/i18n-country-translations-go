# i18n-country-translations-go

Go module for localized country name translations. Covers 257 territory codes across 168 locales, sourced from Unicode CLDR.

## Install

```bash
go get github.com/onomojo/i18n-country-translations-go
```

## Usage

```go
package main

import (
    "fmt"
    countries "github.com/onomojo/i18n-country-translations-go"
)

func main() {
    countries.RegisterLocale("de")
    countries.SetDefaultLocale("de")

    name, _ := countries.GetName("US")
    fmt.Println(name) // Vereinigte Staaten
}
```

## API

- `RegisterLocale(locale string) error`
- `RegisterAllLocales() error`
- `SetDefaultLocale(locale string) error`
- `GetDefaultLocale() string`
- `GetName(code string) (string, bool)` — code auto-uppercased
- `GetNameForLocale(locale, code string) (string, bool)`
- `GetAllNames(locale string) (map[string]string, error)`
- `ListLocales() []string` — available in embedded data
- `ListRegisteredLocales() []string` — currently loaded

## License

MIT
