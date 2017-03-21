@echo off
powershell.exe -ExecutionPolicy Unrestricted %~dp0\compile.ps1 %1 %2 %3
