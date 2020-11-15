define({ "api": [
  {
    "type": "post",
    "url": "/auth/login",
    "title": "user login",
    "version": "1.0.0",
    "name": "login",
    "group": "auth",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>user email</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>user password</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>Bearer token</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>api success message</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "web/user.go",
    "groupTitle": "auth"
  },
  {
    "type": "post",
    "url": "/auth/register",
    "title": "user register",
    "version": "1.0.0",
    "name": "register",
    "group": "auth",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>fullname of user</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>user email</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>user password</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>Bearer token</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>api success message</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "web/user.go",
    "groupTitle": "auth"
  },
  {
    "type": "post",
    "url": "/encode",
    "title": "encode url",
    "version": "1.0.0",
    "name": "encode",
    "group": "link",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "link",
            "description": "<p>orginal link to be shorten</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "expire",
            "description": "<p>expire time compatible with all standard layouts</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "link",
            "description": "<p>export shorten link of url</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "web/link.go",
    "groupTitle": "link"
  },
  {
    "type": "post",
    "url": "/:hash/info",
    "title": "link info",
    "version": "1.0.0",
    "name": "info",
    "group": "link",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "URLParam",
            "optional": false,
            "field": "hash",
            "description": "<p>hash of encoded url</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "web/link.go",
    "groupTitle": "link"
  }
] });
