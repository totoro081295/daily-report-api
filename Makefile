build:
	go build
	./daily-report-api

seeds:
	go build -o seeds db/seeds/seed.go
	./seeds