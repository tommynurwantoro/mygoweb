pretty:
	gofmt -s -w .

mod:
	go mod tidy

build:
	go build -o bin/mygoweb cmd/web/main.go

bin:
	./bin/mygoweb

run: pretty mod build bin

deploy:
	sudo cp mygoweb.service /lib/systemd/system/mygoweb.service

# Only for development
dev: build bin
