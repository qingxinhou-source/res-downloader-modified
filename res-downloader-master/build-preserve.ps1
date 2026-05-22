$ErrorActionPreference = "Stop"

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Build package (preserve history)" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

$timestamp = Get-Date -Format "yyyyMMdd_HHmmss"
$sourceFile = "build/bin/res-downloader-modified.exe"
$newFileName = "res-downloader-modified_$timestamp.exe"
$destinationFile = "build/bin/$newFileName"

Write-Host "Building frontend..." -ForegroundColor Yellow
Push-Location "frontend"
npm run build
$frontendExitCode = $LASTEXITCODE
Pop-Location

if ($frontendExitCode -ne 0) {
    Write-Host ""
    Write-Host "Frontend build failed. Exit code: $frontendExitCode" -ForegroundColor Red
    exit $frontendExitCode
}

Write-Host ""
Write-Host "Building Windows executable..." -ForegroundColor Yellow
wails build -f -s -skipbindings -platform "windows/amd64" -webview2 Embed -ldflags "-s -w"
$goExitCode = $LASTEXITCODE

if ($goExitCode -ne 0) {
    Write-Host ""
    Write-Host "Executable build failed. Exit code: $goExitCode" -ForegroundColor Red
    exit $goExitCode
}

if (-not (Test-Path $sourceFile)) {
    Write-Host ""
    Write-Host "Build finished, but output was not found: $sourceFile" -ForegroundColor Red
    exit 1
}

Copy-Item -Path $sourceFile -Destination $destinationFile

Write-Host ""
Write-Host "========================================" -ForegroundColor Green
Write-Host "  Build complete" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Green
Write-Host ""
Write-Host "Generated file:" -ForegroundColor Cyan
Write-Host "   $destinationFile" -ForegroundColor Yellow
Write-Host ""

Write-Host "All historical builds:" -ForegroundColor Cyan
Get-ChildItem "build/bin/*.exe" | Sort-Object LastWriteTime -Descending | ForEach-Object {
    $sizeMB = [math]::Round($_.Length / 1MB, 2)
    Write-Host "   $($_.Name) ($sizeMB MB)" -ForegroundColor Gray
}
Write-Host ""
