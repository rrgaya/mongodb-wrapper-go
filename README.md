# Mongodb Wrapper Go

![Github Actions](https://github.com/rrgaya/mongodb-wrapper-go/actions/workflows/mongodb.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/rrgaya/mongodb-wrapper-go?status.svg)](https://godoc.org/github.com/rrgaya/mongodb-wrapper-go)

O package mongodb wrapper Go é um wrapper do mongo-go-driver que define as coleções e as APIs CRUD nelas.

```golang
    package your_pagkage

    import (
        mwgo "github.com/rrgaya/mongodb-wrapper-go"
    )
    
    // MWGo NewColletion já faz trata o parse da url e também gerencianmento de context
    collection, err := mwgo.NewCollection("your_url_connection", "your_database", "your_collections") 
    if err != nil{
        log.Println("error on connections with mongodb.")
    }
    collection.InsertOne({})
```
