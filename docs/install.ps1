$exeDownload = "https://github.com/wlfstn/imp/releases/download/v1.3.0/imp.exe"
$destinationDir = "$env:USERPROFILE\wlfstn\imp"
$destinationFile = Join-Path $destinationDir "imp.exe"

if (-not (Test-Path $destinationDir)) {
	New-Item -ItemType Directory -Path $destinationDir -Force | Out-Null
}

Invoke-WebRequest -Uri $exeDownload -OutFile $destinationFile

$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if (-not $userPath) {
	$userPath = ""
}

if ($userPath -split ";" -contains $destinationDir) {
	Write-Host "Imp directory is already in PATH."
} else {
	$addToPath = Read-Host "Would you like to add imp to your PATH? (y/n)"
	if ($addToPath -match '^[Yy]$') {
		$newPath = if ($userPath -eq "") {
			$destinationDir
		} else {
			"$userPath;$destinationDir"
		}
		[Environment]::SetEnvironmentVariable("Path", $newPath, "User")
		Write-Host "Added $destinationDir to user PATH."
	} else {
		Write-Host "Installed, but skipped adding to PATH."
	}
}