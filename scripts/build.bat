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

REM Create a separate folder for the project
set TOOL_FOLDER=%TOOLS_PATH%\node-project-tool

if exist "%TOOL_FOLDER%" (
    echo Tools folder already exists.
) else (
    echo Creating tools folder already exists.
    mkdir %TOOL_FOLDER%

    if errorlevel 1 (
        echo Error: Failed to create a folder for the tool.
        exit /b 1
    )
)

REM Building the go project
set exeFile=node-project-tool.exe
go build -o %exeFile%

if errorlevel 1 (
    echo Error: Failed to build the project.
    exit /b 1
)

REM Copy the executable to the tools folder
echo Copying %exeFile% to %TOOL_FOLDER%
copy %exeFile% %TOOL_FOLDER%

if errorlevel 1 (
    echo Error: Failed to copy
    exit /b 1
)

REM Copy templates folder to tools directory
echo Copying templates folder to %TOOL_FOLDER%
xcopy /E /I templates %TOOL_FOLDER%\templates

if errorlevel 1 (
    echo Error: Failed to copy templates folder
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