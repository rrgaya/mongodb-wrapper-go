# Mongodb Wrapper Go

[![GoDoc](https://godoc.org/github.com/rrgaya/mongodb-wrapper-go?status.svg)](https://godoc.org/github.com/rrgaya/mongodb-wrapper-go)

O pacote mwgo mongodb wrapper Go é um wrapper do mongo-go-driver que define as coleções e as APIs CRUD nelas.

```go 
    urlMongo := "mongodb://admin:password123@localhost:6000/"

    collection := mwgo.NewCollection(urlMongo, "your_database", "your_collections") 

    collection.InsertOne("")
```