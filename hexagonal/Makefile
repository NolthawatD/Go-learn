mysql:
	docker run --name mysqldb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:8.0

createdb:
	docker exec -it mysqldb mysql -uroot -psecret -e "CREATE DATABASE banking;"

dropdb:
	docker exec -it mysqldb mysql -uroot -psecret -e "DROP DATABASE banking;"

createAccount: 
	curl localhost:8000/customers/2003/accounts -i -X POST -H "content-type:application/json" -d \
	'{"account_type": "saving", "amount": 5000}'

getAccount: 
	curl localhost:8000/customers/2003/accounts -i 