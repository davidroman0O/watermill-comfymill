package comfymill

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill-sql/v3/pkg/sql"
)

type DefaultSQLite3OffsetsAdapter struct {
	// GenerateMessagesOffsetsTableName may be used to override how the messages/offsets table name is generated.
	GenerateMessagesOffsetsTableName func(topic string) string
}

func (a DefaultSQLite3OffsetsAdapter) SchemaInitializingQueries(topic string) []sql.Query {
	return []sql.Query{
		{
			Query: `
				CREATE TABLE IF NOT EXISTS ` + a.MessagesOffsetsTable(topic) + ` (
				consumer_group TEXT NOT NULL,
				offset_acked INTEGER,
				offset_consumed INTEGER NOT NULL,
				PRIMARY KEY(consumer_group)
			)`,
		},
	}
}

func (a DefaultSQLite3OffsetsAdapter) AckMessageQuery(topic string, row sql.Row, consumerGroup string) sql.Query {
	ackQuery := `INSERT INTO ` + a.MessagesOffsetsTable(topic) + ` (offset_consumed, offset_acked, consumer_group)
		VALUES (?, ?, ?) ON CONFLICT(consumer_group) DO UPDATE SET offset_consumed=excluded.offset_consumed, offset_acked=excluded.offset_acked`

	return sql.Query{ackQuery, []any{row.Offset, row.Offset, consumerGroup}}
}

func (a DefaultSQLite3OffsetsAdapter) NextOffsetQuery(topic, consumerGroup string) sql.Query {
	return sql.Query{
		Query: `SELECT COALESCE(
				(SELECT offset_acked
				 FROM ` + a.MessagesOffsetsTable(topic) + `
				 WHERE consumer_group=?
				), 0)`,
		Args: []any{consumerGroup},
	}
}

func (a DefaultSQLite3OffsetsAdapter) MessagesOffsetsTable(topic string) string {
	if a.GenerateMessagesOffsetsTableName != nil {
		return a.GenerateMessagesOffsetsTableName(topic)
	}
	return fmt.Sprintf("watermill_offsets_%s", topic)
}

func (a DefaultSQLite3OffsetsAdapter) ConsumedMessageQuery(topic string, row sql.Row, consumerGroup string, consumerULID []byte) sql.Query {
	ackQuery := `INSERT INTO ` + a.MessagesOffsetsTable(topic) + ` (offset_consumed, consumer_group)
		VALUES (?, ?) ON CONFLICT(consumer_group) DO UPDATE SET offset_consumed=excluded.offset_consumed`
	return sql.Query{ackQuery, []interface{}{row.Offset, consumerGroup}}
}

func (a DefaultSQLite3OffsetsAdapter) BeforeSubscribingQueries(topic, consumerGroup string) []sql.Query {
	return nil
}
