package dosc

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
        "/memberships/{id}": {
            "get": {
                "description": "멤버십 정보를 조회합니다. (상세 설명)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Memberships"
                ],
                "summary": "멤버십 정보 단건 조회",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Membership uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/memberships.GetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/memberships.Fail400GetResponse"
                        }
                    }
                }
            }
        },
        "/v1/memberships": {
            "post": {
                "description": "멤버십을 생성합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Memberships"
                ],
                "summary": "멤버십 생성",
                "parameters": [
                    {
                        "description": "user_name:사용자의 이름, membership_type:naver,toss,pacyco 중 하나",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/memberships.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/memberships.CreateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "memberships.CreateRequest": {
            "type": "object",
            "properties": {
                "membership_type": {
                    "type": "string",
                    "example": "toss"
                },
                "user_name": {
                    "type": "string",
                    "example": "andy"
                }
            }
        },
        "memberships.CreateResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "354660dc-f798-11ec-b939-0242ac120002"
                },
                "membership_type": {
                    "type": "string",
                    "example": "toss"
                },
                "user_name": {
                    "type": "string",
                    "example": "andy"
                }
            }
        },
        "memberships.Fail400GetResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Bad Request"
                }
            }
        },
        "memberships.GetResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "354660dc-f798-11ec-b939-0242ac120002"
                },
                "membership_type": {
                    "type": "string",
                    "example": "toss"
                },
                "user_name": {
                    "type": "string",
                    "example": "andy"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Membership API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
