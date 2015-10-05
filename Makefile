all: token hello_unauth
install: repository/NodeTypes/*.yaml
	mkdir -p /tmp/restmdw/repository/NodeTypes	
	cp repository/NodeTypes/* /tmp/restmdw/repository/NodeTypes
run:
	go run server.go


hello_unauth:
	echo "Test Without token"
	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:5000/test/hello

offering:
	curl -i -H "Authorization: Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NDQzMTA0OTYsImlhdCI6MTQ0NDA1MTI5Niwic3ViIjoiIn0.nbla7CXfgmqW0xAdLVuGhWBSJdwdvEo7cGezKIK5TSnwbrl-x3AQsRgsATMYIdRVVCLDHUkXb4CwZDmk9J05cln-hrtaXI3O5lryJnVjEfAU7KGlLdNDm265_VGy3q7uER4xoIuoD59EO8hJ4YDrvSLwX9mZ3VXt1dnRApwCTt3AS-RyXWCbIGQtRegmMwRZ9MTYSsXxBB6nUIY9aQ6iLcbYl0zz9BanQtrgDft1zP5UcNZphd-y7Jk5lIpe7nA3Rle3w_sTDrdrLkZFKi-rw-EOA92Nr_pnyeLLGGakbgBg9YamJ80_64NWfXJ05RkNUXiZloHCKQYdvyUAryPIQQ" http://localhost:5000/offering 

tests/flags/token.json: core/authentication/*.go core/authentication/module/sample.go
	curl -s -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST  -d '{"Username":"test","Password":"test"}'  http://localhost:5000/token > tests/flags/token.json 
	cat tests/flags/token.json

token: tests/flags/token.json

