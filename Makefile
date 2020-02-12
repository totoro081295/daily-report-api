build:
	go build
	./daily-report-api

# deploy to staging
deploy:
	git checkout -b deploy
	GOOS=linux GOARCH=amd64 go build -o db/seeds/seeds db/seeds/seed.go
	GOOS=linux GOARCH=amd64 go build
	git add -f daily-report-api db/seeds/seeds
	git commit -m 'deploy to heroku'
	git push -f heroku deploy:master
