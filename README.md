# Golang graphql example
This project uses [graph-gophers/graphql-go](github.com/graph-gophers/graphql-go), [graph-gophers/dataloader](github.com/graph-gophers/dataloader) to write a sample graphql server in golang.

## Sample queries
`POST: /graphql`

#### Find posts by ids
```graphql
{
    findPosts(ids: [1,11,21,31,41]){
        id
        body
        user{
            id
            name
        }
    }
}
```

#### Get fist N posts
```graphql
{
    listPosts(limit: 2){
        id
        body
        user{
            name
            address
        }
    }
}
```

#### Sample CURL
```shell
curl --location 'http://localhost:9000/graphql' \
--header 'Content-Type: application/json' \
--data '{"query":"{\n    listPosts(limit: 2){\n        id\n        body\n        user{\n            name\n            address\n        }\n    }\n}"}'
```

## Project structure
```
├── adapter // data source
├── repository // data source
├── config
├── graphql
│    ├── resolver
│    │    ├── post.go // model resolver
│    │    ├── root.go // root resolver
│    └── schema
│        ├── schema.go
│        ├── schema.graphql
│        └── type
│            ├── post.graphql // model type
│            ├── root.graphql // root type
├── handler.go
├── loader
│    ├── loader.go // dataloader
│    └── user.go // dataloader implementation
├── model // models
├── pkg // librabries
└── server.go

```

## Reference
- https://jsonplaceholder.typicode.com/