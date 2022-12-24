# API Documentation

## Users
#### Create User
```http request
POST /users
```
```json lines
Body Request
{
  "name" : "Zakaria"
}

Response
{
  "message": "create data successfully",
  "user": {
    "id": 1,
    "name": "Zakaria"
  }
}
```
#### Get All User
```http request
GET /users
```
```json lines
Response
[
  {
    "id": 1,
    "name": "Zakaria"
  }, 
  {
    "id": 2,
    "name": "Wahyu"
  }   
]
```
#### Get User By ID
```http request
GET /users/:id
```
```json lines
Response
{
  "id": 1,
  "name": "Zakaria"
}
```
#### Update User By ID
```http request
PUT /users/:id
```
#### Delete User BY ID
```http request
DELETE /users/:id
```

## Lockers
#### Create Locker
```http request
POST /lokers
```
```json lines
Body Request
{
  "user_id" : 1,
  "code" : "123"
}

Response
{
  "message": "create data successfully",
  "locker": {
    "id": 2,
    "code": "123",
    "user_id": 1
  },
}
```
#### Get All Locker
```http request
GET /lockers
```
```json lines
Response
[
  {
    "id": 1, 
    "code": "123",
    "user_id": 1,
    "user": {
      "id": 1,
      "name": "Zakaria"
    }
  }, 
  {
    "id": 2,
    "code": "456",
    "user_id": 2,
    "user": {
      "id": 2, 
      "name": "Wahyu"
    }
  }
]
```
#### Get Locker By ID
```http request
GET /lockers/:id
```
```json lines
Response
{
  "id": 1,
  "code": "123",
  "user_id": 1,
  "user": {
    "id": 1,
    "name": "Zakaria"
  }
}
```
#### Update Locker By ID
```http request
PUT /lockers/:id
```
#### Delete Locker By ID
```http request
DELETE /lockers/:id
```

## Tags
#### Create Tags
```http request
POST /tags
```
```json lines
Body Request
{
  "name" : "Jakarta"
}

Response
{
  "message": "create data successfully",
  "tag": {
    "id": 1,
    "name": "Jakarta"
  }
}
```
#### Get All Tags
```http request
GET /tags
```
```json lines
Response
[
  {
    "id": 1,
    "name": "Jakarta"
  },
  {
    "id": 2,
    "name": "Kota Kolaborasi"
  }
]
```
#### Get Tags By ID
```http request
GET /tags/:id
```
```json lines
Response
{
  "id": 2,
  "name": "Kota Kolaborasi"
}
```
#### Update Tags By ID
```http request
PUT /tags/:id
```
#### Delete Tags By ID
```http request
DELETE /tags/:id
```

## Post
#### Create Post
```http request
POST /posts
```
```json lines
Body Request
{
  "title" : "Berita 1",
  "body" : "Isi berita 1",
  "user_id" : 1,
  "tags_id" : [
    1,
    2,
    3
  ]
}

Response
{
  "message": "create data successfully",
  "post": {
    "id": 1,
    "title": "Berita 1",
    "body": "Isi berita 1",
    "user_id": 1,
    "tags_id": [
      1,
      2,
      3
    ]
  }
}
```
#### Get All Posts
```http request
GET /posts
```
```json lines
Response
[
  {
    "id": 1,
    "title": "Berita 2",
    "body": "Isi berita 2",
    "user_id": 1,
    "user": {
      "id": 1,
      "name": "Zakaria",
      "locker": {
        "id": 1,
        "code": "123",
        "user_id": 1
      }
    },
    "tags": [
      {
        "id": 1,
        "name": "Jakarta"
      },
      {
        "id": 2,
        "name": "Kota Kolaborasi"
      },
      {
        "id": 3,
        "name": "Metropolitan"
      }
    ]
  },
  {
    "id": 2,
    "title": "Berita Jakarta",
    "body": "Isi berita Jakarta",
    "user_id": 2,
    "user": {
      "id": 2,
      "name": "Wahyu",
      "locker": {
        "id": 2,
        "code": "456",
        "user_id": 2
      }
    },
    "tags": [
      {
        "id": 4,
        "name": "Banjir"
      },
      {
        "id": 5,
        "name": "BUMN"
      }
    ]
  }
]
```
#### Get Post By ID
```http request
GET /posts/:id
```
```json lines
Response
{
  "id": 11,
  "title": "Berita Jakarta",
  "body": "Isi berita Jakarta",
  "user_id": 2,
  "user": {
    "id": 2,
    "name": "Wahyu",
    "locker": {
      "id": 2,
      "code": "456",
      "user_id": 2
    }
  },
  "tags": [
    {
      "id": 4,
      "name": "Banjir"
    },
    {
      "id": 5,
      "name": "BUMN"
    }
  ]
}
```
#### Update Post By ID
```http request
PUT /posts/:id
```
### Delete Post By ID
```http request
DELETE /posts/:id
```

