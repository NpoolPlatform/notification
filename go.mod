module github.com/NpoolPlatform/notification

go 1.16

require (
	entgo.io/ent v0.10.1
	github.com/NpoolPlatform/api-manager v0.0.0-20220205130236-69d286e72dba
	github.com/NpoolPlatform/go-service-framework v0.0.0-20220120091626-4e8035637592
	github.com/NpoolPlatform/message v0.0.0-20220321085959-60d96a8d8fe1
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.8.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.1-0.20210427113832-6241f9ab9942
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.45.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
