// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/auth/refresh": {
            "patch": {
                "description": "create user email and phone must be unique",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "updateAccessToken",
                "operationId": "update-access-token",
                "parameters": [
                    {
                        "description": "refresh_token",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateTokenInputDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "should get access token",
                        "schema": {
                            "$ref": "#/definitions/dto.JWT"
                        }
                    },
                    "400": {
                        "description": "invalid request body or error request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "invalid input parameter",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "should receive 3 tokens: refresh_token, access_token, rtc_token. Access to private methods is done using access_token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "signIn",
                "operationId": "sign-in",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.SignIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "should get 3 tokens",
                        "schema": {
                            "$ref": "#/definitions/dto.SignInOutputDto"
                        }
                    },
                    "400": {
                        "description": "invalid request body or error request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "invalid input parameter",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "create user email and phone must be unique",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "signUp",
                "operationId": "sign-up",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignUpByEmailInputDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "should get uuid",
                        "schema": {
                            "$ref": "#/definitions/dto.SignUpOutputDto"
                        }
                    },
                    "400": {
                        "description": "invalid request body or error request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "invalid input parameter",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/log/push-text-log": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "push-text-log",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "log"
                ],
                "summary": "pushTextLog",
                "operationId": "push-text-log",
                "parameters": [
                    {
                        "description": "log text and uuid rooms ",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.LogTextInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "just status code",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid request body or error request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "unauth",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "invalid input parameter",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get me user information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Me",
                "operationId": "me",
                "responses": {
                    "200": {
                        "description": "should get user",
                        "schema": {
                            "$ref": "#/definitions/dto.MeOutputDto"
                        }
                    },
                    "401": {
                        "description": "unauth",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/op/update-subscribes": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "when a user loses connection with centrifugo, subscriptions to all channels are automatically lost, in order to restore the subscription when the token expires or disconnects, this route is used",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "operating"
                ],
                "summary": "updateSubscribes",
                "operationId": "update-subscribes",
                "responses": {
                    "200": {
                        "description": "just status code",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "unauth",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/room/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "show rooms",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "summary": "showRooms",
                "operationId": "show-rooms",
                "responses": {
                    "200": {
                        "description": "collection of rooms",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid request body or error request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "unauth",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "invalid input parameter",
                        "schema": {
                            "type": "string"
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
                "description": "create room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "summary": "createRoom",
                "operationId": "create-room",
                "parameters": [
                    {
                        "description": "room_name",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.CreateRoom"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "only status code",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid request body or error request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "unauth",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "invalid input parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/room/invite": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "invite user to room. User can only join by invitation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "summary": "inviteRoom",
                "operationId": "invite-room",
                "parameters": [
                    {
                        "description": "user_id and room_id",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.InviteRoom"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "only status code",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid request body or error request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "unauth",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "invalid input parameter",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/room/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "show rooms",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "summary": "showRoomById",
                "operationId": "show-room-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "room uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "collection of rooms",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid request body or error request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "unauth",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "invalid input parameter",
                        "schema": {
                            "type": "string"
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
                "description": "join room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "summary": "joinRoom",
                "operationId": "join-room",
                "parameters": [
                    {
                        "type": "string",
                        "description": "room uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "collection of rooms",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid request body or error request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "unauth",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "invalid input parameter",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.SignIn": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.JWT": {
            "type": "object",
            "properties": {
                "expired_at": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.MeOutputDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fio": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.SignInOutputDto": {
            "type": "object",
            "properties": {
                "access_token": {
                    "$ref": "#/definitions/dto.JWT"
                },
                "refresh_token": {
                    "$ref": "#/definitions/dto.JWT"
                },
                "rtc_host": {
                    "type": "string"
                },
                "rtc_token": {
                    "$ref": "#/definitions/dto.JWT"
                }
            }
        },
        "dto.SignUpByEmailInputDto": {
            "type": "object",
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fio": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.SignUpOutputDto": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateTokenInputDto": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "input.CreateRoom": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "input.InviteRoom": {
            "type": "object",
            "properties": {
                "room_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "input.LogTextInput": {
            "type": "object",
            "properties": {
                "room_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "text": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
