// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Service storage brands and cosmetics.",
    "title": "Service-storage",
    "version": "0.0.1"
  },
  "basePath": "/api",
  "paths": {
    "/brand": {
      "get": {
        "description": "Get list brand",
        "operationId": "brandList",
        "responses": {
          "200": {
            "$ref": "#/responses/ListBrand"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      },
      "post": {
        "description": "Create new brand",
        "operationId": "brandPOST",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/BrandCreate"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/ListBrand"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      },
      "delete": {
        "description": "Delete brand",
        "operationId": "brandDELETE",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/NoContent"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/login": {
      "post": {
        "security": [],
        "description": "Auth method.",
        "operationId": "login",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/LoginInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/UserInfo"
            },
            "headers": {
              "AdminCookie": {
                "type": "string",
                "description": "Session token."
              }
            }
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/product": {
      "post": {
        "description": "Create new product",
        "operationId": "productPOST",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ProductCreate"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/OneProduct"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      },
      "delete": {
        "description": "Delete product",
        "operationId": "productDELETE",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/NoContent"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/product/{productID}": {
      "get": {
        "security": [],
        "operationId": "product",
        "parameters": [
          {
            "type": "integer",
            "description": "Track ID.",
            "name": "productID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/OneProduct"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/products": {
      "get": {
        "security": [],
        "operationId": "listProduct",
        "parameters": [
          {
            "type": "string",
            "name": "query",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "brand",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 25,
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 0,
            "name": "offset",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/ListProduct"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    }
  },
  "definitions": {
    "Brand": {
      "type": "object",
      "required": [
        "id",
        "name"
      ],
      "properties": {
        "id": {
          "$ref": "#/definitions/ID"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "BrandCreate": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string",
          "maxLength": 100,
          "minLength": 1
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ID": {
      "description": "Object ID.",
      "type": "integer"
    },
    "LoginInfo": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "$ref": "#/definitions/Password"
        },
        "username": {
          "$ref": "#/definitions/Username"
        }
      }
    },
    "Pagination": {
      "type": "object",
      "required": [
        "limit",
        "offset"
      ],
      "properties": {
        "limit": {
          "type": "integer",
          "maximum": 500,
          "minimum": 1
        },
        "offset": {
          "type": "integer"
        }
      }
    },
    "Password": {
      "description": "User password.",
      "type": "string",
      "format": "password",
      "maxLength": 100,
      "minLength": 1
    },
    "Product": {
      "type": "object",
      "required": [
        "id",
        "name",
        "description",
        "apply",
        "price",
        "brandName"
      ],
      "properties": {
        "apply": {
          "type": "string"
        },
        "brandName": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "$ref": "#/definitions/ID"
        },
        "name": {
          "type": "string",
          "maxLength": 100,
          "minLength": 1
        },
        "price": {
          "type": "integer"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "ProductCreate": {
      "type": "object",
      "required": [
        "name",
        "description",
        "apply",
        "price",
        "brandID"
      ],
      "properties": {
        "apply": {
          "type": "string"
        },
        "brandID": {
          "type": "integer"
        },
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string",
          "maxLength": 100,
          "minLength": 1
        },
        "price": {
          "type": "integer"
        }
      }
    },
    "UserInfo": {
      "type": "object",
      "required": [
        "username",
        "token"
      ],
      "properties": {
        "token": {
          "type": "string"
        },
        "username": {
          "$ref": "#/definitions/Username"
        }
      }
    },
    "Username": {
      "description": "User login.",
      "type": "string",
      "maxLength": 32,
      "minLength": 1
    }
  },
  "responses": {
    "GenericError": {
      "description": "Generic error response.",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "ListBrand": {
      "description": "Getting a list of brands.",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/Brand"
        }
      }
    },
    "ListProduct": {
      "description": "Receive a list of products.",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/Product"
        }
      }
    },
    "NoContent": {
      "description": "The server successfully processed the request and is not returning any content."
    },
    "OneProduct": {
      "description": "Getting one product.",
      "schema": {
        "$ref": "#/definitions/Product"
      }
    }
  },
  "securityDefinitions": {
    "apiKey": {
      "description": "Session token.",
      "type": "apiKey",
      "name": "AdminCookie",
      "in": "header"
    }
  },
  "security": [
    {
      "apiKey": null
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Service storage brands and cosmetics.",
    "title": "Service-storage",
    "version": "0.0.1"
  },
  "basePath": "/api",
  "paths": {
    "/brand": {
      "get": {
        "description": "Get list brand",
        "operationId": "brandList",
        "responses": {
          "200": {
            "description": "Getting a list of brands.",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Brand"
              }
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Create new brand",
        "operationId": "brandPOST",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/BrandCreate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Getting a list of brands.",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Brand"
              }
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete brand",
        "operationId": "brandDELETE",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The server successfully processed the request and is not returning any content."
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/login": {
      "post": {
        "security": [],
        "description": "Auth method.",
        "operationId": "login",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/LoginInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/UserInfo"
            },
            "headers": {
              "AdminCookie": {
                "type": "string",
                "description": "Session token."
              }
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/product": {
      "post": {
        "description": "Create new product",
        "operationId": "productPOST",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ProductCreate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Getting one product.",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete product",
        "operationId": "productDELETE",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The server successfully processed the request and is not returning any content."
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/product/{productID}": {
      "get": {
        "security": [],
        "operationId": "product",
        "parameters": [
          {
            "type": "integer",
            "description": "Track ID.",
            "name": "productID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Getting one product.",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/products": {
      "get": {
        "security": [],
        "operationId": "listProduct",
        "parameters": [
          {
            "type": "string",
            "name": "query",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "brand",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 25,
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 0,
            "name": "offset",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Receive a list of products.",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Product"
              }
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Brand": {
      "type": "object",
      "required": [
        "id",
        "name"
      ],
      "properties": {
        "id": {
          "$ref": "#/definitions/ID"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "BrandCreate": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string",
          "maxLength": 100,
          "minLength": 1
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ID": {
      "description": "Object ID.",
      "type": "integer"
    },
    "LoginInfo": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "$ref": "#/definitions/Password"
        },
        "username": {
          "$ref": "#/definitions/Username"
        }
      }
    },
    "Pagination": {
      "type": "object",
      "required": [
        "limit",
        "offset"
      ],
      "properties": {
        "limit": {
          "type": "integer",
          "maximum": 500,
          "minimum": 1
        },
        "offset": {
          "type": "integer",
          "minimum": 0
        }
      }
    },
    "Password": {
      "description": "User password.",
      "type": "string",
      "format": "password",
      "maxLength": 100,
      "minLength": 1
    },
    "Product": {
      "type": "object",
      "required": [
        "id",
        "name",
        "description",
        "apply",
        "price",
        "brandName"
      ],
      "properties": {
        "apply": {
          "type": "string"
        },
        "brandName": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "$ref": "#/definitions/ID"
        },
        "name": {
          "type": "string",
          "maxLength": 100,
          "minLength": 1
        },
        "price": {
          "type": "integer"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "ProductCreate": {
      "type": "object",
      "required": [
        "name",
        "description",
        "apply",
        "price",
        "brandID"
      ],
      "properties": {
        "apply": {
          "type": "string"
        },
        "brandID": {
          "type": "integer"
        },
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string",
          "maxLength": 100,
          "minLength": 1
        },
        "price": {
          "type": "integer"
        }
      }
    },
    "UserInfo": {
      "type": "object",
      "required": [
        "username",
        "token"
      ],
      "properties": {
        "token": {
          "type": "string"
        },
        "username": {
          "$ref": "#/definitions/Username"
        }
      }
    },
    "Username": {
      "description": "User login.",
      "type": "string",
      "maxLength": 32,
      "minLength": 1
    }
  },
  "responses": {
    "GenericError": {
      "description": "Generic error response.",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "ListBrand": {
      "description": "Getting a list of brands.",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/Brand"
        }
      }
    },
    "ListProduct": {
      "description": "Receive a list of products.",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/Product"
        }
      }
    },
    "NoContent": {
      "description": "The server successfully processed the request and is not returning any content."
    },
    "OneProduct": {
      "description": "Getting one product.",
      "schema": {
        "$ref": "#/definitions/Product"
      }
    }
  },
  "securityDefinitions": {
    "apiKey": {
      "description": "Session token.",
      "type": "apiKey",
      "name": "AdminCookie",
      "in": "header"
    }
  },
  "security": [
    {
      "apiKey": []
    }
  ]
}`))
}
