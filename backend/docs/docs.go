// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/students": {
            "get": {
                "description": "Get all students from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Gets all students",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Student"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Edits a student from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Updates a student",
                "parameters": [
                    {
                        "description": "Student data",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.StudentEdit"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Student"
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a student to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Create a new student",
                "parameters": [
                    {
                        "description": "Student data",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.StudentCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Student"
                        }
                    }
                }
            }
        },
        "/students/{id}": {
            "delete": {
                "description": "Deletes a student from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Deletes a student",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.Student"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Student": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "grade": {
                    "description": "The grade assigned to this student",
                    "type": "integer"
                },
                "id": {
                    "description": "The ID of the Student",
                    "type": "integer"
                },
                "name": {
                    "description": "The name of the Student",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.StudentCreate": {
            "type": "object",
            "properties": {
                "grade": {
                    "description": "The Grade of the Student",
                    "type": "integer"
                },
                "name": {
                    "description": "The ID of the Student",
                    "type": "string"
                }
            }
        },
        "models.StudentEdit": {
            "type": "object",
            "properties": {
                "grade": {
                    "description": "The Grade of the Student",
                    "type": "integer"
                },
                "id": {
                    "description": "The Id of the student",
                    "type": "integer"
                },
                "name": {
                    "description": "The Name of the student",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Student API",
	Description:      "A simple CRUD API using Gin and GORM.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
