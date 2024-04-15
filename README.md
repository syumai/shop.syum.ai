# shop.syum.ai

* fork of [shop.ojisan.io](https://github.com/sadnessOjisan/shop.ojisan.io).

## Development

### Prerequisite

* buf
  - see: https://connectrpc.com/docs/go/getting-started/#install-tools

* Install buf-related tools

```
$ go install github.com/bufbuild/buf/cmd/buf@latest
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
```

### Generating code from Protobuf schema

* buf generate

