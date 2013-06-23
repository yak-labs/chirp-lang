@echo off

if "%GOPATH%" == "" (
    set GOPATH=%cd%
)

if "%1" == "" goto help

if "%1" == "help" (
    :help
    echo.Please use `make ^<target^>` where ^<target^> is one of
    echo.  build        to build Chirp.
    echo.  reflections  to build Chirp's reflection library.
    echo.  test         to run tests.
    goto end
)

if "%1" == "clean" (
    del /q /s goterp.exe
    del /q /s src\generated\reflections.go
    del /q /s pkg\*
    del /q /s demo\_*.err
    rmdir /q /s .\data
)

if "%1" == "build" (
    go build -x goterp
    goto end
)

if "%1" == "reflections" (
    go run src\mkreflections.go > src\generated\reflections.go
    goto end
)

if "%1" == "test" (
    go test -i terp
    go test terp
    goto end
)

:end
