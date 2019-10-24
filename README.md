# GRPC-Middleware-Example
This is an example usage of interceptor chaining in GRPC-Go.

1. Generate the pb file.
```
$ protoc -I releases releases.proto --go_out=plugins=grpc:releases
```
2. Start server
```
$ go run server/main.go
```
Console(server) shows the following message.
```
2019/10/23 16:23:57 Listening on  localhost:10000
```

3. Run Client
```
$ go run client/main.go
```

Console(client) shows the following logs
```
PS \GOPATH\src\GRPC-Middleware-Example> go run .\client\main.go
2019/10/23 16:24:04 A - start:  2019-10-23 16:24:04.913782 +0530 IST m=+0.011979601
2019/10/23 16:24:04 B - start:  2019-10-23 16:24:04.913782 +0530 IST m=+0.011979601
2019/10/23 16:24:04 B - Invoked RPC method=/releases.GoReleaseService/ListReleases; Duration=16.4963ms; Error=<nil>
2019/10/23 16:24:04 B - End:  2019-10-23 16:24:04.9302785 +0530 IST m=+0.028475901
2019/10/23 16:24:04 A - Invoked RPC method=/releases.GoReleaseService/ListReleases; Duration=17.0375ms; Error=<nil>
2019/10/23 16:24:04 A - End:  2019-10-23 16:24:04.9312815 +0530 IST m=+0.029478901
Version Release Date    Release Notes
1.1     21.10.2009      First release
1.13    22.10.2019      Latest release
```

and console(server) shows the following logs
```
2019/10/23 16:23:57 Listening on  localhost:10000
2019/10/23 16:24:04 A - start:  2019-10-23 16:24:04.927779 +0530 IST m=+7.782712601
2019/10/23 16:24:04 B - start:  2019-10-23 16:24:04.9282791 +0530 IST m=+7.783212701
2019/10/23 16:24:04 B - Request - Method:/releases.GoReleaseService/ListReleases        Duration:499.6µs        Error:<nil>
2019/10/23 16:24:04 B - End:  2019-10-23 16:24:04.9287787 +0530 IST m=+7.783712301
2019/10/23 16:24:04 A - Request - Method:/releases.GoReleaseService/ListReleases        Duration:1.5014ms       Error:<nil>
2019/10/23 16:24:04 A - End:  2019-10-23 16:24:04.929779 +0530 IST m=+7.784712501
```