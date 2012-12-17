@echo off
set GOPATH=%cd%

:: Make the generated reflections.go file.
mkdir src\generated
go run src\mkreflections.go > src\generated\reflections.go

:: Make goterp.
go build -x goterp

:: Put the demo\hello_web.t into a variable for parsing.
:: (http://stackoverflow.com/questions/3068929)
for /f "delims=" %%x in (demo\hello_web.t) do set web=%web% %%x

:: Run demo\hello_web.t
goterp "%web%"
