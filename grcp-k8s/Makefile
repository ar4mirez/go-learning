build:
	protoc -I . --go_out=plugins=grpc:. ./pb/*.proto

gcd:
	docker build -t local/gcd -f gcd/Dockerfile .

api:
	docker build -t local/api -f api/Dockerfile .
