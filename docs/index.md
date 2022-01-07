
# golang_project



<!--- If we have only one grouop/collection, then no need for the "ungrouped" heading -->

* [Auth](#auth)
    1. [Login](#1-login)
        * [Success](#i-example-request-success)
        * [Wrong credentials](#ii-example-request-wrong-credentials)
        * [Bad Request](#iii-example-request-bad-request)
    1. [Register](#2-register)
        * [Success](#i-example-request-success-1)
        * [Bad Request](#ii-example-request-bad-request)
        * [User Exists](#iii-example-request-user-exists)
* [Category](#category)
    1. [Create Category](#1-create-category)
        * [Success](#i-example-request-success-2)
        * [Duplicate Error](#ii-example-request-duplicate-error)
        * [Bad Request](#iii-example-request-bad-request-1)
    1. [Delete Category](#2-delete-category)
        * [Success](#i-example-request-success-3)
    1. [Get Categories](#3-get-categories)
        * [Success](#i-example-request-success-4)
    1. [Get Category By Id](#4-get-category-by-id)
        * [Success](#i-example-request-success-5)
        * [Not Found](#ii-example-request-not-found)
        * [Bad Request](#iii-example-request-bad-request-2)
    1. [Update Category](#5-update-category)
        * [Success](#i-example-request-success-6)
        * [Bad Request](#ii-example-request-bad-request-1)
        * [Not Found](#iii-example-request-not-found)
* [Movie](#movie)
    1. [Add Movie To Favourites](#1-add-movie-to-favourites)
        * [Success](#i-example-request-success-7)
        * [Duplicate Error](#ii-example-request-duplicate-error-1)
    1. [Create Movie](#2-create-movie)
        * [Success](#i-example-request-success-8)
        * [Bad Request](#ii-example-request-bad-request-2)
    1. [Delete Movie](#3-delete-movie)
        * [Success](#i-example-request-success-9)
    1. [Get Favourite Movies](#4-get-favourite-movies)
        * [Success](#i-example-request-success-10)
    1. [Get Movie By Id](#5-get-movie-by-id)
        * [Success](#i-example-request-success-11)
        * [Not Found](#ii-example-request-not-found-1)
    1. [Get Movies](#6-get-movies)
        * [Success](#i-example-request-success-12)
        * [Bad Request](#ii-example-request-bad-request-3)
    1. [Remove Movie From Favourites](#7-remove-movie-from-favourites)
        * [Success](#i-example-request-success-13)
    1. [Update Movie](#8-update-movie)
        * [Success](#i-example-request-success-14)
        * [Bad Request](#ii-example-request-bad-request-4)

--------



## Auth



### 1. Login



***Endpoint:***

```bash
Method: POST
Type: URLENCODED
URL: localhost:8080/users/login
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| username | test_user |  |
| password | pass |  |



***More example Requests/Responses:***


##### I. Example Request: Success



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| username | test_user |  |
| password | pass |  |



##### I. Example Response: Success
```js
{
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDE2MDQ3NDAsImlhdCI6MTY0MTU4MDc0MCwiaXNzIjoiZ29sYW5nX2FwaSIsInN1YiI6ImFjY2Vzc190b2tlbiIsImlkIjoyfQ.N7Tj05Cglv-BLd_IHXpDsAERCM02UqkXbOPkU3QwDTQ"
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Wrong credentials



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| username | test_user |  |
| password | pass1 |  |



##### II. Example Response: Wrong credentials
```js
{
    "message": "Login failed"
}
```


***Status Code:*** 401

<br>



##### III. Example Request: Bad Request



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| username | test_user |  |



##### III. Example Response: Bad Request
```js
{
    "message": "Data is not valid"
}
```


***Status Code:*** 400

<br>



### 2. Register



***Endpoint:***

```bash
Method: POST
Type: URLENCODED
URL: localhost:8080/users/register
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| username | test_user |  |
| password | pass |  |
| email | testemail@gmail.com |  |



***More example Requests/Responses:***


##### I. Example Request: Success



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| username | test_user |  |
| password | pass |  |
| email | testemail@gmail.com |  |



##### I. Example Response: Success
```js
{
    "message": "success",
    "status": 200
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Bad Request



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| username | test_user |  |
| password | pass |  |



##### II. Example Response: Bad Request
```js
{
    "message": "Data is not valid"
}
```


***Status Code:*** 400

<br>



##### III. Example Request: User Exists



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| username | test_user |  |
| password | pass |  |
| email | testemail@gmail.com |  |



##### III. Example Response: User Exists
```js
{
    "message": "User exists"
}
```


***Status Code:*** 409

<br>



## Category



### 1. Create Category



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: localhost:8080/categories
```



***Body:***

```js        
{
    "name": "A New Category",
    "description": "A New Description"
}
```



***More example Requests/Responses:***


##### I. Example Request: Success



***Body:***

```js        
{
    "name": "A New Category",
    "description": "A New Description"
}
```



##### I. Example Response: Success
```js
{
    "data": {
        "id": 4,
        "name": "A New Category",
        "description": "A New Description"
    }
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Duplicate Error



***Body:***

```js        
{
    "name": "A New Category",
    "description": "A New Description"
}
```



##### II. Example Response: Duplicate Error
```js
{
    "message": "Category exists"
}
```


***Status Code:*** 409

<br>



##### III. Example Request: Bad Request



***Body:***

```js        
{
    "name": "A New Category"
}
```



##### III. Example Response: Bad Request
```js
{
    "message": "Data is not valid"
}
```


***Status Code:*** 400

<br>



### 2. Delete Category



***Endpoint:***

```bash
Method: DELETE
Type: 
URL: localhost:8080/categories/3
```



***More example Requests/Responses:***


##### I. Example Request: Success



##### I. Example Response: Success
```js
{
    "data": {
        "id": 3,
        "created_at": "2022-01-07T20:55:20.970061+03:00",
        "updated_at": "2022-01-07T20:55:20.970061+03:00",
        "is_deleted": true,
        "deleted_at": {
            "Time": "2022-01-07T21:56:33.971136+03:00",
            "Valid": true
        },
        "name": "Horror",
        "description": "Example description"
    }
}
```


***Status Code:*** 200

<br>



### 3. Get Categories



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:8080/categories
```



***More example Requests/Responses:***


##### I. Example Request: Success



##### I. Example Response: Success
```js
{
    "data": [
        {
            "id": 1,
            "name": "Adventure",
            "description": "Example description"
        },
        {
            "id": 2,
            "name": "Sci-Fi",
            "description": "Example description"
        },
        {
            "id": 3,
            "name": "Horror",
            "description": "Example description"
        }
    ]
}
```


***Status Code:*** 200

<br>



### 4. Get Category By Id



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:8080/categories/1
```



***More example Requests/Responses:***


##### I. Example Request: Success



##### I. Example Response: Success
```js
{
    "data": {
        "id": 1,
        "name": "Adventure",
        "description": "Example description"
    }
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Not Found



##### II. Example Response: Not Found
```js
{
    "message": "Resource not found"
}
```


***Status Code:*** 404

<br>



##### III. Example Request: Bad Request



##### III. Example Response: Bad Request
```js
{
    "message": "Bad Request"
}
```


***Status Code:*** 400

<br>



### 5. Update Category



***Endpoint:***

```bash
Method: PUT
Type: RAW
URL: localhost:8080/categories/1
```



***Body:***

```js        
{
    "name": "Updated Category",
    "description": "Updated Description"
}
```



***More example Requests/Responses:***


##### I. Example Request: Success



***Body:***

```js        
{
    "name": "Updated Category",
    "description": "Updated Description"
}
```



##### I. Example Response: Success
```js
{
    "data": {
        "id": 1,
        "name": "Updated Category",
        "description": "Updated Description"
    }
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Bad Request



***Body:***

```js        
{
    "name": "Updated Category"
}
```



##### II. Example Response: Bad Request
```js
{
    "message": "Data is not valid"
}
```


***Status Code:*** 400

<br>



##### III. Example Request: Not Found



***Body:***

```js        
{
    "name": "Updated Category",
    "description": "Updated Description"
}
```



##### III. Example Response: Not Found
```js
{
    "message": "Resource not found"
}
```


***Status Code:*** 404

<br>



## Movie



### 1. Add Movie To Favourites



***Endpoint:***

```bash
Method: POST
Type: 
URL: localhost:8080/movies/id/1/favourites
```



***More example Requests/Responses:***


##### I. Example Request: Success



##### I. Example Response: Success
```js
{
    "data": true
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Duplicate Error



##### II. Example Response: Duplicate Error
```js
{
    "message": "Movie is already in favourites"
}
```


***Status Code:*** 400

<br>



### 2. Create Movie



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: localhost:8080/movies
```



***Body:***

```js        
{
    "name": "A New Adventure Movie",
    "plot": "A New plot",
    "categories": [1, 2]
}
```



***More example Requests/Responses:***


##### I. Example Request: Success



***Body:***

```js        
{
    "name": "A New Adventure Movie",
    "plot": "A New plot",
    "categories": [1]
}
```



##### I. Example Response: Success
```js
{
    "data": {
        "id": 10,
        "created_at": "2022-01-07T21:59:56.736353+03:00",
        "updated_at": "2022-01-07T21:59:56.736353+03:00",
        "is_deleted": false,
        "deleted_at": {
            "Time": "0001-01-01T00:00:00Z",
            "Valid": false
        },
        "name": "A New Adventure Movie",
        "plot": "A New plot",
        "categories": null
    }
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Bad Request



***Body:***

```js        
{
    "name": "A New Adventure Movie",
    "plot": "A New plot",
    "categories": [1, 1]
}
```



##### II. Example Response: Bad Request
```js
{
    "message": "Data is not valid"
}
```


***Status Code:*** 400

<br>



### 3. Delete Movie



***Endpoint:***

```bash
Method: DELETE
Type: 
URL: localhost:8080/movies/id/1
```



***More example Requests/Responses:***


##### I. Example Request: Success



##### I. Example Response: Success
```js
{
    "data": {
        "id": 1,
        "created_at": "2022-01-07T20:55:20.970061+03:00",
        "updated_at": "2022-01-07T20:55:20.970061+03:00",
        "is_deleted": true,
        "deleted_at": {
            "Time": "2022-01-07T22:03:56.716373+03:00",
            "Valid": true
        },
        "name": "Updated Movie Name",
        "plot": "Updated Movie Plot",
        "categories": null
    }
}
```


***Status Code:*** 200

<br>



### 4. Get Favourite Movies



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:8080/movies/favourites
```



***More example Requests/Responses:***


##### I. Example Request: Success



##### I. Example Response: Success
```js
{
    "data": [
        {
            "id": 7,
            "name": "A horror movie",
            "plot": "movie plot",
            "categories": [
                {
                    "id": 3,
                    "name": "Horror"
                }
            ]
        },
        {
            "id": 8,
            "name": "A adventure sci-fi movie 2",
            "plot": "movie plot",
            "categories": [
                {
                    "id": 2,
                    "name": "Sci-Fi"
                },
                {
                    "id": 1,
                    "name": "Updated Category"
                }
            ]
        }
    ]
}
```


***Status Code:*** 200

<br>



### 5. Get Movie By Id



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:8080/movies/id/1
```



***More example Requests/Responses:***


##### I. Example Request: Success



##### I. Example Response: Success
```js
{
    "data": {
        "id": 8,
        "name": "A adventure sci-fi movie 2",
        "plot": "movie plot",
        "categories": [
            {
                "id": 2,
                "name": "Sci-Fi",
                "description": "Example description"
            },
            {
                "id": 1,
                "name": "Updated Category",
                "description": "Updated Description"
            }
        ]
    }
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Not Found



##### II. Example Response: Not Found
```js
{
    "message": "Resource not found"
}
```


***Status Code:*** 404

<br>



### 6. Get Movies



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:8080/categories/1/movies
```



***More example Requests/Responses:***


##### I. Example Request: Success



##### I. Example Response: Success
```js
{
    "data": [
        {
            "id": 1,
            "name": "An adventure movie",
            "plot": "movie plot"
        },
        {
            "id": 2,
            "name": "An adventure movie 2",
            "plot": "movie plot"
        },
        {
            "id": 3,
            "name": "An adventure movie 3",
            "plot": "movie plot"
        },
        {
            "id": 8,
            "name": "A adventure sci-fi movie 2",
            "plot": "movie plot"
        }
    ]
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Bad Request



##### II. Example Response: Bad Request
```js
{
    "message": "Bad Request"
}
```


***Status Code:*** 400

<br>



### 7. Remove Movie From Favourites



***Endpoint:***

```bash
Method: DELETE
Type: 
URL: localhost:8080/movies/id/1/favourites
```



***More example Requests/Responses:***


##### I. Example Request: Success



##### I. Example Response: Success
```js
{
    "data": true
}
```


***Status Code:*** 200

<br>



### 8. Update Movie



***Endpoint:***

```bash
Method: PUT
Type: RAW
URL: localhost:8080/movies/id/1
```



***Body:***

```js        
{
    "name": "Updated Movie Name",
    "plot": "Updated Movie Plot"
}
```



***More example Requests/Responses:***


##### I. Example Request: Success



***Body:***

```js        
{
    "name": "Updated Movie Name",
    "plot": "Updated Movie Plot"
}
```



##### I. Example Response: Success
```js
{
    "data": {
        "id": 1,
        "created_at": "2022-01-07T20:55:20.970061+03:00",
        "updated_at": "2022-01-07T20:55:20.970061+03:00",
        "is_deleted": false,
        "deleted_at": {
            "Time": "0001-01-01T00:00:00Z",
            "Valid": false
        },
        "name": "Updated Movie Name",
        "plot": "Updated Movie Plot",
        "categories": null
    }
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Bad Request



***Body:***

```js        
{
    "name": "Updated Movie Name"
}
```



##### II. Example Response: Bad Request
```js
{
    "message": "Data is not valid"
}
```


***Status Code:*** 400

<br>



---
[Back to top](#golang_project)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2022-01-07 23:08:32 by [docgen](https://github.com/thedevsaddam/docgen)
