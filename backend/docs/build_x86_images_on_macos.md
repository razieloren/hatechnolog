- `docker buildx create --use` (Lets us use `buildx`, creates a container for that).
- Supported: `docker buildx ls`
- Now, you can for instance build an image for amd64 linux:
```bash
docker buildx build --platform linux/amd64 --push -t tag .
```