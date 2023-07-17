### Build the image:
```
docker build -t accunox:1.0 .
```

### Run the container
```
docker run -p 10000:10000 accunox:1.0
```

### Pull image off docker hub
```
docker pull harshil18/accuknox:1.0
```

### Signup
```
[POST]
localhost:10000/signup
```

```json
{
"email": "h@g.com",
"name": "harshil",
"password": "pass123"
}
```

### Login
```
[POST]
localhost:10000/login
```

```json
{
  "email": "h@g.com",
  "password": "pass123"
}
```
Response
```json
{
    "s_id": "h@g.comloggedIn"
}
```

### Create Note
```
[POST]
localhost:10000/note
```

```json
{
  "s_id":"h@g.comloggedIn",
  "note":"test note1"
}
```

### Get Notes
```
[GET]
localhost:10000/note
```
Response
```json
{
  "notes": [
    {
      "id": 1,
      "note": "test note1"
    }
  ]
}
```

### Delete Note
```
[DELETE]
localhost:10000/note
```
```json
{
    "s_id":"h@g.comloggedIn",
    "id":2
}
```

