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
  "swagger": "2.0",
  "info": {
    "description": "A mini project for practice",
    "title": "Inventory Management",
    "contact": {
      "name": "AAA",
      "url": "http://aaa.com"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/inventory": {
      "get": {
        "tags": [
          "Stock"
        ],
        "summary": "fetches all inventory data",
        "responses": {
          "200": {
            "description": "Fetched all records",
            "schema": {
              "$ref": "#/definitions/Items"
            }
          }
        }
      },
      "post": {
        "tags": [
          "Stock"
        ],
        "summary": "saves an item to inventory",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Item saved successfully",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        }
      },
      "patch": {
        "tags": [
          "Stock"
        ],
        "summary": "update an item",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully updated the record",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        }
      }
    },
    "/inventory/search/{itemName}": {
      "get": {
        "tags": [
          "Stock"
        ],
        "summary": "get item by name",
        "responses": {
          "200": {
            "description": "Item fetched",
            "schema": {
              "$ref": "#/definitions/Items"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "itemName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/inventory/{itemId}": {
      "delete": {
        "tags": [
          "Stock"
        ],
        "summary": "delete an item",
        "responses": {
          "200": {
            "description": "Item deleted"
          },
          "404": {
            "description": "Item not found"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "itemId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/login": {
      "post": {
        "tags": [
          "Authorization"
        ],
        "summary": "login user",
        "parameters": [
          {
            "type": "string",
            "name": "email",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "password",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "login sucessful"
          },
          "401": {
            "description": "invalid credentials"
          }
        }
      }
    },
    "/user": {
      "put": {
        "tags": [
          "User"
        ],
        "summary": "user updated",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Users"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User Updated",
            "schema": {
              "$ref": "#/definitions/Users"
            }
          }
        }
      },
      "post": {
        "tags": [
          "User"
        ],
        "summary": "new user registration to inventory",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Users"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User created",
            "schema": {
              "$ref": "#/definitions/Users"
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "tags": [
          "User"
        ],
        "responses": {
          "200": {
            "description": "get user details",
            "schema": {
              "$ref": "#/definitions/Users"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "User"
        ],
        "responses": {
          "200": {
            "description": "user deleted"
          },
          "404": {
            "description": "internal server error"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Item": {
      "type": "object",
      "properties": {
        "MRP": {
          "type": "number",
          "example": 130
        },
        "description": {
          "type": "string",
          "example": "Bru"
        },
        "id": {
          "type": "integer",
          "example": 1
        },
        "itemName": {
          "type": "string",
          "example": "Coffee"
        },
        "quantity": {
          "type": "integer",
          "example": 6
        },
        "sellingPrice": {
          "type": "number",
          "example": 125.5
        }
      }
    },
    "Items": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Item"
      }
    },
    "Users": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "example": "rk@gmail.com"
        },
        "id": {
          "type": "integer",
          "example": 1
        },
        "name": {
          "type": "string",
          "example": "Ram"
        },
        "password": {
          "type": "string",
          "format": "password"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about user",
      "name": "User"
    },
    {
      "description": "operations on stock items",
      "name": "Stock"
    },
    {
      "description": "check user credentials",
      "name": "Authorization"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "A mini project for practice",
    "title": "Inventory Management",
    "contact": {
      "name": "AAA",
      "url": "http://aaa.com"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/inventory": {
      "get": {
        "tags": [
          "Stock"
        ],
        "summary": "fetches all inventory data",
        "responses": {
          "200": {
            "description": "Fetched all records",
            "schema": {
              "$ref": "#/definitions/Items"
            }
          }
        }
      },
      "post": {
        "tags": [
          "Stock"
        ],
        "summary": "saves an item to inventory",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Item saved successfully",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        }
      },
      "patch": {
        "tags": [
          "Stock"
        ],
        "summary": "update an item",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully updated the record",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        }
      }
    },
    "/inventory/search/{itemName}": {
      "get": {
        "tags": [
          "Stock"
        ],
        "summary": "get item by name",
        "responses": {
          "200": {
            "description": "Item fetched",
            "schema": {
              "$ref": "#/definitions/Items"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "itemName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/inventory/{itemId}": {
      "delete": {
        "tags": [
          "Stock"
        ],
        "summary": "delete an item",
        "responses": {
          "200": {
            "description": "Item deleted"
          },
          "404": {
            "description": "Item not found"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "itemId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/login": {
      "post": {
        "tags": [
          "Authorization"
        ],
        "summary": "login user",
        "parameters": [
          {
            "type": "string",
            "name": "email",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "password",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "login sucessful"
          },
          "401": {
            "description": "invalid credentials"
          }
        }
      }
    },
    "/user": {
      "put": {
        "tags": [
          "User"
        ],
        "summary": "user updated",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Users"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User Updated",
            "schema": {
              "$ref": "#/definitions/Users"
            }
          }
        }
      },
      "post": {
        "tags": [
          "User"
        ],
        "summary": "new user registration to inventory",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Users"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User created",
            "schema": {
              "$ref": "#/definitions/Users"
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "tags": [
          "User"
        ],
        "responses": {
          "200": {
            "description": "get user details",
            "schema": {
              "$ref": "#/definitions/Users"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "User"
        ],
        "responses": {
          "200": {
            "description": "user deleted"
          },
          "404": {
            "description": "internal server error"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Item": {
      "type": "object",
      "properties": {
        "MRP": {
          "type": "number",
          "example": 130
        },
        "description": {
          "type": "string",
          "example": "Bru"
        },
        "id": {
          "type": "integer",
          "example": 1
        },
        "itemName": {
          "type": "string",
          "example": "Coffee"
        },
        "quantity": {
          "type": "integer",
          "example": 6
        },
        "sellingPrice": {
          "type": "number",
          "example": 125.5
        }
      }
    },
    "Items": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Item"
      }
    },
    "Users": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "example": "rk@gmail.com"
        },
        "id": {
          "type": "integer",
          "example": 1
        },
        "name": {
          "type": "string",
          "example": "Ram"
        },
        "password": {
          "type": "string",
          "format": "password"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about user",
      "name": "User"
    },
    {
      "description": "operations on stock items",
      "name": "Stock"
    },
    {
      "description": "check user credentials",
      "name": "Authorization"
    }
  ]
}`))
}
