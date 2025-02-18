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
        "/cancel-order/{email}/{id}": {
            "delete": {
                "description": "Cancel an order by order ID and email",
                "produces": [
                    "application/json"
                ],
                "summary": "Cancel an order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order Cancelled Successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-all-orders": {
            "get": {
                "description": "Retrieve all active orders",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.Order"
                                }
                            }
                        }
                    },
                    "404": {
                        "description": "No active orders found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-order": {
            "get": {
                "description": "Retrieve all orders for a given email",
                "produces": [
                    "application/json"
                ],
                "summary": "Get user orders",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "404": {
                        "description": "Order not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Check server availability",
                "produces": [
                    "text/plain"
                ],
                "summary": "Ping Server",
                "responses": {
                    "200": {
                        "description": "Hello From the Server!!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/place-order": {
            "post": {
                "description": "Create a new food order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Place an order",
                "parameters": [
                    {
                        "description": "Order Details",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order Details",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "Invalid Request Payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/update-address/{email}/{id}": {
            "put": {
                "description": "Update the delivery address for an order",
                "produces": [
                    "application/json"
                ],
                "summary": "Update address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "New Address",
                        "name": "new_address",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "unable to update new address",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Order": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "delivery_time": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8383",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "WeServeFood Delivery Order Management API",
	Description:      "API for managing food delivery orders",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
