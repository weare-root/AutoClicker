if not exist "build" mkdir build

rmdir build\ui
mkdir build\ui

xcopy ui build\ui /E /y

go build -o build -ldflags -H=windowsgui

echo DONT FORGET TO COPY THE DLLS INCASE YOU HAVENT DONE THAT
