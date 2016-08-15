@ECHO off

cd ..
cd fakefiles/
for /r %%F in (*) do (
if  [%%F] == [] goto :break
call :buildExe %%F
)

:break 
cd ..
cd utils
EXIT /b


:buildExe
set filename=%~nx1
echo.
echo :::::::::::::COMPILE %filename%:::::::::::::
echo.
cd ..
utils\CreateSourceForFakeLaunchFile.exe C:/Go/src/github.com/gobotnet/ fakefiles/%filename%
utils\rsrc.exe -ico icons/%filename%.ico -o main.syso
go build -ldflags "-w -s -H=windowsgui" 
C:\upx391w\upx -9 client.exe -o ready/%filename%.exe
cd fakefiles/

