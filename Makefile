pre-commit:
	go mod tidy
	go vet ./...
	go fmt ./...
help:
	go run main.go --help
version:
	go run main.go version
generate-root-ca-key:
	go run main.go generate-rsa-key 4096 root-ca-key
generate-root-ca-cert:
	go run main.go generate-cert ./root-ca-key.pem 1 root-ca-cert
generate-service-key:
	go run main.go generate-rsa-key 4096 service-key
generate-csr:
	go run main.go generate-csr service-key.pem service test@example.com example.com AU Some-State MyCity Company-Ltd IT
sign-csr:
	go run main.go sign-csr root-ca-key.pem root-ca-cert.pem service.csr 1 service-cert
