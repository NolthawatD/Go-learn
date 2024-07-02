mysql:
	docker run --name mysqldb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:8.0

createdb:
	docker exec -it mysqldb mysql -uroot -psecret -e "CREATE DATABASE bank;"