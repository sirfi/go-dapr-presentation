module github.com/sirfi/go-dapr-presentation/2-go-dapr-service/invoice-api

go 1.20

require (
	github.com/redis/go-redis/v9 v9.0.3
	github.com/sirfi/go-dapr-presentation/2-go-dapr-service/models v0.0.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)

require (
	github.com/dapr/go-sdk v1.7.0
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230104163317-caabf589fcbf // indirect
	google.golang.org/grpc v1.51.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/sirfi/go-dapr-presentation/2-go-dapr-service/models => ../models
