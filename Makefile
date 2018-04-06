binary:
	GOOS=linux GOARCH=arm \
	go build \
	-o app \
	main.go
image:
	docker build \
	-t nycdavid/hello-machine-id:latest \
	.
