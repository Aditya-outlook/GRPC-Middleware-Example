# GRPC-Middleware-Example
This is an example usage of interceptor chaining in GRPC-Go.

1. Generate the pb file.
#protoc -I releases releases.proto --go_out=plugins=grpc:releases

2. Start server
# go run server/main.go

3. Run Client
# go run client/main.go

4. FAQs
Why Middleware?
What are interceptors?
What is the flow of execution during chaining?
