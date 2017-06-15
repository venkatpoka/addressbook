# addressbook
Packages used:
"github.com/gorilla/mux"
"github.com/mattn/go-sqlite3"

Use Cases:
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:12345/addressBook/2
curl: (52) Empty reply from server
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:12345/addressBook/1
curl: (52) Empty reply from server
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:12345/addressBook/1
curl: (52) Empty reply from server
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:12345/addressBook/1
HTTP/1.1 200 OK
Date: Fri, 09 Jun 2017 20:43:11 GMT
Content-Length: 70
Content-Type: text/plain; charset=utf-8
 
{"ID":"1","Firstname":"Nic","Lastname":"Raboy","Email":"","Phone":""}

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:12345/addressBook/1
HTTP/1.1 200 OK
Date: Fri, 09 Jun 2017 20:54:11 GMT
Content-Length: 70
Content-Type: text/plain; charset=utf-8
 
{"ID":"1","Firstname":"Nic","Lastname":"Raboy","Email":"","Phone":""}

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:12345/addressBook/2
HTTP/1.1 200 OK
Date: Fri, 09 Jun 2017 20:54:13 GMT
Content-Length: 61
Content-Type: text/plain; charset=utf-8
 
{"ID":"","Firstname":"","Lastname":"","Email":"","Phone":""}

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST http://localhost:12345/addressBook/3 --data '{"ID":"3","Firstname":"Andrew"}'
curl: (52) Empty reply from server

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST http://localhost:12345/addressBook/3 --data '{"ID":"3","Firstname":"Andrew"}'
HTTP/1.1 200 OK
Date: Fri, 09 Jun 2017 20:55:22 GMT
Content-Length: 68
Content-Type: text/plain; charset=utf-8
 
{"ID":"3","Firstname":"Andrew","Lastname":"","Email":"","Phone":""}

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:12345/addressBook/3
HTTP/1.1 200 OK
Date: Fri, 09 Jun 2017 20:55:28 GMT
Content-Length: 68
Content-Type: text/plain; charset=utf-8
 
{"ID":"3","Firstname":"Andrew","Lastname":"","Email":"","Phone":""}