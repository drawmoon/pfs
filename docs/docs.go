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
        "/api/fs/directories/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "directories"
                ],
                "summary": "获取目录详细信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "目录id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/directory.GetInfoCaseRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    }
                }
            }
        },
        "/api/fs/directories/{id}/directories": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "directories"
                ],
                "summary": "获取指定目录下一层级的目录列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "目录id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/directory.GetDirectoriesCaseRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    }
                }
            }
        },
        "/api/fs/directories/{id}/files": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "directories"
                ],
                "summary": "获取指定目录下一层级的文件列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "目录id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/directory.GetFilesCaseRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    }
                }
            }
        },
        "/api/fs/files/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "files"
                ],
                "summary": "获取文件详细信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文件id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/file.GetInfoCaseRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProblemDetails"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "directory.GetDirectoriesCaseRes": {
            "type": "object",
            "properties": {
                "hasNextPages": {
                    "type": "boolean",
                    "format": "bool",
                    "example": true
                },
                "hasPrevPages": {
                    "type": "boolean",
                    "format": "bool",
                    "example": true
                },
                "items": {
                    "type": "array",
                    "items": {
                        "format": "[]GetDirectoryCaseRes",
                        "$ref": "#/definitions/directory.GetDirectoryCaseRes"
                    }
                },
                "pageIndex": {
                    "type": "integer",
                    "format": "int",
                    "example": 1
                },
                "pageSize": {
                    "type": "integer",
                    "format": "int",
                    "example": 10
                },
                "totalCount": {
                    "type": "integer",
                    "format": "int",
                    "example": 50
                },
                "totalPages": {
                    "type": "integer",
                    "format": "int",
                    "example": 5
                }
            }
        },
        "directory.GetDirectoryCaseRes": {
            "type": "object"
        },
        "directory.GetFilesCaseRes": {
            "type": "object"
        },
        "directory.GetInfoCaseRes": {
            "type": "object"
        },
        "file.GetInfoCaseRes": {
            "type": "object"
        },
        "http.ProblemDetails": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string",
                    "format": "string",
                    "example": "a error detail"
                },
                "instance": {
                    "type": "string",
                    "format": "string",
                    "example": "/api/to/path"
                },
                "status": {
                    "type": "integer",
                    "format": "int",
                    "example": 400
                },
                "title": {
                    "type": "string",
                    "format": "string",
                    "example": "a error title"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Power File Storage API",
	Description:      "DDD-based file storage management project built with Gin.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
