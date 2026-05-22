@echo off
echo Taking ownership of dist\assets...
icacls "G:\res-downloader-master\res-downloader-master\frontend\dist\assets" /setowner %USERNAME% /T /C
echo.
echo Granting full control...
icacls "G:\res-downloader-master\res-downloader-master\frontend\dist\assets" /grant %USERNAME%:(OI)(CI)F /T /C
echo.
echo Deleting dist...
rd /s /q "G:\res-downloader-master\res-downloader-master\frontend\dist"
echo.
if exist "G:\res-downloader-master\res-downloader-master\frontend\dist" (
    echo FAILED - dist still exists
) else (
    echo SUCCESS - dist deleted
)
pause
