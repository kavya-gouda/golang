@echo off
echo ========================================
echo    Go Installation Guide for Windows
echo ========================================
echo.
echo METHOD 1: Download and Install Manually
echo   1. Visit: https://go.dev/dl/
echo   2. Download: go1.x.x.windows-amd64.msi
echo   3. Run the .msi file and follow the installer
echo   4. Open NEW CMD window after installation
echo   5. Verify with: go version
echo.
echo METHOD 2: Using Chocolatey Package Manager
echo   If you have Chocolatey installed, run:
echo      choco install golang
echo.
echo To install Chocolatey:
echo   Open PowerShell as Administrator and run:
echo   powershell -Command "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))"
echo.
echo ========================================
echo    After Installation
echo ========================================
echo.
echo 1. Verify Go installation:
echo    go version
echo.
echo 2. Test with our enum example:
echo    cd "C:\Users\kavya\Documents\golang\05_enums"
echo    go run main.go
echo.
echo ========================================
echo    Troubleshooting
echo ========================================
echo.
echo If 'go' command not found:
echo   1. Close and reopen CMD
echo   2. Check PATH: echo %%PATH%%
echo   3. Make sure C:\Go\bin is in PATH
echo.
echo Press any key to exit...
pause > nul