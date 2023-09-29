gentypes:
	rm -rf gen
	buf generate
	cd web && rm -rf src/gen
	cd web && pnpm generate

build:
	cd web && pnpm build
	go build -o connect-go ./cmd/server/main.go
