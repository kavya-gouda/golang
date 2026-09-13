@echo off
echo ========================================
echo    Go Terminal with PATH Fixed
echo ========================================
echo.
echo This terminal has Go added to PATH.
echo You can run 'go' commands here!
echo.

REM Add Go to PATH for this session
set PATH=%PATH%;C:\Program Files\Go\bin

REM Show current directory
echo Current directory: %CD%
echo.

REM Show Go version
echo Go version:
go version
echo.

REM Menu
:menu
echo ========================================
echo    Available Commands
echo ========================================
echo.
echo 1. Run enum example
echo 2. Check Go environment
echo 3. Go to enum directory
echo 4. Create new Go file
echo 5. Clear screen
echo 6. Exit
echo.

set /p choice="Enter choice (1-6): "

if "%choice%"=="1" goto run_enum
if "%choice%"=="2" goto go_env
if "%choice%"=="3" goto enum_dir
if "%choice%"=="4" goto new_file
if "%choice%"=="5" goto clear
if "%choice%"=="6" goto exit

echo Invalid choice!
goto menu

:run_enum
cd "C:\Users\kavya\Documents\golang\01-basics\05_enums"
echo Running enum example...
go run main.go
goto menu

:go_env
echo Go environment:
go env
goto menu

:enum_dir
cd "C:\Users\kavya\Documents\golang\01-basics\05_enums"
dir
goto menu

:new_file
set /p filename="Enter filename (e.g., test.go): "
echo package main > %filename%
echo. >> %filename%
echo import "fmt" >> %filename%
echo. >> %filename%
echo func main() { >> %filename%
echo     fmt.Println("Hello from Go!") >> %filename%
echo } >> %filename%
echo Created %filename%
goto menu

:clear
cls
goto menu

:exit
echo Goodbye!
pause
exit