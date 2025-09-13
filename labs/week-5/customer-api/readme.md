## Running mysql db container with default tables and data

```sh
docker run --rm -p 3306:3306 --name customers-db -e MYSQL_ROOT_PASSWORD=thecodecamp thecodecamp/customers-db
```

