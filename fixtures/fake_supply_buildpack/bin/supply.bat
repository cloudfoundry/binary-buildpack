@echo off

ECHO Running Fake Supply Buildpack
REM source: ./bin/supply.bat -- ./my_buildpack_assets/main.exe
set BIN_SOURCE_PATH=%0\..\..\my_buildpack_assets\main.exe
set BUILDPACK_DEST_BIN_DIR=%3\%4\bin

REM Create buildpack destination bin directory
MKDIR %BUILDPACK_DEST_BIN_DIR%

REM Copy main to bin directory which libbuildpack will append to the PATH
COPY %BIN_SOURCE_PATH% %BUILDPACK_DEST_BIN_DIR%\main.exe

EXIT 0