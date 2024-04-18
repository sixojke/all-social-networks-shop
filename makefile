run:
	sudo docker-compose up --build -d

swag:
	swag init -g internal/app/app.go