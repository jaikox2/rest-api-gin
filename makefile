dev:
		go run httpd/main.go

test:
		go test ./...

test-cover:
		go test --cover ./...