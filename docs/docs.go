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
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "942801422@qq.com"
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
        "/article/info/{id}": {
            "get": {
                "description": "查询单个文章接口",
                "tags": [
                    "文章接口"
                ],
                "summary": "查询单个文章",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "文章编号",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询文章成功",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseUser"
                        }
                    },
                    "400": {
                        "description": "查询文章失败",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "用户登录，需要账号、密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户接口"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户登录",
                        "name": "userinfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Success"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Token"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/add": {
            "post": {
                "description": "新增用户，需要账号、密码、权限码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户接口"
                ],
                "summary": "新增用户",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "userinfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Success"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.UserInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "通过用户ID删除用户（管理员有此权限）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户接口"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":200, \"msg\":\"删除用户成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Error": {
            "type": "object",
            "properties": {
                "msg": {
                    "description": "返回信息",
                    "type": "string"
                },
                "status": {
                    "description": "返回错误码",
                    "type": "integer"
                }
            }
        },
        "model.Success": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "返回数据"
                },
                "msg": {
                    "description": "返回信息",
                    "type": "string"
                },
                "status": {
                    "description": "返回成功码",
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "model.Token": {
            "type": "object",
            "properties": {
                "token": {
                    "description": "token",
                    "type": "string"
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "role": {
                    "description": "权限码1：管理员；2：普通用户",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "model.UserLogin": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "v1.ResponseError": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "Error"
                },
                "status": {
                    "type": "integer",
                    "example": 500
                }
            }
        },
        "v1.ResponseUser": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "OK"
                },
                "status": {
                    "type": "integer",
                    "example": 200
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
	Version:     "1.0",
	Host:        "localhost:2333",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Gin Blog API v1",
	Description: "Gin博客接口文档",
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