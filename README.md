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

## API_SECRET

`API_SECRET` is used as the static API key for authenticating gRPC calls.

### Recommended ways to generate API_SECRET

```sh
openssl rand -base64 32
```

Copy the generated string and set it as `API_SECRET` in `config.yaml` (and keep it private).
