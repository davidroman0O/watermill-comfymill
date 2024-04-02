# Watermill SQLite3 Pub/Sub

This is an unofficial `sqlite3` provider for the [Watermill](https://watermill.io/) project.

It's a simple implementation of the MySql code for Sqlite3, here an example usage of it:

```go

subscriber, err := sql.NewSubscriber(
    db,
    sql.SubscriberConfig{
        SchemaAdapter:    DefaultSQLite3Schema{},
        OffsetsAdapter:   DefaultSQLite3OffsetsAdapter{},
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
        SchemaAdapter: DefaultSQLite3Schema{},
    },
    logger,
)
if err != nil {
    panic(err)
}

```