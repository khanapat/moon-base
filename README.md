# Moon Base API (Backend API, Database, Frontend Web(Coming soon))

[![Github Actions Status](https://github.com/gbrlsnchs/jwt/workflows/Linux,%20macOS%20and%20Windows/badge.svg)](https://github.com/gbrlsnchs/jwt/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/gbrlsnchs/jwt)](https://goreportcard.com/report/github.com/gbrlsnchs/jwt)
[![GoDoc](https://godoc.org/github.com/gbrlsnchs/jwt?status.svg)](https://pkg.go.dev/github.com/gbrlsnchs/jwt/v3)
[![Version compatibility with Go 1.11 onward using modules](https://img.shields.io/badge/compatible%20with-go1.11+-5272b4.svg)](https://github.com/gbrlsnchs/jwt#installing)

## ABOUT

This repository is a backend api for [Go](https://golang.org) (or Golang). This api connect to web moon base to provide data and store total supply.

## API Service

swaager

http://localhost:9090/moon-coin/swagger/index.html#/

* Buy Coin
> Validate request, Check total supply, Update total supply and record history

* Get History
> Query history transaction with string param from and to date

* Get Supply
> Query remaining supply at this moment

* Reset and Setup Database
> Create table from supply coin and history and Insert initialize moon base coin in table


## Makefile
### MSSQL
```bash
make create-mssql
make delete-mssql
```

### Backend-Api
```bash
make run

make create-api
```
