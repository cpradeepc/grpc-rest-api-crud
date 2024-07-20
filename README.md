# grpc-REST API example application

This is application to understand the rest-api  crud operaton  with grpc service.

Some little steps before start it:

`Make gen` It generate  and make compiled proto file.

`Make del` It delet the generate compiled proto file.

# REST API

The REST API to the example app is described below.

## Person list in [][]byte

### Request

`POST api/person`

    http://localhost:8080/api/person
    Json Body:-
    {
    "name": "name-person",
    "city": "dl",
    "phone": 4444888822,
    "height": 6.5,
    "married": true
    }

### Response
    {
    "error": null,
    "response": {
        "id": 1,
        "name": "name-person",
        "city": "dl",
        "phone": 4444888822,
        "height": 6.5,
        "married": true
        }
    }




`GET api/person/id`

    http://localhost:8080/api/person/1
    

### Response
    {
    "error": null,
    "response": {
        "id": 1,
        "name": "name-person",
        "city": "dl",
        "phone": 4444888822,
        "height": 6.5,
        "married": true
        }
    }



`GET api/persons`

    http://localhost:8080/api/persons
    Json Body:-
    {
        "city": "dl",
    }

### Response
    {
    [
    "error": null,
    "response": {
        "id": 1,
        "name": "reemmeenaa",
        "city": "dl",
        "phone": 1111111111,
        "height": 7.1,
        "married": true
        },
        {
            "id": 2,
            "name": "john",
            "city": "dl",
            "phone": 8888777733,
            "height": 5.5,
            "married": true
        }
    ]
    }


`DELETE api/person/id`

    http://localhost:8080/api/person/2
    

### Response
    {
    "error": null,
    "response": "deleted:2"
    }
