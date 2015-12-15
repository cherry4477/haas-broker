package catalog

import (
	"fmt"
	"net/http"
)

const (
	HandlerPath = "/v2/catalog"
)

func Get() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		response := `{
			"services": [{
				"id": "5a9b9f22-a08d-11e5-8062-7831c1d4f660",
				"name": "pez-haas",
				"description": "Lease on-demand hardware as a service",
				"metadata":{
            "displayName":"PEZ-HaaS",
            "imageUrl":"https://pez.pezapp.io/img/pez-clear.png",
						"providerDisplayName":"PEZ",
						"documentationUrl":"https://www.pezapp.io",
            "supportUrl":"https://www.pezapp.io"
         },
				"bindable": false,
				"plans": [{
					"id": "6a977311-a08d-11e5-8062-7831c1d4f660",
					"name": "m1.small",
					"description": "A small instance of hardware as a service",
					"metadata":{
						"bullets":[
							 "48gb Mem", 
							 "Supermicro", 
							 "2.7ghz X5650 2 socket", 
							 "24 core",
							 "10 x 2TB disk sata"
						]
					}
				}],
				"dashboard_client": {
					"id": "pez-haas-client",
          "secret": "pez-haas-secret",
					"redirect_uri": "https://www.pezapp.io"
				}
			}]
		}`
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, response)
	}
}