# todotxt-frontend

## Usage

```bash
cd frontend && yarn install && yarn build

docker build -t todotxt-frontend .
docker run \
  -p 3000:3000 \
  -v $(pwd)/todotxt:/opt/todotxt \
  -e TODO_PATH=/opt/todotxt/todo.txt \
  -e LISTEN_ADDR=:3000 \
  todotxt-frontend
```
