# Watermill Comfylite3 Pub/Sub

This is an unofficial `sqlite3` provider leveraging `comfylite3` for the [Watermill](https://watermill.io/) project.

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