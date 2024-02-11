# go-lclocale

[![Test](https://github.com/elliotcourant/go-lclocale/actions/workflows/test.yaml/badge.svg)](https://github.com/elliotcourant/go-lclocale/actions/workflows/test.yaml)
[![Documentation](https://pkg.go.dev/badge/github.com/elliotcourant/go-lclocale)](https://pkg.go.dev/github.com/elliotcourant/go-lclocale)

This package is to expose Linux locale APIs in Golang via Cgo. The main purpose of this is to make it easier to write
localization packages that rely on the hosts actual localization data, rather than relying on some other data source
that might be inconsistent between applications.

