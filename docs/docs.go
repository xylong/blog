// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/categories": {
            "get": {
                "description": "分类列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "分类列表",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/dto.TagListOutput"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建分类",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "创建分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":\"\",\"msg\":\"ok\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改分类",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "修改分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":\"\",\"msg\":\"ok\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/categories/{id}": {
            "delete": {
                "description": "删除分类",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "删除分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "分类ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":\"\",\"msg\":\"ok\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "用户登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户接口"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":\"token\",\"msg\":\"ok\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/register": {
            "post": {
                "description": "用户注册",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户接口"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":null,\"msg\":\"ok\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/tags": {
            "get": {
                "description": "标签列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "标签列表",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/dto.TagListOutput"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改标签",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "修改标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TagUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":\"\",\"msg\":\"ok\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建标签",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "创建标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TagInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":\"\",\"msg\":\"ok\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/{id}": {
            "delete": {
                "description": "删除标签",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "删除标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "标签ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":\"\",\"msg\":\"ok\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取个人信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户接口"
                ],
                "summary": "个人信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/dto.Profile"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "ping",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"pong\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CategoryInput": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "编程"
                }
            }
        },
        "dto.CategoryUpdate": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "description": "id",
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "编程"
                }
            }
        },
        "dto.LoginInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "abc123@qq.com"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "dto.Profile": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "name": {
                    "description": "昵称",
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterInput": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "abc123@qq.com"
                },
                "name": {
                    "description": "昵称",
                    "type": "string",
                    "example": "小明"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string",
                    "example": "19999999999"
                }
            }
        },
        "dto.TagInput": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "description": "标签名",
                    "type": "string",
                    "example": "mysql"
                }
            }
        },
        "dto.TagListOutput": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.TagOutput"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "dto.TagOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "name": {
                    "description": "标签名",
                    "type": "string"
                }
            }
        },
        "dto.TagUpdate": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "description": "id",
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "description": "标签名",
                    "type": "string",
                    "example": "mysql"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
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
