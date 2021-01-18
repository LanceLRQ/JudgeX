module github.com/wejudge/wejudge-polygon

go 1.14

require (
	github.com/LanceLRQ/deer-common v0.0.0-20201113090955-db272c7977ec // indirect
	github.com/envoyproxy/go-control-plane v0.9.7 // indirect
	github.com/go-kratos/kratos v0.6.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2 // indirect
	github.com/google/wire v0.4.0
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0 // indirect
	gorm.io/driver/mysql v1.0.3 // indirect
	gorm.io/gorm v1.20.8
)

replace github.com/LanceLRQ/deer-common => ./pkg/deer-common
