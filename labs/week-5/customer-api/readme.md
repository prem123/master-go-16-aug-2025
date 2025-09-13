## Running mysql db container with default tables and data

```sh
docker run -d --rm -p 3306:3306 --name customers-db -e MYSQL_ROOT_PASSWORD=thecodecamp thecodecamp/customers-db
```

## Mockery

### Installation
https://vektra.github.io/mockery/v3.5/installation/

```bash
brew install mockery
```