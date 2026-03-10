# OVS-AGENT

## Generate

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/ovsagentpb/agent.proto
```

or

```sh
protoc --go_out=. --go_opt=paths=source_relative `
    --go-grpc_out=. --go-grpc_opt=paths=source_relative `
    api/ovsagentpb/agent.proto
```
