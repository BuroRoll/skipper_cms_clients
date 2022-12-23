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
        "/clients/": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Получение списка менторов и менти",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Communication": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "messenger": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Messenger"
                    }
                },
                "parentId": {
                    "description": "gorm.Model",
                    "type": "integer"
                }
            }
        },
        "models.Education": {
            "type": "object",
            "properties": {
                "degree": {
                    "type": "string"
                },
                "endYear": {
                    "type": "integer"
                },
                "institution": {
                    "type": "string"
                },
                "parentId": {
                    "description": "gorm.Model",
                    "type": "integer"
                },
                "startYear": {
                    "type": "integer"
                }
            }
        },
        "models.Messenger": {
            "type": "object",
            "properties": {
                "communications": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Communication"
                    }
                },
                "name": {
                    "description": "gorm.Model",
                    "type": "string"
                }
            }
        },
        "models.OtherInformation": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "parentId": {
                    "description": "gorm.Model",
                    "type": "integer"
                }
            }
        },
        "models.Report": {
            "type": "object",
            "properties": {
                "fromUserId": {
                    "description": "gorm.Model",
                    "type": "integer"
                },
                "reportText": {
                    "type": "string"
                },
                "reportTheme": {
                    "type": "string"
                },
                "toUserId": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "averageClassPrice": {
                    "type": "integer"
                },
                "communications": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Communication"
                    }
                },
                "dateOfBirthday": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "educations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Education"
                    }
                },
                "email": {
                    "description": "gorm.Model\nBase information",
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "isMentor": {
                    "type": "boolean"
                },
                "isVerifyEmail": {
                    "type": "boolean"
                },
                "isVerifyPhone": {
                    "type": "boolean"
                },
                "otherInformation": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OtherInformation"
                    }
                },
                "patronymic": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "profilePicture": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "reports": {
                    "description": "Жалобы на пользователя",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Report"
                    }
                },
                "secondName": {
                    "type": "string"
                },
                "specialization": {
                    "description": "Mentor information",
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "userReports": {
                    "description": "Жалобы пользователя",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Report"
                    }
                },
                "workExperiences": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.WorkExperience"
                    }
                }
            }
        },
        "models.WorkExperience": {
            "type": "object",
            "properties": {
                "endYear": {
                    "type": "string"
                },
                "organization": {
                    "type": "string"
                },
                "parentId": {
                    "description": "gorm.Model",
                    "type": "integer"
                },
                "startYear": {
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
	Title:            "Auth service",
	Description:      "Auth methods for skipper cms",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
