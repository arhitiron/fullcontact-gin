package core
import (
	"log"
	"gopkg.in/gcfg.v1"
)

var Cfg Config

type Config struct {
	Global   struct {
				 Port string
			 }
	Database struct {
				 DriverName string
				 Database   string
				 UserName   string
				 Password   string
				 Host       string
				 Charset    string
				 Engine     string
				 Encoding   string
			 }
	Kafka    struct {
				 Brokers      string
				 DefaultTopic string
			 }
}

func InitCfg() {
	err := gcfg.ReadFileInto(&Cfg, "./config.gcfg")
	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
	}
}