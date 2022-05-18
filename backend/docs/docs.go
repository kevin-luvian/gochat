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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/login/google": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Create google oauth redirect",
                "parameters": [
                    {
                        "description": "login google request",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginGoogleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginGoogleRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.ValidationError"
                        }
                    }
                }
            }
        },
        "/api/auth/signup": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "Create new user account",
                "parameters": [
                    {
                        "description": "create new user",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.SignupReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.TokenRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.ValidationError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ErrResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/temp": {
            "get": {
                "description": "check for server connection",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Temporary Handler",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/user/temp": {
            "get": {
                "description": "test REST on User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Temporary Handler",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "app.ErrResponse": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "an error has occured"
                }
            }
        },
        "app.VErr": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string",
                    "example": "input_field"
                },
                "message": {
                    "type": "string",
                    "example": "input_field must have a value!"
                }
            }
        },
        "app.ValidationError": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app.VErr"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "invalid request parameters"
                }
            }
        },
        "auth.LoginGoogleReq": {
            "type": "object",
            "properties": {
                "redirect_url": {
                    "type": "string",
                    "example": "http://localhost:8000/auth/google"
                }
            }
        },
        "auth.LoginGoogleRes": {
            "type": "object",
            "properties": {
                "oauth_url": {
                    "type": "string",
                    "example": "https://accounts.google.com/o/oauth2/auth?..."
                },
                "state": {
                    "type": "string",
                    "example": "GoogleAuthCredential_12345"
                }
            }
        },
        "auth.SignupReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "abc@def.gh"
                },
                "password": {
                    "type": "string",
                    "example": "pass1234"
                },
                "username": {
                    "type": "string",
                    "example": "rick"
                }
            }
        },
        "auth.TokenRes": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:7000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "GoChat API documentation",
	Description:      "This is the core server for GoChat to manage accounts and contacts.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
