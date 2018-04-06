binary:
	GOOS=linux GOARCH=amd64 \
	go build \
	-o app \
	main.go
image:
	docker build \
	-t nycdavid/hello-machine-id:latest \
	.
