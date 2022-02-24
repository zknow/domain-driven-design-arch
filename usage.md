# docker cmd

Builds the images if the images do not exist and starts the containers:
```
docker-compose up
```

It is forced to build the images even when not needed
```
docker-compose up --build -d
```

Recreate containers even if their configuration and image haven't changed.
```
docker-compose up --force-recreate -d
```

Clean up multiple types of objects at once
```
docker 'object' prune
```

# go test cmd

testing *_test.go file each folder
```
go test -v ./...
```