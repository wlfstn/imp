# The Wolfstone IMP
Short for "Imperial" as in the imperial measurement system primarily used in the United States as well as some industries globally.

## What does IMP provide?
- Split feet/inch inputs evenly into feet/inch (`imp -d 8f8 -s 4`)
- Get total inches from feet/inch input (`impt -in -d 8f8`)

## Example
`imp -d 8f6i -s 2` this will divide 8'6" into two equal sections resulting in 4'3"
`imp -d 14f -s 4` this will divide 14'0" into four equal sections resulting in 3'6"
Note the "i" is optional

# Quick Install with Powershell
```powershell
$exeDownload = "https://github.com/wlfstn/imp/releases/download/V1.2/imp.exe"
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
```