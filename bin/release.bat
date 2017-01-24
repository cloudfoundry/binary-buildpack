@echo off

set build_dir=%1
:: output valid yaml containing the start command

echo ---
echo default_process_types:
echo   web: cmd.exe /c exit 1

exit /b 0
