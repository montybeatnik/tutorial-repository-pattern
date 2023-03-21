SHELL := /bin/bash

db-up:
	docker volume create device_inventory
	docker run --name device_inventory \
	-e POSTGRES_PASSWORD=password \
	-v device_inventory:/var/lib/postgresql/data \
	-d postgres 

db-sleep:
	sleep 3

db-create-inventory:
	docker exec device_inventory psql -U postgres -c "create database device_inventory"

db-create-table:
	docker exec device_inventory psql -U postgres -d device_inventory -c "CREATE TABLE IF NOT EXISTS devices(id SERIAL PRIMARY KEY, hostname TEXT UNIQUE, mgmt_address INET);"

db-all-up: db-up db-sleep db-create-inventory db-create-table

db-shell:
	docker exec -it device_inventory psql -U postgres -d device_inventory

db-down:
	docker container rm -f device_inventory
	docker volume rm device_inventory