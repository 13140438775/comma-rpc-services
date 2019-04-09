# comma-rpc-services
Comma public services include（schedule、user、course）

## rpc services sh
```
    span := opentracing.StartSpan("service." + name)
    defer func() {
      if err != nil {
        span.SetTag("error", err)
      }
      span.Finish()
    }()
    ctx := opentracing.ContextWithSpan(ctx, span)


# build schedule.thrift file (thrift 0.11.0)
thrift -out . -r --gen go schedule.thrift

# rpc Server
go run apps/schedule/schedule.go

# build rpc server
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build apps/schedule/schedule.go

# upload schedule dev env
scp -i /Users/pb/Documents/51.pem schedule root@120.79.101.51:/export/apps/comma-rpc-services
```