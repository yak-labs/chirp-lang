@echo off
set GOPATH=%cd%

go test -i %1%
go test %1%
