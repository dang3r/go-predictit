
build: pred.go structs.go
	go build -o gopred pred.go structs.go

test : test.go
	go build -o test test.go
	./test
