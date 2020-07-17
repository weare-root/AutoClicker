if not exist "build" mkdir build

xcopy ui build /E

go build -o build

echo DONT FORGET TO COPY THE DLLS INCASE YOU HAVENT DONE THAT
