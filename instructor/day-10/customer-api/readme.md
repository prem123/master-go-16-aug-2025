## Running mysql db container with default tables and data

```sh
docker run -d --rm -p 3306:3306 --name customers-db -e MYSQL_ROOT_PASSWORD=thecodecamp thecodecamp/customers-db
```
## MySql connect and view tables from cli

```bash
docker exec -it customers-db bash

mysql -u root -p banking
```


## Mockery

### Installation
https://vektra.github.io/mockery/v3.5/installation/

```bash
brew install mockery
```