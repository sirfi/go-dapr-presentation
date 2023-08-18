module github.com/sirfi/go-dapr-presentation/4-go-dapr-grpc-service/order-api

go 1.20

require (
	github.com/dapr/go-sdk v1.8.0
	github.com/redis/go-redis/v9 v9.1.0
	github.com/sirfi/go-dapr-presentation/4-go-dapr-grpc-service/models v0.0.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.55.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/sirfi/go-dapr-presentation/4-go-dapr-grpc-service/models => ../models
