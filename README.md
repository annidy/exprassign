```
./exprassign -f main.go "COMPILE_TIME=\"`date`\""
```
equal
```
go build -ldflags "-X 'main.COMPILE_TIME=$(date)'" main.go
```
