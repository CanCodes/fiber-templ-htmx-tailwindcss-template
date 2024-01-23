package_manager = pnpm

all: build

templ: tailwind
	templ generate -path ./components

tailwind: init
	$(package_manager) run build:css

run:
	templ generate --watch --cmd="go run ."

build: templ
	go build -o app

init:
	$(package_manager) install
