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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/EstateType": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new EstateType",
                "tags": [
                    "EstateType"
                ],
                "summary": "Create a new EstateType",
                "parameters": [
                    {
                        "description": "The EstateType to create  ",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EstateType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EstateType"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/EstateType/": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update Assginmenttype",
                "tags": [
                    "EstateType"
                ],
                "summary": "update EstateType",
                "parameters": [
                    {
                        "description": "The EstateType to update  ",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EstateType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EstateType"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/Province": {
            "post": {
                "description": "Create a new Province",
                "tags": [
                    "Province"
                ],
                "summary": "Create a new Province",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Province"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/Province/": {
            "put": {
                "description": "update Assginmenttype",
                "tags": [
                    "Province"
                ],
                "summary": "update Province",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Province"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/Unit": {
            "post": {
                "description": "Create a new Unit",
                "tags": [
                    "Unit"
                ],
                "summary": "Create a new Unit",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Unit"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/Unit/": {
            "put": {
                "description": "update Assginmenttype",
                "tags": [
                    "Unit"
                ],
                "summary": "update Unit",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Unit"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/assignmenttype/": {
            "put": {
                "description": "update Assginmenttype",
                "tags": [
                    "AssignmentType"
                ],
                "summary": "update Assginmenttype",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AssignmentType"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Assginmenttype",
                "tags": [
                    "AssignmentType"
                ],
                "summary": "Create a new Assginmenttype",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AssignmentType"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/assignmenttype/id": {
            "get": {
                "description": "Get AssignmentType by id",
                "tags": [
                    "AssignmentType"
                ],
                "summary": "Get AssignmentType by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AssignmentType"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete AssignmentType by id",
                "tags": [
                    "AssignmentType"
                ],
                "summary": "Delete AssignmentType by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AssignmentType"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/assignmenttypes": {
            "get": {
                "description": "Get All AssignmentType",
                "tags": [
                    "AssignmentType"
                ],
                "summary": "Get All AssignmentType",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.AssignmentType"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/estatetype/id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get EstateType by id",
                "tags": [
                    "EstateType"
                ],
                "summary": "Get EstateType by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EstateType"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete EstateType by id",
                "tags": [
                    "EstateType"
                ],
                "summary": "Delete EstateType by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EstateType"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/estatetypes": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get All EstateType",
                "tags": [
                    "EstateType"
                ],
                "summary": "Get All EstateType",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.EstateType"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/form/": {
            "put": {
                "description": "update Form",
                "tags": [
                    "Froms"
                ],
                "summary": "update Form",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Form"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Form",
                "tags": [
                    "Froms"
                ],
                "summary": "Delete Form",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Form"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/form/id": {
            "get": {
                "description": "Get Form",
                "tags": [
                    "Froms"
                ],
                "summary": "Get Form",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Form"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/forms/": {
            "get": {
                "description": "Get Forms",
                "tags": [
                    "Froms"
                ],
                "summary": "Get Forms",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Form"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/province/city/id": {
            "post": {
                "description": "Add City to Province",
                "tags": [
                    "Province"
                ],
                "summary": "Add City to Province",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Province ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.City"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete City in Province",
                "tags": [
                    "Province"
                ],
                "summary": "Delete City in Province",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Province ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.City"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/province/id": {
            "get": {
                "description": "Get Province by id",
                "tags": [
                    "Province"
                ],
                "summary": "Get Province by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Province"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Province by id",
                "tags": [
                    "Province"
                ],
                "summary": "Delete Province by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Province"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/provinces": {
            "get": {
                "description": "Get All Province",
                "tags": [
                    "Province"
                ],
                "summary": "Get All Province",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Province"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "description": "Signin",
                "tags": [
                    "User"
                ],
                "summary": "Signin",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/unit/id": {
            "get": {
                "description": "Get Unit by id",
                "tags": [
                    "Unit"
                ],
                "summary": "Get Unit by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Unit"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Unit by id",
                "tags": [
                    "Unit"
                ],
                "summary": "Delete Unit by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Unit"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/units": {
            "get": {
                "description": "Get All Unit",
                "tags": [
                    "Unit"
                ],
                "summary": "Get All Unit",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Unit"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/id": {
            "get": {
                "description": "get user by id",
                "tags": [
                    "User"
                ],
                "summary": "Get user by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "get all users",
                "tags": [
                    "User"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AssignmentType": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "required: false",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.City": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.EstateType": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "required: false",
                    "type": "string"
                },
                "name": {
                    "description": "The name for a EstateType\nexample: خرید\nrequired: true",
                    "type": "string"
                }
            }
        },
        "models.Field": {
            "type": "object",
            "properties": {
                "fieldtype": {
                    "type": "string"
                },
                "fileds": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Field"
                    }
                },
                "id": {
                    "type": "string"
                },
                "max": {
                    "type": "number"
                },
                "min": {
                    "type": "number"
                },
                "option": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "optional": {
                    "type": "boolean"
                },
                "title": {
                    "description": "Name      string        ` + "`" + `json:\"name\" bson:\"name\"` + "`" + `",
                    "type": "string"
                },
                "value": {
                    "$ref": "#/definitions/models.Valuetype"
                }
            }
        },
        "models.Form": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "sections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Section"
                    }
                },
                "title": {
                    "description": "Name string        ` + "`" + `json:\"name\" bson:\"name\"` + "`" + `\nSections []map[string]interface{} ` + "`" + `json:\"sections\" bson:\"sections\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "models.Province": {
            "type": "object",
            "properties": {
                "cites": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.City"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Section": {
            "type": "object",
            "properties": {
                "fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Field"
                    }
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "description": "Name   string        ` + "`" + `json:\"name\" bson:\"name\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "models.Unit": {
            "type": "object",
            "properties": {
                "- bson:": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "description": "owner 0 Admin 1 User 2",
                    "type": "integer"
                }
            }
        },
        "models.Valuetype": {
            "type": "object"
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
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
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "User API documentation",
	Description: "",
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
	swag.Register(swag.Name, &s{})
}
