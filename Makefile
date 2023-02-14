css:
	~/tailwindcss -i ./ui/static/css/index.css -o ./ui/static/css/output.min.css

run: css
	go run ./cmd/web