## Data Details
### Get All Details Users
##### Detail with Lockers, Posts and Tags
```http request
GET /users-details
```
```json lines
Response
[
  {
    "id": 1,
    "name": "Zakaria",
    "locker": {
      "id": 1,
      "code": "123",
      "user_id": 1
    },
    "post": [
      {
        "id": 1,
        "title": "Berita 2",
        "body": "Isi berita 2",
        "user_id": 1,
        "tags": [
          {
            "id": 1,
            "name": "Jakarta"
          },
          {
            "id": 2,
            "name": "Kota Kolaborasi"
          },
          {
            "id": 3,
            "name": "Metropolitan"
          }
        ]
      }
    ]
  },
  {
    "id": 2,
    "name": "Wahyu",
    "locker": {
      "id": 2,
      "code": "456",
      "user_id": 2
    },
    "post": [
      {
        "id": 11,
        "title": "Berita Jakarta",
        "body": "Isi berita Jakarta",
        "user_id": 2,
        "tags": [
          {
            "id": 4,
            "name": "Banjir"
          },
          {
            "id": 5,
            "name": "BUMN"
          }
        ]
      }
    ]
  }
]
```
### Get Details Users By ID
```http request
GET /users-details/:id
```
```json lines
Response
{
  "id": 1,
  "name": "Zakaria",
  "locker": {
    "id": 1,
    "code": "123",
    "user_id": 1
  },
  "post": [
    {
      "id": 1,
      "title": "Berita 2",
      "body": "Isi berita 2",
      "user_id": 1,
      "tags": [
        {
          "id": 1,
          "name": "Jakarta"
        },
        {
          "id": 2,
          "name": "Kota Kolaborasi"
        },
        {
          "id": 3,
          "name": "Metropolitan"
        }
      ]
    },
    {
      "id": 2,
      "title": "Berita Jakarta",
      "body": "Isi berita Jakarta",
      "user_id": 1,
      "tags": [
        {
          "id": 4,
          "name": "Banjir"
        },
        {
          "id": 5,
          "name": "BUMN"
        }
      ]
    }
  ]
}
```
### Get All Details Tags
#### Detail with Posts
```http request
GET /tags-details
```
```json lines
Response
[
  {
    "id": 1,
    "name": "Jakarta",
    "post": [
      {
        "id": 1,
        "title": "Berita 2",
        "body": "Isi berita 2"
      },
      {
        "id": 2,
        "title": "Berita 2",
        "body": "Isi berita 2"
      }
    ]
  },
  {
    "id": 2,
    "name": "Kota Kolaborasi",
    "post": [
      {
        "id": 1,
        "title": "Berita 2",
        "body": "Isi berita 2"
      },
      {
        "id": 2,
        "title": "Berita 2",
        "body": "Isi berita 2"
      }
    ]
  },
  {
    "id": 3,
    "name": "Metropolitan",
    "post": [
      {
        "id": 1,
        "title": "Berita 2",
        "body": "Isi berita 2"
      },
      {
        "id": 2,
        "title": "Berita 2",
        "body": "Isi berita 2"
      }
    ]
  },
  {
    "id": 4,
    "name": "Banjir",
    "post": [
      {
        "id": 10,
        "title": "Berita Jakarta",
        "body": "Isi berita Jakarta"
      },
      {
        "id": 11,
        "title": "Berita Jakarta",
        "body": "Isi berita Jakarta"
      }
    ]
  },
  {
    "id": 5,
    "name": "BUMN",
    "post": [
      {
        "id": 10,
        "title": "Berita Jakarta",
        "body": "Isi berita Jakarta"
      },
      {
        "id": 11,
        "title": "Berita Jakarta",
        "body": "Isi berita Jakarta"
      }
    ]
  }
]
```
### Get Details Tags By ID
```http request
GET /tags-details/:id
```
```json lines
Response
{
  "id": 4,
  "name": "Banjir",
  "post": [
    {
      "id": 10,
      "title": "Berita Jakarta",
      "body": "Isi berita Jakarta"
    },
    {
      "id": 11,
      "title": "Berita Jakarta",
      "body": "Isi berita Jakarta"
    }
  ]
}
```
