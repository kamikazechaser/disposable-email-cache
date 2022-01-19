## disposable-email-cache

> A self-hosted alternative to Kickbox's OpenAPI for checking if an email is disposable/temporary.

### Build and run

```bash
# single binary
$ make build
$ ./disposable-email-cache
```

or

```bash
# build Docker container
$ DOCKER_BUILDKIT=1 docker build -t disposable-email-cache .
$ docker run --name disposable-email-cache -p 5000:5000 disposable-email-cache:latest

# or use pre-built Docker image
$  docker run --name disposable-email-cache -p 5000:5000 kamikazechaser/disposable-email-cache:latest
```

### API Endpoints

**Check if email domain is disposable**

```bash
$ curl localhost:5000/check/n8.gs
# {"disposable":true}
```

**Update and reload cache**

```bash
$ curl localhost:5000/update-cache
# {"cacheUpdated":true}
```
