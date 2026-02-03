module github.com/poymanov/codemania-task-board/board

go 1.25.5

replace github.com/poymanov/codemania-task-board/shared => ../shared

require (
	github.com/caarlos0/env/v11 v11.3.1
	github.com/joho/godotenv v1.5.1
	github.com/poymanov/codemania-task-board/shared v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.78.0
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.3.0 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
