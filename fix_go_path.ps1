Write-Host "=== Fix Go PATH Issue ===" -ForegroundColor Green
Write-Host ""

# Check current PATH
Write-Host "Current PATH contains Go?" -ForegroundColor Cyan
$goInPath = $env:Path -split ";" | Where-Object { $_ -match "\\Go\\bin" }
if ($goInPath) {
    Write-Host "✅ Go is already in PATH:" -ForegroundColor Green
    $goInPath | ForEach-Object { Write-Host "   $_" -ForegroundColor White }
} else {
    Write-Host "❌ Go is NOT in PATH" -ForegroundColor Red
}

Write-Host ""
Write-Host "Go installation found at: C:\Program Files\Go\bin\go.exe" -ForegroundColor Cyan

Write-Host ""
Write-Host "=== Fix Options ===" -ForegroundColor Yellow
Write-Host "1. Add Go to PATH temporarily (current session only)" -ForegroundColor White
Write-Host "2. Create a script to always add Go to PATH" -ForegroundColor White
Write-Host "3. Show manual instructions to add to system PATH" -ForegroundColor White
Write-Host "4. Test Go installation" -ForegroundColor White

$choice = Read-Host "`nEnter choice (1-4)"

switch ($choice) {
    "1" {
        # Temporary fix
        $env:Path += ";C:\Program Files\Go\bin"
        Write-Host "`n✅ Added Go to PATH temporarily" -ForegroundColor Green
        Write-Host "Go will work in THIS terminal session only" -ForegroundColor Yellow
        Write-Host "Test with: go version" -ForegroundColor Cyan
        go version
    }
    "2" {
        # Create profile script
        $profileDir = Split-Path $PROFILE -Parent
        if (-not (Test-Path $profileDir)) {
            New-Item -ItemType Directory -Path $profileDir -Force
        }
        
        $goPathLine = '`$env:Path += ";C:\Program Files\Go\bin"'
        Add-Content -Path $PROFILE -Value "`n# Add Go to PATH`n$goPathLine"
        
        Write-Host "`n✅ Created PowerShell profile script" -ForegroundColor Green
        Write-Host "Go will be added to PATH every time you open PowerShell" -ForegroundColor Cyan
        Write-Host "Profile location: $PROFILE" -ForegroundColor White
    }
    "3" {
        Write-Host "`n=== Manual Instructions ===" -ForegroundColor Green
        Write-Host "1. Press Win + X → 'System'" -ForegroundColor Cyan
        Write-Host "2. Click 'Advanced system settings'" -ForegroundColor Cyan
        Write-Host "3. Click 'Environment Variables'" -ForegroundColor Cyan
        Write-Host "4. Under 'System variables', find 'Path' → 'Edit'" -ForegroundColor Cyan
        Write-Host "5. Click 'New' and add: C:\Program Files\Go\bin" -ForegroundColor Cyan
        Write-Host "6. Click OK, OK, OK" -ForegroundColor Cyan
        Write-Host "7. CLOSE and REOPEN all terminal windows" -ForegroundColor Red
    }
    "4" {
        Write-Host "`n=== Testing Go Installation ===" -ForegroundColor Green
        Write-Host "Trying with full path:" -ForegroundColor Cyan
        & "C:\Program Files\Go\bin\go.exe" version
        
        Write-Host "`nTrying with 'go' command:" -ForegroundColor Cyan
        try {
            go version
        } catch {
            Write-Host "❌ 'go' command not found in PATH" -ForegroundColor Red
        }
        
        Write-Host "`nTrying enum example:" -ForegroundColor Cyan
        try {
            cd "c:\Users\kavya\Documents\golang\01-basics\05_enums"
            go run main.go
        } catch {
            Write-Host "❌ Could not run enum example" -ForegroundColor Red
        }
    }
}

Write-Host "`n=== Quick Test ===" -ForegroundColor Green
Write-Host "Try these commands in NEW terminal after fixing:" -ForegroundColor Cyan
Write-Host "go version" -ForegroundColor White
Write-Host "cd `"c:\Users\kavya\Documents\golang\01-basics\05_enums`"" -ForegroundColor White
Write-Host "go run main.go" -ForegroundColor White

Write-Host "`nPress any key to exit..." -ForegroundColor Yellow
Read-Host