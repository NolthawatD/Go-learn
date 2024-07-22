package main

import (
	"consumer/repositories"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/Shopify/sarama.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	dial := mysql.Open(dsn)

	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {

	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_0 // specify appropriate version
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	db := initDatabase()
	accountRepo := repositories.NewAccountRepository(db)
	accountEventHandler := services.NewAccountEventHandler(accountRepo)
	accountConsumerHandler := services.NewConsumerHandler(accountEventHandler)

	fmt.Println("Account consumer started...")
	for {
		consumer.Consume(context.Background(), events.Topics, accountConsumerHandler)
	}
}

// func main() {

// 	servers := []string{"localhost:9092"}

// 	consumer, err := sarama.NewConsumer(servers, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer consumer.Close()

// 	partitionConsumer, err := consumer.ConsumePartition("nolhello", 0, sarama.OffsetNewest)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer partitionConsumer.Close()

// 	fmt.Println("Consumer start.")
// 	for {
// 		select {
// 		case err := <-partitionConsumer.Errors():
// 			fmt.Println(err)
// 		case msg := <-partitionConsumer.Messages():
// 			fmt.Println(string(msg.Value))
// 		}
// 	}

// }
