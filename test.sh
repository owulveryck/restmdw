#!/bin/sh
echo "Trying to access hello without token"
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:5000/test/hello
echo ""
echo "Getting the token"
echo "-----------------"
echo ""
token=$(curl -s -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST  -d '{"Username":"test","Password":"test"}'  http://localhost:5000/token | egrep -o '[^"]{439}')
echo ""
echo "Testing with the token"
echo "----------------------"
echo ""
curl -i -H "Authorization: Bearer $token" http://localhost:5000/test/hello
echo ""
echo "Getting the offering list"
echo "-------------------------"
echo ""
curl -i -H "Authorization: Bearer $token" http://localhost:5000/offering
