install:
	go build -o csc main.go
	sudo mv csc /usr/bin
	sudo mv Modules /usr/bin
	sudo mv main.go /usr/bin
	sudo mv go.mod /usr/bin

dev:
	go build -o csc_dev main.go
