@echo off
set GOPATH=%cd%

:: Make the generated reflections.go file.
mkdir src\generated
go run src\mkreflections.go > src\generated\reflections.go

:: Make goterp.
go build -x goterp
