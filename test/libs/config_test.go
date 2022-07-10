package libs

import (
	"log"
	"subdomain/libs"
	"testing"
)

func TestConfig(t *testing.T) {
	var options libs.Options
	err := libs.InitConfig(options)
	if err != nil {
		log.Println(err)
	}
}
