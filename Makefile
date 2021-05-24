run:
	protoc --proto_path=./pbfiles --proto_path=. --go_out=plugins=grpc:. ./pbfiles/Prod.proto
	protoc --proto_path=./pbfiles --proto_path=. --grpc-gateway_out=logtostderr=true:. ./pbfiles/Prod.proto
	protoc --proto_path=./pbfiles --proto_path=. --openapiv2_out=logtostderr=true:. ./pbfiles/Prod.proto