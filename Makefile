up: 
	docker compose up
down:
	docker compose down
table-create:
	docker compose exec api migrate create -ext sql -dir ./database/migrations -seq $(table)

