package ipctransitrouter

import (
	"github.com/kr/pretty"
)

type TransitRouterError struct {
	What string
}

func (e TransitRouterError) Error() string {
	return pretty.Sprintf("%v", e.What)
}

func doRoute(sendMessage map[string]interface{}, route map[string]interface{}) (bool, error) {
	pretty.Println(route)
	return false, nil
}

//See reference implementation:
//Combination of:
//https://github.com/dana/perl-Message-Router
//https://github.com/dana/perl-IPC-Transit-Router
func Route(sendMessage map[string]interface{}, config map[string]interface{}) error {
	if _, ok := config["routes"]; !ok {
		return TransitRouterError{"'routes' attribute required in the config"}
	}
	routes := config["routes"]
	switch routes.(type) {
	case []interface{}:
		for _, route := range routes.([]interface{}) {
			doShortCircuit, routeErr := doRoute(sendMessage, route.(map[string]interface{}))
			if routeErr != nil {
				return TransitRouterError{"route error"}
			}
			if doShortCircuit {
				break
			}
			pretty.Println(doShortCircuit)
		}
	case map[string]interface{}:
		panic("map routes type unimplemented")
	default:
		return TransitRouterError{"'routes' attribute must be either a map or array"}
	}
	return nil
}
func routeNo() {
	pretty.Println("Yeah")
}
