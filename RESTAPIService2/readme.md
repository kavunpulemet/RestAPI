docker run --name=todo-app -e POSTGRES_PASSWORD='1488' -p 5436:5432 -d --rm postgres

migrate -path ./schema -database 'postgres://postgres:1488@localhost:5436/postgres?sslmode=disable' up

docker ps
docker exec -it bf39dfdd9f49 /bin/bash
psql -U postgres
\d