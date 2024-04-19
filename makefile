run:
	sudo docker-compose up --build -d

swag:
	swag init -g internal/app/app.go

export-go:
	export GOPATH=$HOME/go
	export GOBIN=$GOPATH/bin
	export PATH=$PATH:$GOBIN
