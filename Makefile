export ROOT_MOD=github.com/meilingluolingluo/gomall
.PHONY: gen-demo-proto

##@ Open Browser

.PHONY: open.gomall
open-gomall: ## open `gomall` website in the default browser
	@open "http://localhost:8080/"

.PHONY: open.consul
open-consul: ## open `consul ui` in the default browser
	@open "http://localhost:8500/ui/"

.PHONY: open.jaeger
open-jaeger: ## open `jaeger ui` in the default browser
	@open "http://localhost:16686/search"

.PHONY: open.prometheus
open-prometheus: ## open `prometheus ui` in the default browser
	@open "http://localhost:9090"

.PHONY: open.grafana
open-grafana: ## open `prometheus ui` in the default browser
	@open "http://localhost:3000"

#	cwgo -h
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --module github.com/meilingluolingluo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --module  github.com/meilingluolingluo/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: demo-link-fix
demo-link-fix:
	cd demo/demo_proto && golangci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m
.PHONY:gen-home
gen-home:
	@cd app/home && cwgo server -I ../../idl --type HTTP --service frontend --module github.com/meilingluolingluo/gomall/app/frontend --idl ../../idl/frontend/home.proto

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/order_page.proto --service frontend --module github.com/meilingluolingluo/gomall/app/frontend -I ../../idl

.PHONY: gen-rpc
gen-rpc:
	@cd app/rpc && cwgo client --type RPC --service user --module github.com/meilingluolingluo/gomall/rpc_gen --I ../idl --idl ../idl/user.proto
.PHONY: gen-user
gen-user: 
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/meilingluolingluo/gomall/rpc_gen  -I ../idl  --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module github.com/meilingluolingluo/gomall/app/user --pass "-use github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/user.proto

.PHONY: gen-product
gen-product: 
	@cd rpc_gen && cwgo client --type RPC --service product --module github.com/meilingluolingluo/gomall/rpc_gen  -I ../idl  --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module github.com/meilingluolingluo/gomall/app/product --pass "-use github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/product.proto


.PHONY: gen-cart
gen-cart: 
	@cd rpc_gen && cwgo client --type RPC --service cart --module github.com/meilingluolingluo/gomall/rpc_gen  -I ../idl  --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module github.com/meilingluolingluo/gomall/app/cart --pass "-use github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/cart.proto
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/cart_page.proto --service frontend --module github.com/meilingluolingluo/gomall/app/frontend -I ../../idl
.PHONY: gen-payment
gen-payment: 
	@cd rpc_gen && cwgo client --type RPC --service payment --module github.com/meilingluolingluo/gomall/rpc_gen  -I ../idl  --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module github.com/meilingluolingluo/gomall/app/payment --pass "-use github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/payment.proto

.PHONY: gen-checkout
gen-checkout: 
	@cd rpc_gen && cwgo client --type RPC --service checkout --module github.com/meilingluolingluo/gomall/rpc_gen  -I ../idl  --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module github.com/meilingluolingluo/gomall/app/checkout --pass "-use github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/checkout.proto
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/checkout.proto --service frontend --module github.com/meilingluolingluo/gomall/app/frontend -I ../../idl

.PHONY: gen-order
gen-order: 
	@cd rpc_gen && cwgo client --type RPC --service order --module github.com/meilingluolingluo/gomall/rpc_gen  -I ../idl  --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module github.com/meilingluolingluo/gomall/app/order --pass "-use github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/order.proto


.PHONY: gen-auth
gen-auth:
	@cd  rpc_gen && cwgo client --type RPC --service auth --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/auth.proto
	@cd  app/auth && cwgo server --type RPC --service auth --module ${ROOT_MOD}/app/auth --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/auth.proto

