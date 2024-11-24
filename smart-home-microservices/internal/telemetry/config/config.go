package config

type Config struct {
	Secret               string
	PostgresDNS          string
	KafkaDNSs            []string
	KafkaConsumerGroupID string
	KafkaClientID        string
}

func DefaultConfig() Config {
	return Config{
		Secret:               "secret",
		PostgresDNS:          "postgres://postgres:postgres@localhost:5432/sh_devices",
		KafkaDNSs:            []string{"localhost:9092"},
		KafkaConsumerGroupID: "devices-consumer-group",
	}
}
