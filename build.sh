GOOS=linux GOARCH=amd64 go build -o build/go-pic-bed_linux
GOOS=windows GOARCH=amd64 go build -o build/go-pic-bed_windows.exe
GOOS=darwin GOARCH=amd64 go build -o build/go-pic-bed_macos
echo "exec build success"
docker build -t go-pic-bed .
echo "docker build success"