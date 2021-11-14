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
        "/": {
            "get": {
                "description": "测试 Index 页",
                "tags": [
                    "测试"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"gcp\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/es/create/mytype": {
            "post": {
                "description": "创建es索引",
                "tags": [
                    "elasticsearch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "intvalue",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "{\"success\": false, \"message\": \"该ID已存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"success\": false, \"message\": \"创建错误500\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/es/get/mytype": {
            "post": {
                "description": "获取es索引",
                "tags": [
                    "elasticsearch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"该ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"success\": false, \"message\": \"错误500\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/es/get/paper/animer": {
            "post": {
                "description": "es获取Paper详细信息",
                "tags": [
                    "elasticsearch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"该PaperID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"success\": false, \"message\": \"错误500\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/es/get/paper/msg": {
            "post": {
                "description": "es获取Paper详细信息",
                "tags": [
                    "elasticsearch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"该PaperID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"success\": false, \"message\": \"错误500\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/es/update/mytype": {
            "post": {
                "description": "更新es索引",
                "tags": [
                    "elasticsearch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"更新成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"该ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"success\": false, \"message\": \"创建错误500\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/confirm": {
            "post": {
                "description": "验证邮箱",
                "tags": [
                    "用户管理"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "confirm_number",
                        "name": "confirm_number",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"用户验证邮箱成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "{\"success\": false, \"message\": \"用户已验证邮箱\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "402": {
                        "description": "{\"success\": false, \"message\": \"用户输入验证码错误}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"用户不存在}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "600": {
                        "description": "{\"success\": false, \"message\": \"用户待修改，传入false 更新验证码，否则为验证正确}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "查看用户个人信息",
                "tags": [
                    "用户管理"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查看用户信息成功\", \"data\": \"model.User的所有信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"用户ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "登录",
                "tags": [
                    "用户管理"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"登录成功\", \"detail\": user的信息}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "{\"success\": false, \"message\": \"没有该用户\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "402": {
                        "description": "{\"success\": false, \"message\": \"密码错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/modify": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "修改用户信息（支持修改用户名和密码）",
                "tags": [
                    "用户管理"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户个人信息",
                        "name": "user_info",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "原密码",
                        "name": "password_old",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "password_new",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"修改成功\", \"data\": \"model.User的所有信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"success\": false, \"message\": \"用户未登录\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "{\"success\": false, \"message\": \"原密码输入错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"用户ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"success\": false, \"message\": \"数据库操作时的其他错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "注册",
                "tags": [
                    "用户管理"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户邮箱",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": false, \"message\": \"用户已存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
	Schemes:     []string{"https"},
	Title:       "Slime Scholar Golang Backend",
	Description: "hzh company",
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
