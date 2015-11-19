package ipctransitrouter

import (
	"github.com/dana/go-ipc-transit"
	"github.com/dana/go-message-match"
	"github.com/dana/go-message-transform"
	"github.com/kr/pretty"
)

type TransitRouterError struct {
	What string
}

func (e TransitRouterError) Error() string {
	return pretty.Sprintf("%v", e.What)
}

func doForward(sendMessage map[string]interface{}, forward map[string]interface{}) error {
	if _, ok := forward["qname"]; !ok {
		return TransitRouterError{"missing required field qname in forward"}
	}
	//TODO need to validate the type of qname here
	qname := forward["qname"].(string)
	sendErr := ipctransit.Send(sendMessage, qname)
	if sendErr != nil {
		return TransitRouterError{"sendErr"}
	}
	return nil
}
func doRoute(sendMessage map[string]interface{}, route map[string]interface{}) (bool, error) {
	if _, ok := route["match"]; !ok {
		return false, TransitRouterError{"'match' attribute required in each route"}
	}
	//TODO need to validate the type of match here
	match := route["match"].(map[string]interface{})
	doesMatch, matchErr := messagematch.Match(sendMessage, match)
	if matchErr != nil {
		return false, TransitRouterError{"matchErr"}
	}
	if !doesMatch {
		return false, nil
	}
	if _, ok := route["forwards"]; !ok {
		return false, nil
	}
	//TODO need to validate the type of forwards here
	forwards := route["forwards"].([]interface{})
	if _, ok := route["transform"]; ok {
		//TODO need to validate the type of transform here
		transformErr := messagetransform.Transform(&sendMessage, route["transform"].(map[string]interface{}))
		if transformErr != nil {
			return false, TransitRouterError{"transformErr"}
		}
	}
	for _, forward := range forwards {
		//TODO need to validate the type of forward here
		forwardErr := doForward(sendMessage, forward.(map[string]interface{}))
		if forwardErr != nil {
			return false, TransitRouterError{"forwardErr"}
		}
	}
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
