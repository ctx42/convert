[![Go Report Card](https://goreportcard.com/badge/github.com/ctx42/convert)](https://goreportcard.com/report/github.com/ctx42/convert)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/ctx42/convert)
![Tests](https://github.com/ctx42/convert/actions/workflows/go.yml/badge.svg?branch=master)

`convert` is Go module providing functions for safe converting between Go types. 
Conversion functions return an error if conversion causes truncation, overflow,
or semantic loss.
