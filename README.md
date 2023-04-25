# Golang graphql example
This project uses [graphql-go](github.com/graph-gophers/graphql-go), [dataloader](github.com/graph-gophers/dataloader) to write a sample graphql server in golang.

Here, the DataLoader helps us prevent N+1 queries.

## Sample queries
#### Find posts by ids
```shell
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
//TODO: Why doesn't server batch `GetUser` requests
```shell
{
    listPosts(limit: 100){
        id
        body
        user{
            name
        }
    }
}
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