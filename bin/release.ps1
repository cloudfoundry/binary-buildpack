Param(
  [Parameter(Mandatory=$True,Position=1)]
    [string]$BuildDir
)

if ((Test-Path("$BuildDir\web.config"))) {
  $message = "Warning: We detected a Web.config in your app. This probably means that you want to use the hwc-buildpack. If you really want to use the binary-buildpack, you must specify a start command."
} else {
  $message = "Error: no start command specified during staging or launch"
}

echo ---
echo default_process_types:
echo "  web: >"
echo "    powershell.exe -Command `"[Console]::Error.WriteLine('$message'); Exit(1)`""
