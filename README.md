# Hello-GraphQL

> My hello-world practice using [GraphQL](https://graphql.org) with [Golang](https://golang.org)

![hello-graphql](assets/graphql.png)

## Content

<!-- TOC -->

- [Hello-GraphQL](#hello-graphql)
  - [Content](#content)
  - [Introduction](#introduction)
  - [Libraries](#libraries)
    - [graphql-go/graphql](#graphql-gographql)
    - [graph-gophers/graphql-go](#graph-gophersgraphql-go)
    - [thunder](#thunder)
    - [gqlgen](#gqlgen)

<!-- /TOC -->

## Introduction

I'm tired of using `RESTful` api in `Golang` since I have to mangage which field to return in a big struct manually over and over again. And then I found [GraphQL](https://graphql.org) which really seems to solve my problem. So I create this repo and try some _hello-world_ practice for using `GraphQL` in `Golang`.

There are several libraries to help developing a `GraphQL` server in `Golang`:

1. [graphql-go/graphql](https://github.com/graphql-go/graphql)
2. [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go)
3. [thunder](https://github.com/samsarahq/thunder)
4. [gqlgen](https://github.com/vektah/gqlgen)

## Libraries

> comparing reference:
>
> [Medium](https://medium.com/open-graphql/choosing-a-graphql-server-library-in-go-8836f893881b)
>
> [99designs](https://99designs.com/blog/engineering/gqlgen-a-graphql-server-generator-for-go/)

### graphql-go/graphql

It says:

> Follows the official reference implementation [graphql-js](https://github.com/graphql/graphql-js).

The resolver looks like this:

```golang
func(p graphql.ResolveParams) (interface{}, error) {
    return "Hello, world!", nil
}
```

Apperantly that there are ~~many~~ huge numbers of `interface{}` which relay on _reflect_ and _type assert_.

And also you have to define the _query_ and _mutation_ as well as the _models_ in your `Go` codes by `graphql.{someType}`, which is duplicated and boring.

But it should be a good way to learn basic `GraphQL` _types_, _query_ and _mutation_.

One of the seniors I know used this library `=>` [Fredliang](https://github.com/fredliang44/family-tree)

### graph-gophers/graphql-go

This one looks like this:

```golang
func (q *query) Hello() string {
    return "Hello, world!"
}
```

Seems great, I will give it a try later.

### thunder

This one is a reflection-based approach, it looks like this:

```golang
object.FieldFunc("posts", func() []post {
    return posts
})
```

I'm doubt about it's performance.

### gqlgen

This is a generator. Only thing you need to do is writing *resolvers*, which like this:

```golang
func (a *MyApp) Query_todos(ctx context.Context) ([]Todo, error) {
    return a.todos, nil
}
```

I'm now tring this. I will try out whether the *resolvers* can work when I `add/delete/update` *schema*.