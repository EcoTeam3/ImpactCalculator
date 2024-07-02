CURRENT_DIR=$(shell pwd)

proto-gen:
	./script/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://postgers:parol@localhost:5432/impactcalculator?sslmode=disable'

mig-up:
	migrate -path storage/migrations -database 'postgres://postgers:parol@localhost:5432/impactcalculator?sslmode=disable' -verbose up

mig-down:
	migrate -path storage/migrations -database &{DBURL} -verbose down

mig-create:
	migrate create -ext sql -dir storage/migrations -seq create_tables

mig-insert:
	migrate create -ext sql -dir storage/migrations -seq insert_table