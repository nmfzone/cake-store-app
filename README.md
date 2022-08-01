# Cake Store App

## Requirements
- Go version >= 1.18
- Golang migrate (optional)
- Docker (optional)

## How to run
1. Execute `WEB_HOST_PORT=8080 MARIADB_HOST_PORT=3306 make run`.
Change port to another port if they're not available
2. Navigate to `http://localhost:8080`

## Available endpoints
#### [GET] /public/api/cakes
Used to get list of cakes. Example response:
```json
{
    "data": [
        {
            "id": 71220999374245888,
            "title": "Roti Bakar",
            "description": "It's delicious cakes from Indonesia",
            "rating": 100,
            "image": null,
            "created_at": "2022-07-17T07:00:00+07:00",
            "updated_at": "2022-07-17T07:00:00+07:00"
        },
        {
            "id": 71221185135775744,
            "title": "Roti Maryam",
            "description": null,
            "rating": 80,
            "image": null,
            "created_at": "2022-07-17T08:00:00+07:00",
            "updated_at": "2022-07-17T08:00:00+07:00"
        }
    ]
}
```

By default, it shows 10 cakes. You can change using `limit` query param to get
less or more cakes.

If the data in storage is more than limit, it will respond with `X-CURSOR`
header, so you can retrieve next data by passing those value to the `cursor` query param.

#### [POST] /public/api/cakes
Used to add cake to the storage. Example response:
```json
{
    "data": {
        "id": 71221185135775744,
        "title": "Roti Maryam",
        "description": null,
        "rating": 80,
        "image": null,
        "created_at": "2022-07-17T08:00:00+07:00",
        "updated_at": "2022-07-17T08:00:00+07:00"
    },
    "message": "Cake created successfully."
}
```

You need to send data using form-data, with `application/x-www-form-urlencoded`
content type.

#### [GET] /public/api/cakes/:id
Used to see detail of cake base on its ID. Example response:
```json
{
    "data": {
        "id": 71221185135775744,
        "title": "Roti Maryam",
        "description": null,
        "rating": 80,
        "image": null,
        "created_at": "2022-07-17T08:00:00+07:00",
        "updated_at": "2022-07-17T08:00:00+07:00"
    }
}
```

#### [PUT] /public/api/cakes/:id
Used to update cake base on its ID. Example response:
```json
{
    "data": {
        "id": 71220999374245888,
        "title": "Roti Bakar",
        "description": "It's delicious cakes from Indonesia",
        "rating": 100,
        "image": null,
        "created_at": "2022-07-17T08:00:00+07:00",
        "updated_at": "2022-07-17T08:02:00+07:00"
    },
    "message": "Cake updated successfully."
}
```

You need to send data using form-data, with `application/x-www-form-urlencoded`
content type.

#### [DELETE] /public/api/cakes/:id
Used to remove cake from storage base on its ID. Example response:
```json
{
    "message": "Cake deleted successfully."
}
```

## Test
Execute `make unittest` to run the test

## License
MIT
