install package sarama
https://pkg.go.dev/gopkg.in/Shopify/sarama.v1

install viper
https://pkg.go.dev/github.com/spf13/viper

check partition and insert to partition comsumer
watch-partition:
	kafka-topics --bootstrap-server=localhost:9092 --topic=nolhello --describe

refer to other module by write reference
replace events => ../events //refer to module events
go get events