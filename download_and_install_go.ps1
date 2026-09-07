Write-Host "=== Go Installation Script ===" -ForegroundColor Green
Write-Host ""

# Check if Go is already installed
try {
    $goVersion = go version 2>$null
    if ($goVersion) {
        Write-Host "Go is already installed: $goVersion" -ForegroundColor Yellow
        Write-Host "You can run your enum examples now!" -ForegroundColor Green
        exit 0
    }
} catch {
    # Go is not installed
}

Write-Host "Go is not installed. Let me help you install it..." -ForegroundColor Yellow
Write-Host ""

# Get the latest Go version
Write-Host "Fetching latest Go version..." -ForegroundColor Cyan
$downloadPage = Invoke-WebRequest -Uri "https://go.dev/dl/" -UseBasicParsing
$links = $downloadPage.Links | Where-Object { $_.href -match 'go\d+\.\d+\.\d+\.windows-amd64\.msi$' }

if ($links.Count -eq 0) {
    Write-Host "Could not find download link. Please install manually:" -ForegroundColor Red
    Write-Host "1. Visit: https://go.dev/dl/" -ForegroundColor Cyan
    Write-Host "2. Download the Windows MSI installer" -ForegroundColor Cyan
    Write-Host "3. Run the installer" -ForegroundColor Cyan
    exit 1
}

$downloadUrl = "https://go.dev" + $links[0].href
$fileName = $links[0].href -split '/' | Select-Object -Last 1

Write-Host "Latest Go version: $fileName" -ForegroundColor Green
Write-Host ""

Write-Host "Download location options:" -ForegroundColor Cyan
Write-Host "1. Download to Desktop" -ForegroundColor White
Write-Host "2. Download to Downloads folder" -ForegroundColor White
Write-Host "3. Download to current directory ($PWD)" -ForegroundColor White
Write-Host "4. I'll install manually" -ForegroundColor White

$choice = Read-Host "Enter choice (1-4)"

switch ($choice) {
    "1" { $downloadPath = [Environment]::GetFolderPath("Desktop") }
    "2" { $downloadPath = [Environment]::GetFolderPath("UserProfile") + "\Downloads" }
    "3" { $downloadPath = $PWD.Path }
    "4" { 
        Write-Host ""
        Write-Host "Manual installation instructions:" -ForegroundColor Yellow
        Write-Host "1. Visit: https://go.dev/dl/" -ForegroundColor Cyan
        Write-Host "2. Download: $fileName" -ForegroundColor Cyan
        Write-Host "3. Run the MSI installer" -ForegroundColor Cyan
        Write-Host "4. Open new terminal and run: go version" -ForegroundColor Cyan
        exit 0
    }
    default { 
        Write-Host "Invalid choice. Downloading to current directory." -ForegroundColor Yellow
        $downloadPath = $PWD.Path
    }
}

$fullPath = Join-Path $downloadPath $fileName

Write-Host ""
Write-Host "Downloading Go installer..." -ForegroundColor Cyan
Write-Host "From: $downloadUrl" -ForegroundColor Gray
Write-Host "To: $fullPath" -ForegroundColor Gray

try {
    Invoke-WebRequest -Uri $downloadUrl -OutFile $fullPath
    Write-Host "Download completed successfully!" -ForegroundColor Green
} catch {
    Write-Host "Download failed. Error: $_" -ForegroundColor Red
    Write-Host "Please download manually from: https://go.dev/dl/" -ForegroundColor Yellow
    exit 1
}

Write-Host ""
Write-Host "=== Installation Instructions ===" -ForegroundColor Green
Write-Host ""
Write-Host "Go installer downloaded to: $fullPath" -ForegroundColor Cyan
Write-Host ""
Write-Host "To install Go:" -ForegroundColor Yellow
Write-Host "1. Double-click the file: $fileName" -ForegroundColor White
Write-Host "2. Follow the installation wizard" -ForegroundColor White
Write-Host "3. Accept the license agreement" -ForegroundColor White
Write-Host "4. Choose installation location (default: C:\Go\)" -ForegroundColor White
Write-Host "5. Click 'Install'" -ForegroundColor White
Write-Host ""
Write-Host "After installation:" -ForegroundColor Yellow
Write-Host "1. Open a NEW terminal/CMD window" -ForegroundColor White
Write-Host "2. Verify with: go version" -ForegroundColor White
Write-Host "3. Test with: cd 'C:\Users\kavya\Documents\golang\05_enums' && go run main.go" -ForegroundColor White
Write-Host ""
Write-Host "=== Quick Test ===" -ForegroundColor Green
Write-Host "Once Go is installed, you can test it by running:" -ForegroundColor Cyan
Write-Host "go version" -ForegroundColor White
Write-Host "go env" -ForegroundColor White
Write-Host ""
Write-Host "Press any key to open the download folder..." -ForegroundColor Yellow
Read-Host

# Open the folder
explorer $downloadPath