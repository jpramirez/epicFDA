

echo "Generating FoodEvent Struct"
protoc --proto_path=api/proto/v1/ --proto_path=api/proto/v1/ --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1/ FoodEnforcement.proto
protoc --proto_path=api/proto/v1/ --proto_path=api/proto/v1/ --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1/ FoodEnforcement.proto



echo "Generating FoodEvent Struct"
protoc --proto_path=api/proto/v1/ --proto_path=api/proto/v1/ --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1/ FoodEvent.proto
protoc --proto_path=api/proto/v1/ --proto_path=api/proto/v1/ --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1/ FoodEvent.proto


