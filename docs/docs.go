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
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/timetable/groups/": {
            "get": {
                "description": "Get all groups list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "timetable"
                ],
                "summary": "Get groups list",
                "operationId": "get-groups",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.ResponseBase"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/timetable.GroupResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/timetable/sessions/{group_name}": {
            "get": {
                "description": "Get session timetable by group name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "timetable"
                ],
                "summary": "Get session timetable",
                "operationId": "get-group-session",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group Name",
                        "name": "group_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.ResponseBase"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/timetable.SessionResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/timetable/timetables/{group_name}": {
            "get": {
                "description": "Get group timetable by group name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "timetable"
                ],
                "summary": "Get group timetable",
                "operationId": "get-group-timetable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group Name",
                        "name": "group_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.ResponseBase"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/timetable.TimetableResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ResponseBase": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "result": {
                    "type": "boolean"
                }
            }
        },
        "timetable.GroupResponse": {
            "type": "object",
            "properties": {
                "direction": {
                    "type": "string"
                },
                "faculty": {
                    "type": "string"
                },
                "group_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "timetable.SessionResponse": {
            "type": "object",
            "properties": {
                "addition": {
                    "type": "string"
                },
                "group_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "table": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/timetable.Subject"
                    }
                }
            }
        },
        "timetable.Subject": {
            "type": "object",
            "properties": {
                "location": {
                    "type": "string"
                },
                "subject_name": {
                    "type": "string"
                },
                "subject_type": {
                    "type": "string"
                },
                "teacher": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "timetable.TimetableResponse": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "string"
                },
                "group_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "table": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/timetable.Subject"
                    }
                },
                "week_num": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
