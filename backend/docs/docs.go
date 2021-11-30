// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        }
    },
    "definitions": {
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
                },
                "value": {
                    "type": "string"
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
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0.0",
	Host:        "localhost:8000",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "GoChat API documentation",
	Description: "This is the core server for GoChat to manage accounts and contacts.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
