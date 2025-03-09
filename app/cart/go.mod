module github.com/meilingluolingluo/gomall/app/cart

go 1.24.0

replace (
	github.com/meilingluolingluo/gomall/rpc_gen => ../../rpc_gen
)
require github.com/golang/protobuf v1.5.4 // indirect
