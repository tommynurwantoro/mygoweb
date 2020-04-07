pretty:
	gofmt -s -w .

mod:
	go mod tidy

build:
	go build -o bin/mygoweb cmd/web/main.go

bin:
	./bin/mygoweb

run: pretty dep build bin
	./bin/mygoweb

deploy:
	sudo cp mygoweb.service /lib/systemd/system/mygoweb.service

# Only for development
dev:
	go build -o bin/mygoweb app/main/main.go
	./bin/mygoweb
