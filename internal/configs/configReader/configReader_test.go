package configreader_test

import (
	configreader "JHETBackend/internal/configs/configReader"
	"fmt"
	"testing"
)

func TestConfigRead(t *testing.T) {
	configreader.Init()
	configreader.GetConfig()
	fmt.Printf("%v", configreader.GetConfig())
}
