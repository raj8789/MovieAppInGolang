module run

go 1.20

require (
	github.com/gorilla/mux v1.8.0
	handler v0.0.0-00010101000000-000000000000
)

require model v0.0.0-00010101000000-000000000000 // indirect

replace handler => ../handler

replace model => ../model
