package ipctransitrouter

import (
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func Route(sendMessage map[string]interface{}, config map[string]interface{}) error {
func TestBasic(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "b",
	}
	config := map[string]interface{}{
		"routes": []interface{}{"hi", "there"},
	}
	err := Route(message, config)
	assert.Nil(err)
}

func routeTestNo() {
	pretty.Println("Yeah")
}
