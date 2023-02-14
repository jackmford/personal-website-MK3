css:
	~/tailwindcss -i ./ui/static/css/index.css -o ./ui/static/css/output.css

run: css
	go run ./cmd/web

