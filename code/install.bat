@echo off

setlocal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end

:ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0;%~dp0..\thirdparty

gofmt -w src

go install chapter01/io

set GOPATH=%OLDGOPATH%

:end
echo finished