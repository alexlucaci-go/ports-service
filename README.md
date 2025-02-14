# Ports service

### Build and run (ports.json should be in root directory to run)

```bash
$ go build -o ports-service cmd/ports-service/main.go
$ ./ports-service
```

### Build and run with docker (ports.json should be in root directory to run)
```bash
$ bash docker-build.sh
$ docker run -d -p 8000:8000 --read-only --name ports-service local/ports-service
```

### How to test

```bash
$ go test ./...
```

### Adding port
```bash
$ curl -X POST "http://localhost:8000/v1/ports" \
     -H "Content-Type: application/json" \
     -d '{
           "id": "TEST_PORT_ID",
           "name": "Ajman",
           "city": "Ajman",
           "country": "United Arab Emirates",
           "alias": [],
           "regions": [],
           "coordinates": [55.5136433, 25.4052165],
           "province": "Ajman",
           "timezone": "Asia/Dubai",
           "unlocs": ["AEAJM"],
           "code": "52000"
         }'
```

### Updating port (partial fields)
```bash
$ curl -X PATCH "http://localhost:8000/v1/ports/TEST_PORT_ID" \
     -H "Content-Type: application/json" \
     -d '{
           "code": "53000"
         }'
```