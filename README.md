# go-lclocale

[![Test](https://github.com/elliotcourant/go-lclocale/actions/workflows/test.yaml/badge.svg)](https://github.com/elliotcourant/go-lclocale/actions/workflows/test.yaml)
[![Documentation](https://pkg.go.dev/badge/github.com/elliotcourant/go-lclocale)](https://pkg.go.dev/github.com/elliotcourant/go-lclocale)
[![Go Report Card](https://goreportcard.com/badge/github.com/elliotcourant/go-lclocale)](https://goreportcard.com/report/github.com/elliotcourant/go-lclocale)

This package is to expose Linux LC_MONETRY locale APIs in Golang via Cgo. The main purpose of this is to make it easier
to write localization packages that rely on the hosts actual localization data, rather than relying on some other data
source that might be inconsistent between applications.

## Requirements

This package requires Cgo, you cannot import depend on it without having `gcc` installed on the system you are building
on.

If you do not have `gcc` installed you will see errors like:

```
# github.com/elliotcourant/go-lclocale
..\go\pkg\mod\github.com\elliotcourant\go-lclocale@v0.0.3\monetary.go:84:2: undefined: localeMutex
..\go\pkg\mod\github.com\elliotcourant\go-lclocale@v0.0.3\monetary.go:85:8: undefined: localeMutex
..\go\pkg\mod\github.com\elliotcourant\go-lclocale@v0.0.3\monetary.go:86:12: undefined: setLocale
..\go\pkg\mod\github.com\elliotcourant\go-lclocale@v0.0.3\monetary.go:90:12: undefined: localeconv
..\go\pkg\mod\github.com\elliotcourant\go-lclocale@v0.0.3\search.go:20:13: undefined: setLocale
```
