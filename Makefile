gb:
	go build -o ./bin/gb ./cmd/main.go

run:
	go run ./cmd/main.go < ./cfg/config.json

clean:
	rm ./bin/gb

