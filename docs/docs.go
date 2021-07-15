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
        "termsOfService": "https://github.com/yangjunyangjun/ledger",
        "contact": {
            "name": "API Support",
            "email": "zzzzzz1916@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ledger/v1/consume/add_consume": {
            "post": {
                "description": "新增消费记录",
                "tags": [
                    "消费管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/consume.AddConsumeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/consume/add_consume_type": {
            "post": {
                "description": "新增消费种类",
                "tags": [
                    "消费管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/consume.AddConsumeTypeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/consume/consume_list": {
            "get": {
                "description": "消费记录列表",
                "tags": [
                    "消费管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "end_time",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "start_time",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/consume/consume_type_list": {
            "get": {
                "description": "消费种类列表",
                "tags": [
                    "消费管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/consume/del_consume_type": {
            "delete": {
                "description": "删除种类列表",
                "tags": [
                    "消费管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/consume.DelConsumeTypeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/consume/update_consume": {
            "put": {
                "description": "更新消费记录",
                "tags": [
                    "消费管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/consume.UpdateConsumeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/consume/update_revenue": {
            "put": {
                "description": "更新收入记录",
                "tags": [
                    "收入管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/revenue.UpdateRevenueRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/file/download": {
            "get": {
                "description": "下载图片",
                "tags": [
                    "文件接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图片连接",
                        "name": "icon",
                        "in": "query",
                        "required": true
                    }
                ]
            }
        },
        "/ledger/v1/file/upload": {
            "post": {
                "description": "上传图片",
                "tags": [
                    "文件接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文件",
                        "name": "icon",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/revenue/add_revenue": {
            "post": {
                "description": "新增收入记录",
                "tags": [
                    "收入管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/revenue.AddRevenueRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/revenue/add_revenue_type": {
            "post": {
                "description": "新增收入种类",
                "tags": [
                    "收入管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/revenue.AddRevenueTypeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/revenue/del_revenue_type": {
            "delete": {
                "description": "删除收入种类",
                "tags": [
                    "收入管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/revenue.DelRevenueTypeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/revenue/revenue_list": {
            "get": {
                "description": "收入记录列表",
                "tags": [
                    "收入管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "end_time",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "start_time",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/revenue/revenue_type_list": {
            "get": {
                "description": "收入种类列表",
                "tags": [
                    "收入管理相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/user/captcha": {
            "get": {
                "description": "获取图片验证码",
                "tags": [
                    "用户相关接口"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/user/captcha_verify": {
            "post": {
                "description": "图片验证码校验",
                "tags": [
                    "用户相关接口"
                ],
                "parameters": [
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.Verify"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/user/login": {
            "post": {
                "description": "用户登录",
                "tags": [
                    "用户相关接口"
                ],
                "parameters": [
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/user/register": {
            "post": {
                "description": "用户注册",
                "tags": [
                    "用户相关接口"
                ],
                "parameters": [
                    {
                        "description": "JSON数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.Register"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/user/send_activate": {
            "get": {
                "description": "发送激活码",
                "tags": [
                    "用户相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        },
        "/ledger/v1/user/user_list": {
            "get": {
                "description": "用户列表",
                "tags": [
                    "用户相关接口"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "user_name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/webbase.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "consume.AddConsumeRequest": {
            "type": "object",
            "properties": {
                "money": {
                    "type": "number"
                },
                "remark": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "consume.AddConsumeTypeReq": {
            "type": "object",
            "properties": {
                "type_name": {
                    "type": "string"
                }
            }
        },
        "consume.DelConsumeTypeReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "consume.UpdateConsumeRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "money": {
                    "type": "number"
                },
                "remark": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "revenue.AddRevenueRequest": {
            "type": "object",
            "properties": {
                "money": {
                    "type": "number"
                },
                "remark": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "revenue.AddRevenueTypeReq": {
            "type": "object",
            "properties": {
                "type_name": {
                    "type": "string"
                }
            }
        },
        "revenue.DelRevenueTypeReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "revenue.UpdateRevenueRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "money": {
                    "type": "number"
                },
                "remark": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "user.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "user.Register": {
            "type": "object",
            "required": [
                "code",
                "password",
                "user_name"
            ],
            "properties": {
                "code": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "user.Verify": {
            "type": "object",
            "required": [
                "code",
                "id"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "webbase.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
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
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang ledger API",
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
