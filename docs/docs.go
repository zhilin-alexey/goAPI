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
        "/info": {
            "get": {
                "description": "get people by passport details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Show a people",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Passport serie",
                        "name": "passportSerie",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Passport number",
                        "name": "passportNumber",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goAPI.Person"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/people": {
            "get": {
                "description": "get people by multiple filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Show a multiple full people",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Surname",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Patronymic",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Address",
                        "name": "address",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Passport serie",
                        "name": "passportSerie",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Passport number",
                        "name": "passportNumber",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/goAPI.Person"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "create people by passport number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Create people",
                "parameters": [
                    {
                        "type": "string",
                        "name": "passportNumber",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goAPI.Person"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "delete people by multiple filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Delete people",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Surname",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Patronymic",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Address",
                        "name": "address",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Passport serie",
                        "name": "passportSerie",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Passport number",
                        "name": "passportNumber",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/people/{id}": {
            "patch": {
                "description": "edit people by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Edit people",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "People data to edit",
                        "name": "people",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.EditParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goAPI.Person"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/people/{id}/task/end": {
            "post": {
                "description": "end task by people id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "End task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "People id",
                        "name": "people",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goAPI.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/people/{id}/task/start": {
            "post": {
                "description": "start task by people id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Start task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "People id",
                        "name": "people",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goAPI.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/people/{id}/tasks": {
            "get": {
                "description": "get tasks by people id and period of time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get people tasks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "People id",
                        "name": "people",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Period start",
                        "name": "periodStart",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Period end",
                        "name": "periodEnd",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/goAPI.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.EditParams": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "minLength": 10,
                    "example": "3-й Автозаводский проезд, вл13, Москва, 115280"
                },
                "name": {
                    "type": "string",
                    "minLength": 3,
                    "example": "Иван"
                },
                "passportNumber": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6,
                    "example": "123456"
                },
                "passportSerie": {
                    "type": "string",
                    "maxLength": 4,
                    "minLength": 4,
                    "example": "1234"
                },
                "patronymic": {
                    "type": "string",
                    "minLength": 3,
                    "example": "Иванович"
                },
                "surname": {
                    "type": "string",
                    "minLength": 3,
                    "example": "Иванов"
                }
            }
        },
        "goAPI.Person": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "3-й Автозаводский проезд, вл13, Москва, 115280"
                },
                "id": {
                    "type": "string",
                    "format": "uuid",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "name": {
                    "type": "string",
                    "example": "Иван"
                },
                "passportNumber": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6,
                    "example": "123456"
                },
                "passportSerie": {
                    "type": "string",
                    "maxLength": 4,
                    "minLength": 4,
                    "example": "1234"
                },
                "patronymic": {
                    "type": "string",
                    "example": "Иванович"
                },
                "surname": {
                    "type": "string",
                    "example": "Иванов"
                }
            }
        },
        "goAPI.Task": {
            "type": "object",
            "properties": {
                "endTime": {
                    "type": "string",
                    "format": "dateTime",
                    "example": "2022-01-01T00:00:00Z"
                },
                "id": {
                    "type": "string",
                    "format": "uuid",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "name": {
                    "type": "string",
                    "example": "Помыть посуду"
                },
                "peopleId": {
                    "type": "string"
                },
                "startTime": {
                    "type": "string",
                    "format": "dateTime",
                    "example": "2022-01-01T00:00:00Z"
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
	Schemes:          []string{"http"},
	Title:            "goAPI",
	Description:      "go API with task tracking",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
