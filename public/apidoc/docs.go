
package main
//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "",
    "swaggerVersion": "1.2",
    "basePath": "{{.}}",
    "apis": [
        {
            "path": "/hello",
            "description": "Says hello"
        }
    ],
    "info": {}
}`
var apiDescriptionsJson = map[string]string{"hello":`{
    "apiVersion": "",
    "swaggerVersion": "1.2",
    "basePath": "{{.}}",
    "resourcePath": "/hello",
    "apis": [
        {
            "path": "/hello",
            "description": "Says hello",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "HelloController",
                    "type": "",
                    "items": {},
                    "summary": "Says hello",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "customer_id",
                            "description": "Customer ID",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "order_id",
                            "description": "Retrieve order with given ID only",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "order_nr",
                            "description": "Retrieve order with given number only",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "created_from",
                            "description": "Date-time string, MySQL format. If specified, API will retrieve orders that were created starting from created_from",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "created_to",
                            "description": "Date-time string, MySQL format. If specified, API will retrieve orders that were created before created_to",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 400,
                            "message": "Customer ID must be specified",
                            "responseType": "object",
                            "responseModel": "github.com.owulveryck.restmdw.web.JsonErr"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.owulveryck.restmdw.web.JsonErr": {
            "id": "github.com.owulveryck.restmdw.web.JsonErr",
            "properties": {
                "code": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "text": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,}
