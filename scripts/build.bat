@echo off
setlocal

REM Get tools folder from the environment variables
set TOOLS_PATH=%TOOLS%
echo Tools folder path: %TOOLS_PATH%

REM Check if the environment variable is not set
if "%TOOLS_PATH%"=="" (
  echo Error: TOOLS environment variable is not set.
  exit /b 1
)

REM Building the go project
set exeFile=node-project-tool.exe
go build -o %exeFile%

if errorlevel 1 (
    echo Error: Failed to build the project.
    exit /b 1
)

REM Copy the executable to the tools folder
echo Copying %exeFile% to %TOOLS_PATH%
copy %exeFile% %TOOLS_PATH%

if errorlevel 1 (
    echo Error: Failed to copy
    exit /b 1
)

echo Deleting %exeFile%
del %exeFile%

if errorlevel 1 (
    echo Error: Failed to delete %exeFile%
    exit /b 1
)

echo Build and copy successful!
endlocal