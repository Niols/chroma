module github.com/Niols/chroma/cmd/chromad

go 1.16

require (
	github.com/Niols/chroma v0.0.0-00010101000000-000000000000
	github.com/alecthomas/kong v0.2.4
	github.com/alecthomas/kong-hcl v0.2.0
	github.com/gorilla/csrf v1.6.2
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
)

replace github.com/Niols/chroma => ../../
