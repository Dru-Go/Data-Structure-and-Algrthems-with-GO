package main

import (
	"context"
	"fmt"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionBuilder interface {
	ConnectToMongoDB(uri string) (Database, error)
	ConnectToRabbitMQ(uri string) (MessageQueue, error)
}

type Database interface {
	GetCollection(name string) Collection
	Close() error
}

type MessageQueue interface {
	Publish(exchange, routingKey string, body []byte) error
	Consume(queue, consumer string, autoAck bool, handler func(delivery amqp.Delivery)) error
	Close() error
}

type Collection interface {
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
}

type Cursor interface {
	Close(ctx context.Context) error
	Next(ctx context.Context) bool
	Decode(v interface{}) error
	Err() error
	All(ctx context.Context, results interface{}) error
}

type MongoDBConnection struct {
	client   *mongo.Client
	database *mongo.Database
}

func (c *MongoDBConnection) GetCollection(name string) Collection {
	return c.database.Collection(name)
}

func (c *MongoDBConnection) Close() error {
	return c.client.Disconnect(context.Background())
}

type RabbitMQConnection struct {
	channel  *amqp.Channel
	consumer string
}

func (c *RabbitMQConnection) Publish(exchange, routingKey string, body []byte) error {
	return c.channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		Body: body,
	})
}

func (c *RabbitMQConnection) Consume(queue, consumer string, autoAck bool, handler func(delivery amqp.Delivery)) error {
	deliveries, err := c.channel.Consume(queue, consumer, autoAck, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for delivery := range deliveries {
			handler(delivery)
		}
	}()
	return nil
}

func (c *RabbitMQConnection) Close() error {
	return c.channel.Close()
}

type ConnectionBuilderImpl struct{}

func (b *ConnectionBuilderImpl) ConnectToMongoDB(uri string) (Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return &MongoDBConnection{client: client, database: client.Database("test")}, nil
}

func (b *ConnectionBuilderImpl) ConnectToRabbitMQ(uri string) (MessageQueue, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMQConnection{channel: channel}, nil
}

func main() {
	builder := &ConnectionBuilderImpl{}
	db, err := builder.ConnectToMongoDB("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	mq, err := builder.ConnectToRabbitMQ("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mq.Close()
	// Use the connections here...
}

/**
In this example, the `ConnectionBuilder` interface defines two methods for connecting to a MongoDB database and a RabbitMQ message queue. The `MongoDBConnection` and `RabbitMQConnection` types implement the `Database` and `MessageQueue` interfaces, respectively, and provide the actual implementation for the connection. The `ConnectionBuilderImpl` type is a concrete builder that implements the `ConnectionBuilder` interface and provides the specific logic for connecting to the MongoDB and RabbitMQ systems. The `main` function uses an instance of the `ConnectionBuilderImpl` to create the connections to the MongoDB and RabbitMQ systems, and then uses these connections to perform various operations.
*/
