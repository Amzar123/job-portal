// Code generated by swaggo/swag. DO NOT EDIT.

package apidocs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/jobs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Get a company",
                "operationId": "job.Handler.GetJobs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by \u003cfield\u003e with \u003cop\u003e is operator",
                        "name": "field_op",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Get related resources",
                        "name": "with",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort by \u003cfield\u003e in ascending order, add dash (-) in front of \u003cfield\u003e to sort in descending order",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Skip the first \u003cskip\u003e records",
                        "name": "skip",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit result to \u003climit\u003e records",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Fiber Example API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}