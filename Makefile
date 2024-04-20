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
generate-token:
	go run main.go generate-token service-key.pem "{\"name\":\"HaoLe\"}"
verify-token:
	go run main.go verify-token root-ca-cert.pem service-cert.pem eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJkYXQiOiIvVXNlcnMvaGFvbGUvRG9jdW1lbnRzL3Jvb3QtY2VydGlmaWNhdGUtYXV0aG9yaXR5LWp3dC97XCJuYW1lXCI6XCJIYW9MZVwifSIsImV4cCI6MTcxMzYyNjkzMSwiaWF0IjoxNzEzNjIzMzMxLCJuYmYiOjE3MTM2MjMzMzF9.KhOptbIQ3ZwInnU5pNIvq5MYVsvGLd8Dk5L31vGgHEHvJldGrYW7rdE_7DoaLWhZX-4j9NHtJMKVqofy_eAFutsOa1dQVl5IZh50ZtFybPDmASlfp1H9uOomI1GPlqq-V_AcIVLD3RQXEwCHSQNorqleHhL3HbDienp4ROdyIw_WezXl4gBOhqZdi0zdRxMGJnzKuqWNv0_4fcPYnH4uYXPTGNSXp0xDSYsY-S-YLPRMCtSpGRwZ2V_DFMd3X3P3fqudZCtqivU8agCKOxS1aeLOn8JE4S9CAZLeqILVoWiQ9y4ZsX0DWVRQm_e8QdHuY6cFeWcg3Pl1QJGiwC4xBZYI38dWeDyy9_vrL5LLQQVhFHAG6Ofe5kwUr8JNWrOn65JGKkEzS76B82Laa5mHTR2q1srF8Y-iXDmd3eR-0oAo_1V83_50Ksh9SIpGr9nAolAhovGARjzQJQaHie1xu3YnKkxt2xVWwpyH5fs1TdQ817PG39Q55URvPq39W2qjrExqEVNZqkJvkZS4u-_znDE3nPWCLhCrYwyZBZd_hs3NfVwHxkDvNpo6VXtbTsLi3oreScLNvop6TfRPoyFiX7zogcLBhoSnFFG6oi3u8I8-MfG3XmMxBhEfXTTKA-h41trM23jr_y9vIrwpqX6oFjrl35idbjN2wc6Q_nXXsow