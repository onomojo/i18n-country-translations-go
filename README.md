# i18n-country-translations-go

> Localized country names for Go -- 168 locales, 257 territory codes, zero external dependencies.

Building a country picker? Displaying addresses internationally? Your users expect to see country names in their own language -- not just in English. Most i18n packages cover 30-50 locales and call it a day, leaving hundreds of millions of speakers without support.

**i18n-country-translations-go** provides country name translations sourced from [CLDR](https://cldr.unicode.org/), the same data that powers ICU, Chrome, and Android. With **168 locales** -- more than double the coverage of alternatives -- you can serve users from Amharic to Zulu without gaps. All translation data is embedded in the binary via `go:embed` -- no external files, no network calls, no filesystem access at runtime.

## Why i18n-country-translations-go?

- **168 locales** -- the most comprehensive coverage available for Go
- **257 territory codes** -- full ISO 3166-1 alpha-2 plus EU, XK, and other commonly used codes
- **CLDR-sourced** -- accurate, professionally reviewed translations (not scraped from Wikipedia)
- **Zero external dependencies** -- translation data is embedded at compile time
- **Thread-safe** -- all operations protected by `sync.RWMutex`, safe for concurrent use
- **Auto-uppercase** -- country codes are automatically uppercased, so `"us"` and `"US"` both work
- **Register what you need** -- load individual locales or all at once

## Install

```bash
go get github.com/onomojo/i18n-country-translations-go
```

Requires Go 1.21+.

## Quick Start

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

## Usage

### Register only what you need

```go
countries.RegisterLocale("de")
countries.RegisterLocale("ja")

name, ok := countries.GetNameForLocale("de", "US")
// name = "Vereinigte Staaten", ok = true

name, ok = countries.GetNameForLocale("ja", "JP")
// name = "日本", ok = true
```

### Register all locales at once

For server-side apps where memory is not a concern:

```go
countries.RegisterAllLocales()
countries.SetDefaultLocale("en")
```

### Using a default locale

Set a default so you don't have to pass a locale every time:

```go
countries.RegisterLocale("de")
countries.SetDefaultLocale("de")

name, _ := countries.GetName("US")   // "Vereinigte Staaten"
name, _ = countries.GetName("JP")    // "Japan"
name, _ = countries.GetName("NO")    // "Norwegen" (not false -- properly handled)
```

### Case-insensitive codes

Country codes are automatically uppercased:

```go
name, _ := countries.GetName("us")   // "Vereinigte Staaten"
name, _ = countries.GetName("Us")    // "Vereinigte Staaten"
```

### Get all names for a locale

```go
all, err := countries.GetAllNames("de")
// all is map[string]string with 257+ entries
```

### List available and registered locales

```go
available := countries.ListLocales()          // all 168 locales in embedded data
registered := countries.ListRegisteredLocales() // only the ones you've loaded
```

## API Reference

| Function | Description |
|----------|-------------|
| `RegisterLocale(locale string) error` | Load translations for a single locale. No-op if already registered. |
| `RegisterAllLocales() error` | Load translations for all 168 available locales. |
| `SetDefaultLocale(locale string) error` | Set the default locale for lookups. Returns error if locale not registered. |
| `GetDefaultLocale() string` | Get the current default locale, or empty string if none set. |
| `GetName(code string) (string, bool)` | Get the localized country name using the default locale. Code auto-uppercased. |
| `GetNameForLocale(locale, code string) (string, bool)` | Get the localized country name for a specific locale. Code auto-uppercased. |
| `GetAllNames(locale string) (map[string]string, error)` | Get all translations for a locale as a new map. |
| `ListLocales() []string` | List all locales available in the embedded data. |
| `ListRegisteredLocales() []string` | List all currently loaded locales. |

All lookup functions return `(value, false)` when a code or locale is not found -- no panics.

## Supported Locales

168 locales covering major and regional languages worldwide:

<details>
<summary>View all 168 locales</summary>

af, ak, am, ar, as, az, be, bg, bm, bn, bo, br, bs, ca, cs, cy, da, de, dz, ee, el, en, eo, es, et, eu, fa, ff, fi, fo, fr, ga, gd, gl, gu, ha, he, hi, hr, hu, hy, ia, id, ig, is, it, ja, ka, ki, kk, kl, km, kn, ko, ky, lg, ln, lo, lt, lu, lv, mg, mk, ml, mn, mr, ms, mt, my, nb, nd, ne, nl, nn, or, pa, pl, ps, pt, pt-BR, rm, rn, ro, ru, se, sg, si, sk, sl, sn, so, sq, sr, sv, sw, ta, te, th, ti, to, tr, uk, ur, uz, vi, yo, zh, zh-CN, zh-HK, zh-TW, zu, asa, bas, bez, brx, byn, cgg, chr, dav, dje, dyo, ebu, ewo, fil, fur, gsw, guz, haw, jmc, kab, kam, kde, kea, khq, kln, ksb, ksf, ksh, lag, luo, luy, mas, mer, mfe, mgh, mua, naq, nmg, nus, nyn, rof, rwk, saq, sbp, seh, ses, shi, swc, teo, tig, twq, tzm, vai, vun, wae, wal, xog, yav

</details>

## Data Source

All translations come from the [Unicode CLDR](https://cldr.unicode.org/) (Common Locale Data Repository) -- the industry-standard source used by every major platform including iOS, Android, Chrome, and Java. This ensures translations are accurate, consistent, and maintained by native speakers through Unicode's established review process.

## Also Available For

- **[Ruby](https://github.com/onomojo/i18n-country-translations)** -- Rails gem with automatic Railtie integration
- **[JavaScript/TypeScript](https://github.com/onomojo/i18n-country-translations-js)** -- NPM package with tree-shaking and reverse lookups
- **[Rust](https://github.com/onomojo/i18n-country-translations-rs)** -- Crate with compile-time embedded data

## License

MIT
