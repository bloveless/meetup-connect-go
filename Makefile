# go install github.com/mitranim/gow@latest
dev: generate
	mprocs "gow run cmd/server/main.go" "cd web && npm run dev"

generate:
	rm -rf gen
	buf generate
	cd web && rm -rf src/gen
	cd web && pnpm generate

build:
	cd web && pnpm build
	go build -o connect-go ./cmd/server/main.go

present:
	cd presentation && npm i && npm start
