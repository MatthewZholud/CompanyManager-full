up:
		docker-compose up
migration:
		cd migrator && docker build -t migrator . && docker run --network host migrator -path=/migrations/ -database "postgresql://postgres:mypassword@localhost:5432/company_manager?sslmode=disable" up
run_api:
		cd cmd/api && go run main.go
run_company:
		cd cmd/company && go run main.go
run_employee:
		cd cmd/employee && go run main.go