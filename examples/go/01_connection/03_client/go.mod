module github.com/joinself/academy/examples/go/01_connection

go 1.24.2

replace github.com/joinself/academy/examples/go/common => ../../common

require (
	github.com/joinself/academy/examples/go/common v0.0.0-00010101000000-000000000000
	github.com/joinself/self-go-sdk v0.59.0
)
