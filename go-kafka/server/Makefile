topics-list:
	kafka-topics --bootstrap-server=localhost:9092 --list

topics-create:
	kafka-topics --bootstrap-server=localhost:9092 --topic=nolhello --create

consumer:
	kafka-console-consumer --bootstrap-server=localhost:9092 --topic=nolhello

producer:
	kafka-console-producer --bootstrap-server=localhost:9092 --topic=nolhello

consumer-group1:
	kafka-console-consumer --bootstrap-server=localhost:9092 --topic=nolhello --group=group1

consumer-group2:
	kafka-console-consumer --bootstrap-server=localhost:9092 --topic=nolhello --group=group2

consumer-include1:
	kafka-console-consumer --bootstrap-server=localhost:9092 --include="nolhello|pondhello" --group=group3

consumer-include2:
	kafka-console-consumer --bootstrap-server=localhost:9092 --include="DepositFundEvent|CloseAccountEvent|OpenAccountEvent|WithdrawFundEvent" --group=log

watch-partition:
	kafka-topics --bootstrap-server=localhost:9092 --topic=nolhello --describe

mysql:
	docker run --name mysqldb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:8.0

createdb:
	docker exec -it mysqldb mysql -uroot -psecret -e "CREATE DATABASE micro;"

open-account:
	{"ID": "krungthai123","AccountHolder": "nol","AccountType": 1,"Balance": 1000}


