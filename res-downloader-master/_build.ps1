$ErrorActionPreference = "Stop"
Set-Location "g:\res-downloader-master\res-downloader-master"

Write-Host "=== Cleaning frontend dist ===" -ForegroundColor Cyan
Remove-Item -Recurse -Force "frontend\dist" -ErrorAction SilentlyContinue

Write-Host "=== Building ===" -ForegroundColor Yellow
wails build 2>&1 | ForEach-Object { $_.ToString() }
$exitCode = $LASTEXITCODE

Write-Host ""
if ($exitCode -eq 0) {
    $exe = Get-Item "build/bin/res-downloader-modified.exe" -ErrorAction SilentlyContinue
    if ($exe) {
        Write-Host "BUILD OK: $($exe.LastWriteTime) ($([math]::Round($exe.Length/1MB,2)) MB)" -ForegroundColor Green
        # 时间戳副本
        $ts = Get-Date -Format "yyyyMMdd_HHmmss"
        Copy-Item $exe.FullName "build/bin/res-downloader-modified_$ts.exe"
        Write-Host "SAVED: res-downloader-modified_$ts.exe" -ForegroundColor Green
        
        Write-Host "`n=== All builds ===" -ForegroundColor Cyan
        Get-ChildItem "build/bin/*.exe" | Sort-Object LastWriteTime -Descending | ForEach-Object {
            Write-Host "  $($_.Name)  $([math]::Round($_.Length/1MB,2)) MB  $($_.LastWriteTime.ToString('HH:mm:ss'))" -ForegroundColor Gray
        }
    } else {
        Write-Host "ERROR: exe not found after 'success' build!" -ForegroundColor Red
    }
} else {
    Write-Host "BUILD FAILED: exit code $exitCode" -ForegroundColor Red
}
