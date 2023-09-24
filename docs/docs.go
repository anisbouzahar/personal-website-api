// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Anis BOUZAHAR",
            "email": "hello@anis-bouzahar.dev"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/subscribe": {
            "post": {
                "description": "Post order writes a new subscriber to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriber"
                ],
                "summary": "Add a new subscriber to the database",
                "parameters": [
                    {
                        "description": "Add subscriber",
                        "name": "subscriber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Subscriber"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "application-specific error code",
                    "type": "integer",
                    "example": 404
                },
                "error": {
                    "description": "application-level error message, for debugging",
                    "type": "string",
                    "example": "The requested resource was not found on the server"
                },
                "status": {
                    "description": "user-level status message",
                    "type": "string",
                    "example": "Resource not found."
                }
            }
        },
        "Subscriber": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@example.com"
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
	Title:            "Anis's portfolio API",
	Description:      "This is the API for Anis's portfolio.\nWill gradually add endpoints based as requirements evolve",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
