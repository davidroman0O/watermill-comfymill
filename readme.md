# Watermill Comfylite3 Pub/Sub

This is an unofficial `sqlite3` provider leveraging [`comfylite3`](https://github.com/davidroman0O/comfylite3) for the [Watermill](https://watermill.io/) project.

Since there are a lot of issues using `sqlite3` for either in memory or with a file when used in concurrency situations, `comfylite3` provide a while bypass for us.

All you need to use `comfymill.NewDatabase`, use your own parameters (check `comfylite3` docs) and tadam!

It's a simple implementation of the MySql code for Sqlite3, here an example usage of it:

```go

db, _ := comfymill.NewDatabase(
    comfylite3.WithMemory(),
)

subscriber, err := sql.NewSubscriber(
    db,
    sql.SubscriberConfig{
        SchemaAdapter:    comfymill.DefaultSQLite3Schema{},
        OffsetsAdapter:   comfymill.DefaultSQLite3OffsetsAdapter{},
        InitializeSchema: true,
    },
    logger,
)
if err != nil {
    panic(err)
}

messages, err := subscriber.Subscribe(context.Background(), "example_topic")
if err != nil {
    panic(err)
}

go process(messages)

publisher, err := sql.NewPublisher(
    db,
    sql.PublisherConfig{
        SchemaAdapter: comfymill.DefaultSQLite3Schema{},
    },
    logger,
)
if err != nil {
    panic(err)
}

```

You can refer to the [basic example too](./_example/basic/main.go)

