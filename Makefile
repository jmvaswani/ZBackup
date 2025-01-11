
SERVERNAME = ${SERVER_NAME}  

proto-gen:
	protoc --go_out=./client/pkg/ --go_opt=paths=source_relative --go-grpc_out=./client/pkg --go-grpc_opt=paths=source_relative proto/grpc.proto 
	protoc --go_out=./server/pkg/ --go_opt=paths=source_relative --go-grpc_out=./server/pkg --go-grpc_opt=paths=source_relative proto/grpc.proto

# This target will build a new set of client and server certificates using certs/CA.crt and certs/CA.key
.PHONY: certs
certs:
	openssl req -newkey rsa:2048 -keyout certs/server.key -out certs/server.csr -days 365 -nodes
	openssl x509 -req -CA certs/CA.crt -CAkey certs/CA.key -in certs/server.csr -out certs/server.crt -days 365 -CAcreateserial -extfile <(printf "subjectAltName= DNS:${SERVERNAME}")

	openssl req -newkey rsa:2048 -keyout certs/client.key -out certs/client.csr -days 365 -nodes
	openssl x509 -req -CA certs/CA.crt -CAkey certs/CA.key -in certs/client.csr -out certs/client.crt -days 365 -CAcreateserial
	
	
#Generates a new CA.crt and CA.key
ca-cert:
	openssl req -x509 -sha256 -days 5000 -newkey rsa:2048 -keyout certs/CA.key -out certs/CA.crt -nodes