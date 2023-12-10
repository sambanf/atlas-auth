# Atlas Auth
## _Auth API with Go Language_

## Docker
By default, the Docker will expose port 8080, so change this within the
Dockerfile if necessary. When ready, simply use the Dockerfile to
build the image.

```sh
docker build -t atlas-auth:0.0.1 .
```

Once done, run the Docker image and map the port to whatever you wish on
your host. In this example, we simply map port 8000 of the host to
port 8080 of the Docker (or whatever port was exposed in the Dockerfile):

```sh
docker run -d -p 8080:8080  atlas-auth:0.0.1
```

Verify the deployment by navigating to your server address in
your preferred browser.

```sh
127.0.0.1:8000/auth
```

## License

MIT
