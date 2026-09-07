# Go Installation Guide for Windows
Write-Host "=== Go Installation Guide ===" -ForegroundColor Green
Write-Host ""
Write-Host "METHOD 1: Download and Install Manually" -ForegroundColor Yellow
Write-Host "1. Visit: https://go.dev/dl/" -ForegroundColor Cyan
Write-Host "2. Download: go1.x.x.windows-amd64.msi (latest version)" -ForegroundColor Cyan
Write-Host "3. Run the .msi file and follow the installer" -ForegroundColor Cyan
Write-Host "4. Open NEW terminal after installation" -ForegroundColor Cyan
Write-Host "5. Verify with: go version" -ForegroundColor Cyan
Write-Host ""
Write-Host "METHOD 2: Using Chocolatey Package Manager" -ForegroundColor Yellow
Write-Host "If you have Chocolatey installed, run:" -ForegroundColor Cyan
Write-Host "   choco install golang" -ForegroundColor White
Write-Host ""
Write-Host "To install Chocolatey:" -ForegroundColor Cyan
Write-Host '   Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString("https://community.chocolatey.org/install.ps1"))' -ForegroundColor White
Write-Host ""
Write-Host "=== After Installation ===" -ForegroundColor Green
Write-Host "1. Verify Go installation:" -ForegroundColor Cyan
Write-Host "   go version" -ForegroundColor White
Write-Host ""
Write-Host "2. Set up your Go workspace:" -ForegroundColor Cyan
Write-Host "   Go will automatically use: C:\Users\kavya\go" -ForegroundColor White
Write-Host ""
Write-Host "3. Test with our enum example:" -ForegroundColor Cyan
Write-Host "   cd 'C:\Users\kavya\Documents\golang\05_enums'" -ForegroundColor White
Write-Host "   go run main.go" -ForegroundColor White
Write-Host ""
Write-Host "=== Troubleshooting ===" -ForegroundColor Green
Write-Host "If 'go' command not found:" -ForegroundColor Yellow
Write-Host "1. Close and reopen your terminal" -ForegroundColor Cyan
Write-Host "2. Check PATH: echo %PATH%" -ForegroundColor Cyan
Write-Host "3. Make sure C:\Go\bin is in PATH" -ForegroundColor Cyan
Write-Host ""
Write-Host "=== Quick Test After Installation ===" -ForegroundColor Green
Write-Host "Run these commands to test:" -ForegroundColor Cyan
Write-Host "go version" -ForegroundColor White
Write-Host "go env GOPATH" -ForegroundColor White
Write-Host "cd 'C:\Users\kavya\Documents\golang\05_enums' && go run main.go" -ForegroundColor White