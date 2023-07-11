# Protobuffer Optional Test

- proto 파일에서 primitive type에 대해 optional로 처리하는 경우, generate 된 Go 파일에서 어떻게 보여지는지 확인
- nil 처리 가능한지 확인.
- grpc-gateway를 사용하여, client에서 default value(zero-value)를 던졌을 경우 해당 값(nil 이 아닌 값)으로 받는지 확인.

## Generate gRPC Stubs

```shell
protoc -I . \
    --go_out . --go_opt paths=source_relative \
    --go-grpc_out . --go-grpc_opt paths=source_relative \
    --go-grpc_opt=require_unimplemented_servers=false \
    api/my_service.proto
```

### Generated Message
```go
type TestMessage struct {
	StringValue         string  `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	OptionalStringValue *string `protobuf:"bytes,2,opt,name=optional_string_value,json=optionalStringValue,proto3,oneof" json:"optional_string_value,omitempty"`
	Int32Value          int32   `protobuf:"varint,3,opt,name=int32_value,json=int32Value,proto3" json:"int32_value,omitempty"`
	OptionalInt32Value  *int32  `protobuf:"varint,4,opt,name=optional_int32_value,json=optionalInt32Value,proto3,oneof" json:"optional_int32_value,omitempty"`
	BoolValue           bool    `protobuf:"varint,5,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	OptionalBoolValue   *bool   `protobuf:"varint,6,opt,name=optional_bool_value,json=optionalBoolValue,proto3,oneof" json:"optional_bool_value,omitempty"`
}
```

### Generate grpc gateway

```shell
protoc -I . --grpc-gateway_out . \    
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    api/my_service.proto
```