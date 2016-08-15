go build -ldflags "-w -s -H=windowsgui"  main.go
rsrc.exe -ico icons/linkedin.ico -arch=amd64 -o main.syso
C:\upx391w\upx -9 main.exe -o test.exe