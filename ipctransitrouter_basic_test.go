package ipctransitrouter

//totally totally look at this:
//https://godoc.org/github.com/stretchr/stew/objects
//up vote
//OR you can use objects.Map from our stew package, it gives you dot accessors for maps:
//
//objects.Map(data).Get("service.auth.token")
//see http://godoc.org/github.com/stretchr/stew/objects
//via http://stackoverflow.com/questions/17056044/golang-quickly-access-data-of-maps-within-maps
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
	f1 := map[string]interface{}{
		"qname": "test",
	}
	r1 := map[string]interface{}{
		"match": map[string]interface{}{
			"a": "b",
		},
		"transform": map[string]interface{}{
			"x": "y",
		},
		"forwards": []interface{}{f1},
	}
	config := map[string]interface{}{
		"routes": []interface{}{r1},
	}
	err := Route(message, config)
	assert.Nil(err)
}

func routeTestNo() {
	pretty.Println("Yeah")
}
