package ipctransitrouter

import (
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"testing"
)

//See tests of the reference implementation:
//https://github.com/dana/perl-IPC-Transit-Router/blob/master/t/basic.t
//https://github.com/dana/perl-Message-Router/blob/master/t/basic.t
//https://github.com/dana/perl-Message-Router/blob/master/t/hash_of_routes.t
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
