# resource-keeper-api

## sign-up

### Request

* `curl -k -X POST https://46.101.253.18:8000/auth/sign-up -d '{"name":"Lika", "email":"lika@gmail.com", "password":"LikaPass"}'`

### Responses

* `{"ok":1}`
* `{"message":"invalid input body"}`
* `{"message":"invalid input value"}`

## sign-in

### Request

* `curl -k -X POST https://46.101.253.18:8000/auth/sign-in -d '{"email":"lika@gmail.com", "password":"LikaPass"}'`

### Responses

* `{"token":"really_long_token"}`
* `{"message":"invalid input body"}`
* `{"message":"invalid input value"}`

## resource (create, read, update, delete)

### Request

* `curl -k -X POST https://46.101.253.18:8000/resource/ -H "Authorization: Bearer really_long_token" -d '{"resource_name":"Instagram", "resource_login":"lika@gmail.com", "resource_password":"LikaInstaPass"}'`

### Responses

* `{"id":1}`
* `{"message":"invalid input body"}`
* `{"message":"invalid input value"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`

### Request

* `curl -k -X GET https://46.101.253.18:8000/resource/ -H "Authorization: Bearer really_long_token"`

### Responses

* `{"data":[{"id":1,"user_id":1,"resource_name":"Instagram","resource_login":"lika@gmail.com","resource_password":"LikaInstaPass","date_creation":"2022-11-04T10:48:30.762652Z","last_update":"2022-11-04T10:48:30.762652Z"}]}`
* `{"message":"signature is invalid"}`
* `{"message":"user is not found"}`
* `{"message":"user id is of invalid type"}`

### Request

* `curl -k -X GET https://46.101.253.18:8000/resource/1 -H "Authorization: Bearer really_long_token"`

### Responses

* `{"data":{"id":1,"user_id":1,"resource_name":"Instagram","resource_login":"lika@gmail.com","resource_password":"LikaInstaPass","date_creation":"2022-11-04T10:48:30.762652Z","last_update":"2022-11-04T10:48:30.762652Z"}}`
* `{"message":"invalid id param"}`
* `{"message":"resource id not found"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`

### Request

* `curl -k -X PUT https://46.101.253.18:8000/resource/1 -H "Authorization: Bearer really_long_token" -d '{"resource_name":"NewResName", "resource_login":"NewResLogin", "resource_password":"NewResPass"}'`

### Responses

* `{"status":"ok"}`
* `{"message":"invalid id param"}`
* `{"message":"resource id not found"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`

### Request 

* `curl -k -X DELETE https://46.101.253.18:8000/resource/1 -H "Authorization: Bearer really_long_token"`

### Responses

* `{"id":1, "status":"ok"}`
* `{"message":"invalid id param"}`
* `{"message":"resource id not found"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`
