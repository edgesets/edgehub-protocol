package protocol

//go:generate protoc -I . -I ./third_party --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --go-http_out=paths=source_relative:. ./shared/*.proto ./proxy/v1/*.proto ./envelope/v1/*.proto ./ssapi/v1/*.proto
