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
        "/api/v1/login": {
            "get": {
                "description": "ユーザー登録",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.LoginResponse"
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
        },
        "/api/v1/nyuka/getjaninfo": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "summary": "JANによって検品情報取得",
                "parameters": [
                    {
                        "description": "jan info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.NyukaInputJANRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/response.NyukaInputJANResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/nyuka/update": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "summary": "入荷数を更新する",
                "parameters": [
                    {
                        "description": "jisekisu　input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.NyukaInputJANNKAJISURequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/api/v1/vender": {
            "get": {
                "description": "Get vender details by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "vender_no",
                        "name": "vender_no",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Vender"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.LoginResponse": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/response.LoginByIDResponse"
                },
                "token": {
                    "type": "string"
                }
            }
        },
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
        "models.Vender": {
            "type": "object",
            "properties": {
                "sire_name": {
                    "type": "string"
                },
                "sire_no": {
                    "type": "string"
                }
            }
        },
        "request.NyukaInputJANNKAJISURequest": {
            "type": "object",
            "properties": {
                "center_no": {
                    "type": "string"
                },
                "jiseki_su": {
                    "type": "integer"
                },
                "order_no": {
                    "type": "string"
                },
                "product_cd": {
                    "type": "string"
                },
                "up_date": {
                    "type": "string"
                }
            }
        },
        "request.NyukaInputJANRequest": {
            "type": "object",
            "properties": {
                "center_no": {
                    "type": "string"
                },
                "order_no": {
                    "type": "string"
                },
                "product_cd": {
                    "type": "string"
                }
            }
        },
        "request.NyukaInputRequest": {
            "type": "object",
            "properties": {
                "center_no": {
                    "type": "string",
                    "example": "0093"
                },
                "nyuka_begin": {
                    "type": "string",
                    "example": "20220101"
                },
                "nyuka_end": {
                    "type": "string",
                    "example": "20240101"
                },
                "sire_no": {
                    "type": "string",
                    "example": "4466"
                }
            }
        },
        "response.LoginByIDResponse": {
            "type": "object",
            "properties": {
                "employee_code": {
                    "type": "string"
                },
                "employee_name": {
                    "type": "string"
                }
            }
        },
        "response.NyukaInputJANResponse": {
            "type": "object",
            "properties": {
                "eda_no": {
                    "type": "string"
                },
                "jiseki_su": {
                    "type": "integer"
                },
                "order_no": {
                    "type": "string"
                },
                "order_su": {
                    "type": "integer"
                },
                "product_cd": {
                    "type": "string"
                },
                "product_name": {
                    "type": "string"
                },
                "sire_name": {
                    "type": "string"
                },
                "sire_no": {
                    "type": "string"
                },
                "yotei_su": {
                    "type": "integer"
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
