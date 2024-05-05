APP_NAME = email-cleaner
VM_DIR = ~/Desktop/Dev/cloudVM

build:
	go mod download && go mod verify
	go mod build -o ${APP_NAME} main.go
	./${APP_NAME}

zip: 
	tar -zcvf ${VM_DIR}/${APP_NAME}.tar.gz .

clean: 
	go clean 
	go mod tidy 