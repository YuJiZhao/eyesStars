// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "耶瞳",
            "url": "http://space.eyesspace.top",
            "email": "eyesyeager@gmail.com"
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
        "/context/contextAdd": {
            "post": {
                "security": [
                    {
                        "sToken,lToken": []
                    }
                ],
                "description": "权限：Admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "context"
                ],
                "summary": "添加配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短token",
                        "name": "sToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "长token",
                        "name": "lToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "请求参数",
                        "name": "receiver",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/receiver.ContextAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/context/contextUpdate": {
            "put": {
                "security": [
                    {
                        "sToken,lToken": []
                    }
                ],
                "description": "权限：Admin；若表中name不存在则添加配置，已存在则修改配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "context"
                ],
                "summary": "修改配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短token",
                        "name": "sToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "长token",
                        "name": "lToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "请求参数",
                        "name": "receiver",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/receiver.ContextUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/context/initSite": {
            "get": {
                "security": [
                    {
                        "sToken,lToken": []
                    }
                ],
                "description": "权限：Visitor；该接口处理方式要求前端必须是服务端渲染，因为接口会明文返回密钥等信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "context"
                ],
                "summary": "网站信息初始化",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短token",
                        "name": "sToken",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "长token",
                        "name": "lToken",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/star/starAdd": {
            "post": {
                "security": [
                    {
                        "sToken,lToken": []
                    }
                ],
                "description": "权限：User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "star"
                ],
                "summary": "添加星星",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短token",
                        "name": "sToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "长token",
                        "name": "lToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "请求参数",
                        "name": "receiver",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/receiver.StarAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/star/starGetById/{id}": {
            "get": {
                "security": [
                    {
                        "sToken,lToken": []
                    }
                ],
                "description": "权限：Visitor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "star"
                ],
                "summary": "通过id获取星星",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短token",
                        "name": "sToken",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "长token",
                        "name": "lToken",
                        "in": "header"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "寄语id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/star/starsGet": {
            "get": {
                "security": [
                    {
                        "sToken,lToken": []
                    }
                ],
                "description": "权限：Visitor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "star"
                ],
                "summary": "批量查询星星",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短token",
                        "name": "sToken",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "长token",
                        "name": "lToken",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "已读id密串",
                        "name": "ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/track/login": {
            "get": {
                "security": [
                    {
                        "sToken,lToken": []
                    }
                ],
                "description": "权限：Visitor，无响应内容",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "track"
                ],
                "summary": "登录埋点",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短token",
                        "name": "sToken",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "长token",
                        "name": "lToken",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "埋点数据加密包",
                        "name": "package",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/track/visit": {
            "get": {
                "security": [
                    {
                        "sToken,lToken": []
                    }
                ],
                "description": "权限：Visitor，无响应内容",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "track"
                ],
                "summary": "进站埋点",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短token",
                        "name": "sToken",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "长token",
                        "name": "lToken",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "埋点数据加密包",
                        "name": "package",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/userThrift/userInfoGet": {
            "get": {
                "security": [
                    {
                        "sToken,lToken": []
                    }
                ],
                "description": "权限：User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userThrift"
                ],
                "summary": "获取用户基本信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短token",
                        "name": "sToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "长token",
                        "name": "lToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/userThrift/userInfoUpdate": {
            "put": {
                "security": [
                    {
                        "sToken,lToken": []
                    }
                ],
                "description": "权限：User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userThrift"
                ],
                "summary": "修改用户基本信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短token",
                        "name": "sToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "长token",
                        "name": "lToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "请求参数",
                        "name": "receiver",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/receiver.UserInfoUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/userThrift/userUpdateAvatar": {
            "post": {
                "security": [
                    {
                        "sToken,lToken": []
                    }
                ],
                "description": "权限：User",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userThrift"
                ],
                "summary": "修改用户头像",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短token",
                        "name": "sToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "长token",
                        "name": "lToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "头像文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "receiver.ContextAdd": {
            "type": "object",
            "required": [
                "content",
                "name"
            ],
            "properties": {
                "content": {
                    "description": "配置内容",
                    "type": "string"
                },
                "name": {
                    "description": "配置名",
                    "type": "string"
                },
                "remarks": {
                    "description": "配置说明",
                    "type": "string"
                }
            }
        },
        "receiver.ContextUpdate": {
            "type": "object",
            "required": [
                "content",
                "name"
            ],
            "properties": {
                "content": {
                    "description": "配置内容",
                    "type": "string"
                },
                "name": {
                    "description": "配置名",
                    "type": "string"
                }
            }
        },
        "receiver.StarAdd": {
            "type": "object",
            "required": [
                "content"
            ],
            "properties": {
                "content": {
                    "description": "\"提交内容\"",
                    "type": "string",
                    "maxLength": 300
                },
                "emailNeed": {
                    "description": "\"是否需要寄送服务\"",
                    "type": "boolean",
                    "default": false
                },
                "name": {
                    "description": "\"署名\"",
                    "type": "string",
                    "maxLength": 10
                }
            }
        },
        "receiver.UserInfoUpdate": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "description": "\"昵称\"",
                    "type": "string",
                    "maxLength": 10
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "自定义错误码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据"
                },
                "msg": {
                    "description": "信息",
                    "type": "string"
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
	Title:            "耶瞳星空",
	Description:      "耶瞳星空，来自星星的寄语",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
