# Motion

Create, Share, Move with Motion.

## Deployment

```
bash deployment/build.sh
docker compose -f deployment/docker-compose.yaml up
```

## Development

Run server:

```
go run .
```

Run client:

```
pnpm install \
&& pnpm run dev:client 
```
