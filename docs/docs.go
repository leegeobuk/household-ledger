// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/household-ledger/ledger": {
            "post": {
                "description": "Add ledger",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basic"
                ],
                "summary": "Ledger",
                "parameters": [
                    {
                        "description": "Add ledger request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/resource.ReqAddLedger"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/resource.ResAddLedger"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/resource.ResErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/resource.ResErr"
                        }
                    }
                }
            }
        },
        "/api/household-ledger/ping": {
            "get": {
                "description": "Ping",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basic"
                ],
                "summary": "Ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/resource.ResErr"
                        }
                    }
                }
            }
        },
        "/api/household-ledger/time": {
            "get": {
                "description": "Time",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basic"
                ],
                "summary": "Time",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resource.ResGetTime"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/resource.ResErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "resource.ReqAddLedger": {
            "type": "object",
            "required": [
                "date",
                "description",
                "income",
                "user_id"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "income": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "resource.ResAddLedger": {
            "type": "object",
            "properties": {
                "ledger_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "resource.ResErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "resource.ResGetTime": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
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
