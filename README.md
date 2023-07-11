# Protobuffer Optional Test

- proto 파일에서 primitive type에 대해 optional로 처리하는 경우, generate 된 Go 파일에서 어떻게 보여지는지 확인
- nil 처리 가능한지 확인.
- grpc-gateway를 사용하여, client에서 default value(zero-value)를 던졌을 경우 해당 값(nil 이 아닌 값)으로 받는지 확인.

## Make Protobuffer file

```protobuf
message TestMessage {
  string string_value = 1;
  optional string optional_string_value = 2;

  int32 int32_value = 3;
  optional int32 optional_int32_value = 4;

  bool bool_value = 5;
  optional bool optional_bool_value = 6;
}

service YourService {
  rpc Echo(TestMessage) returns (TestMessage) {
    option (google.api.http) = {
      post: "/v1/example/echo"
      body: "*"
    };
  }
}
```

## Generate gRPC Stubs

```shell
protoc -I . \
    --go_out . --go_opt paths=source_relative \
    --go-grpc_out . --go-grpc_opt paths=source_relative \
    --go-grpc_opt=require_unimplemented_servers=false \
    api/my_service.proto
```

### generated golang

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

## Generate gRPC Gateway

```shell
protoc -I . --grpc-gateway_out . \    
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    api/my_service.proto
```

### Result directory

```plaintext
api
├── my_service.proto
└── v1
    ├── my_service.pb.go
    ├── my_service.pb.gw.go
    └── my_service_grpc.pb.go
```

## Test

- grpc server와 grpc proxy 서버를 설정 한 후 테스트 진행합니다.

### Case 1) default value 를 모두 요청 한 경우,

Request)

```json
{
  "string_value": "",
  "optional_string_value": "",
  "int32_value": 0,
  "optional_int32_value": 0,
  "bool_value": false,
  "optional_bool_value": false
}
```

response)

```json
{
  "stringValue": "",
  "optionalStringValue": "",
  "int32Value": 0,
  "optionalInt32Value": 0,
  "boolValue": false,
  "optionalBoolValue": false
}
```

in-grpc server log)

```
string_value: ""
optional_string_value: ""
int32_value: 0
optional_int32_value: 0
bool_value: false
optional_bool_value: false
```

### Case 2)  optional 필드들 빼고 요청

Request)

```json
{
  "string_value": "",
  // "optional_string_value": "",
  "int32_value": 0,
  // "optional_int32_value": 0,
  "bool_value": false
  // "optional_bool_value": false
}
```

response)

```json
{
  "stringValue": "",
  "int32Value": 0,
  "boolValue": false
}
```

in-grpc server log)

```
string_value: ""
optional_string_value: nil
int32_value: 0
optional_int32_value: nil
bool_value: false
optional_bool_value: nil
```

### Case 3)  비어있는 Body 요청

Request)

```json
{
  // "string_value": "",
  // "optional_string_value": "",
  // "int32_value": 0,
  // "optional_int32_value": 0,
  // "bool_value": false,
  // "optional_bool_value": false
}
```

response)

```json
{
  "stringValue": "",
  "int32Value": 0,
  "boolValue": false
}
```

in-grpc server log)

```
string_value: ""
optional_string_value: nil
int32_value: 0
optional_int32_value: nil
bool_value: false
optional_bool_value: nil
```

## 결과

- protobuf 의 optional 기능을 사용하여 ‘필드의 부재’와 ‘default value’ 사이의 차이를 알 수 있습니다.

## Ref

- Protocol Buffers Documentation : https://protobuf.dev/programming-guides/proto3/#field-labels
- proto3 문법의 optional label (Field presence) - Ukjae Jeong blog : https://blog.ukjae.io/posts/optional-label-in-proto3/
- grpc gateway : https://github.com/grpc-ecosystem/grpc-gateway
