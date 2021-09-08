curl -X GET http://127.0.0.1:9091/v1/account

curl -X POST http://127.0.0.1:9091/v1/account -H 'content-type: application/json' -d '{ "accountName": "peterydd", "password": "Ydd@19901203" }'

curl -X GET http://127.0.0.1:9091/v1/account/peterydd

curl -X PUT http://127.0.0.1:9091/v1/account/ -H 'content-type: application/json' -d '{ "accountName": "peterydd", "password": "wewdxw@1" }'

curl -X DELETE http://127.0.0.1:9091/v1/account/1938cc8c-5287-3cad-1e4b-f50fc581d858


curl -X POST http://127.0.0.1:9091/v1/account/login -H 'content-type: application/json' -d '{ "accountName": "peterydd", "password": "Ydd@19901203" }'


curl -X POST http://127.0.0.1:9091/v1/account -H 'content-type: application/json' -d '{ "accountName": "xwsoft", "password": "Ydd@19901203" }'
