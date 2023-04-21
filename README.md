# go-studying


DI:
cd cmd
~/go/bin/wire

ls
openapi:
cd api
~/go/bin/swag init -g ./handler/center_handler.go -o ./docs -parseDependency
http://localhost:8080/swagger/index.html#/default/