package configreader_test

import (
	configreader "JHETBackend/configs/configReader"
	"log"
	"testing"
)

func TestConfigRead(t *testing.T) {
	log.Printf("%v", configreader.GetConfig())
	log.Printf("%v", configreader.GetConfig().Database)
	log.Printf("%v", configreader.GetConfig().Database.Host)
}
