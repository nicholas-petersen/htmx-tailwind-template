#!make
include .env

define tailwind_config
module.exports = {
  content: [
	'./web/template/*.html',
	'./web/template/views/*.html',
	'./web/template/partials/*.html',
	],
  theme: {
    extend: {},
  },
  plugins: [],
}
endef
export tailwind_config

run:
	@go run cmd/web/main.go

htmx.install:
	@wget -q https://unpkg.com/htmx.org/dist/htmx.min.js -O web/static/js/htmx.min.js

tailwindcss.install: 
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
	chmod +x tailwindcss-macos-arm64
	mv tailwindcss-macos-arm64 tailwindcss

tailwindcss.init:
	@echo "@tailwind base;\n@tailwind components;\n@tailwind utilities;" > web/main.css 
	@echo "$$tailwind_config" > tailwind.config.js

tailwindcss:
	./tailwindcss -i web/main.css -o web/static/css/style.css

tailwindcss.watch:
	./tailwindcss -i web/main.css -o web/static/css/style.css --watch

.PHONY: tailwindcss
