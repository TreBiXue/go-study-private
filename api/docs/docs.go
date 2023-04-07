// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/center": {
            "get": {
                "description": "Get center details by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "center_no",
                        "name": "center_no",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Center"
                        }
                    }
                }
            }
        },
        "/api/v1/nyuka/getcounts": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "summary": "获取未入荷条数",
                "parameters": [
                    {
                        "description": "test",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.NyukaInputRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/response.NyukaInputResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Center": {
            "type": "object",
            "properties": {
                "center_name": {
                    "type": "string"
                },
                "center_no": {
                    "type": "string"
                }
            }
        },
        "request.NyukaInputRequest": {
            "type": "object",
            "properties": {
                "center_no": {
                    "type": "string"
                },
                "nyuka_begin": {
                    "type": "string"
                },
                "nyuka_end": {
                    "type": "string"
                },
                "sire_no": {
                    "type": "string"
                }
            }
        },
        "response.NyukaInputResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "orderNo_List": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
