package gen

//
////go:generate mkdir -p pb
////work
////go:generate protoc --go_out=. --go-grpc_out=paths=source_relative outliers.proto
////work
//go:generate protoc --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative outliers.proto
//
//go:generate protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative outliers.proto
