if not exist "build" mkdir build

rmdir build\resources
mkdir build\resources

xcopy resources build\resources /E /y

go build -o build -ldflags -H=windowsgui

echo DONT FORGET TO COPY THE DLLS INCASE YOU HAVENT DONE THAT
