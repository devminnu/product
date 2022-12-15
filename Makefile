build-docker:
	docker image build -t devminnu/product:latest .

run-docker: build-docker
	docker container run --name product -p 8080:8080 devminnu/product

build:
	go build -o app ./cmd/product

prune:
	docker container prune -f

run:
	docker-compose up --build

destroy:
	docker-compose down && make prune

re-run:
	make destroy && make run

proto:
	protoc \
	--go_out=./ \
	--go_opt=paths=import \
	--go-grpc_out=./ \
	--go-grpc_opt=paths=import \
	./api/grpc/proto/*/*.proto

# grpcurl -plaintext 54.237.88.89:50051 describe
grpcurl -plaintext localhost:50051 describe
grpcurl -plaintext 18.204.220.190:50051 describe
grpcurl -plaintext product-service-alb-1796865760.us-east-1.elb.amazonaws.com:50051 describe