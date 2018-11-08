run:
	PSQL_HOST=localhost PSQL_USER=postgres PSQL_PASS='' \
	PSQL_DB=postgres PSQL_PORT=5432 \
	go run app.go

db: 
	docker ps | grep -q postgres-webapp || docker run --rm -d -p 5432:5432 --name postgres-webapp postgres 

build:
	docker build -t webapp .

runimage:  build db
	docker run --rm -it --link postgres-webapp:postgres \
		-e PSQL_HOST=postgres -e PSQL_USER=postgres -e PSQL_PASSWORD='' -e PSQL_PORT=5432 -e PSQL_DB=postgres \
		-p 8080:8080 \
		webapp 
