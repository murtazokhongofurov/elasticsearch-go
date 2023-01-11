# elasticsearch-go
  
# RUN 
``` make run ```

# Test qilish
**POST**

```
$ curl --request POST 'http://localhost:8000/api/v1/post' --data-raw '{
"title":"post of title",
"description":"yangi description"
}' 
```

**GET**

```
$ curl --request GET 'http://localhost:8000/api/v1/post/{id}'
```

**DELETE**

```
$ curl --request DELETE 'http://localhost:8000/api/v1/post/{id}'
```

**PATCH**

```
$ curl --request PATCH 'http://localhost:8000/api/v1/post/{id}' --data-raw '{
"title":"update title",
"description":"update dexcription"
}'
```
