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

consumer-include:
	kafka-console-consumer --bootstrap-server=localhost:9092 --include="nolhello|pondhello" --group=group3

watch-partition:
	kafka-topics --bootstrap-server=localhost:9092 --topic=nolhello --describe