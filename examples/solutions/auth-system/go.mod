module github.com/joinself/academy/examples/server/auth-system

go 1.24.2

require (
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/sessions v1.4.0
	github.com/joinself/self-go-sdk v0.59.0
)

require github.com/gorilla/securecookie v1.1.2 // indirect

replace github.com/joinself/academy/examples/server/common => ../common
