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
        "/receipts/process": {
            "post": {
                "description": "Process a receipt calculating generated points",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Receipt"
                ],
                "summary": "Process a receipt",
                "parameters": [
                    {
                        "description": "Receipt to be process",
                        "name": "receipt_to_process",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/receipts.Receipt"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Receipt created data",
                        "schema": {
                            "$ref": "#/definitions/receipts.ProcessOutput"
                        }
                    },
                    "500": {
                        "description": "Something unidentified has occurred"
                    }
                }
            }
        },
        "/receipts/{id}/points": {
            "get": {
                "description": "Returns a receipt registered points by receipt id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Receipt"
                ],
                "summary": "Returns a receipt registered points by receipt id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "receipt UUID to get",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Receipt points data",
                        "schema": {
                            "$ref": "#/definitions/receipts.FindPointsByReceiptIdOutput"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apperrors.MessagedError"
                        }
                    },
                    "500": {
                        "description": "Something unidentified has occurred"
                    }
                }
            }
        }
    },
    "definitions": {
        "apperrors.MessagedError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "receipts.FindPointsByReceiptIdOutput": {
            "type": "object",
            "properties": {
                "points": {
                    "type": "integer"
                }
            }
        },
        "receipts.ProcessOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "receipts.Receipt": {
            "type": "object",
            "required": [
                "purchaseDate",
                "purchaseTime",
                "retailer",
                "total"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/receipts.ReceiptItem"
                    }
                },
                "purchaseDate": {
                    "type": "string"
                },
                "purchaseTime": {
                    "type": "string"
                },
                "retailer": {
                    "type": "string"
                },
                "total": {
                    "type": "string",
                    "example": "0"
                }
            }
        },
        "receipts.ReceiptItem": {
            "type": "object",
            "required": [
                "price",
                "shortDescription"
            ],
            "properties": {
                "price": {
                    "type": "string",
                    "example": "0"
                },
                "shortDescription": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "fetch-receipt-processor-challenge",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
