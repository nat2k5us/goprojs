package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/swaggo/cli"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	credential := options.Credential{
		Username: "root",
		Password: "7CLtPeEdm9",
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/").SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("booksarefun").Collection("books")
}

// You will be using this Trainer type later in the program
type Book struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Title       string             `bson:"text"`
	Author      string             `bson:"text"`
	Description string             `bson:"text"`
	Completed   bool               `bson:"completed"`
}

func createBook(book *Book) error {

	ins := bson.M{
		"ID":          book.ID,
		"CreatedAt":   book.CreatedAt,
		"UpdatedAt":   book.UpdatedAt,
		"Title":       book.Title,
		"Author":      book.Author,
		"Description": book.Description,
		"Completed":   book.Completed,
	}
	_, err1 := collection.InsertOne(ctx, ins)
	return err1
}

func getAll() ([]*Book, error) {
	// passing bson.D{{}} matches all documents in the collection
	filter := bson.D{{}}
	return filterBooks(filter)
}
func getAllBooks() {
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var epi bson.M
		if err = cur.Decode(&epi); err != nil {
			log.Fatal(err)
		}
		s, _ := json.MarshalIndent(epi, "", "\t")
		fmt.Print(string(s))

	}

}

func filterBooks(filter interface{}) ([]*Book, error) {
	// A slice of books for storing the decoded documents
	var books []*Book

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return books, err
	}

	for cur.Next(ctx) {
		var t Book
		err := cur.Decode(&t)
		if err != nil {
			return books, err
		}

		books = append(books, &t)
	}

	if err := cur.Err(); err != nil {
		return books, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(books) == 0 {
		return books, mongo.ErrNoDocuments
	}

	return books, nil
}

func printBooks(books []*Book) {
	for i, v := range books {
		if v.Completed {
			fmt.Printf("%d: %s %s %s\n", i+1, v.Title, v.Author, v.Description)
		} else {
			fmt.Printf("%d: %s %s %s\n", i+1, v.Title, v.Author, v.Description)
		}
	}
}

var (
	NAME    = "BookAreFun"
	VERSION = "null"
	BUILD   = "null"
)

func main() {

	app := &cli.App{
		Name:        NAME,
		Version:     VERSION + "." + BUILD,
		Description: "BookAreFun - A Book Library",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "debug, d"},
		},
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		fmt.Println("quit app...")

		select {
		case <-time.After(time.Second * 3):
			log.Fatal("quit app timeout")
		}
	}()

	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			log.Fatal("set debug level")
		}

		return nil
	}

	app.Commands = append(app.Commands, cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a book to the list",
		Action: func(c *cli.Context) error {
			str := c.Args().First()
			author := c.Args().Get(1)
			description := c.Args().Get(2)
			if str == "" {
				return errors.New("Cannot add an empty book")
			}

			book := &Book{
				ID:          primitive.NewObjectID(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       str,
				Author:      author,
				Description: description,
				Completed:   false,
			}

			return createBook(book)
		},
	})

	app.Commands = append(app.Commands, cli.Command{
		Name:    "getall",
		Aliases: []string{"l"},
		Usage:   "list all books",
		Action: func(c *cli.Context) error {
			getAllBooks()

			return nil
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal("app exit with error: %s", err)
		os.Exit(1)
	}

}

func ConnectMongo1() {
	// mongodb://<username>:<password>@<server_address>:<port>/<database_name>
	// Set client options
	credential := options.Credential{
		Username: "root",
		Password: "qdvPhmeS1Z",
	}
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017").SetAuth(credential)
	//clientOptions := options.Client().ApplyURI("mongodb://admin:password@18.207.2.160:27017/GoLogin")
	//shard 2
	//clientOptions := options.Client().ApplyURI("mongodb://admin:admin@54.161.239.229:27017/admin")
	//Shard 1
	//clientOptions := options.Client().ApplyURI("mongodb://admin:admin@18.206.233.147:27017/admin")
	//Router
	//clientOptions := options.Client().ApplyURI("mongodb://admin:admin@100.27.3.107:27017/admin")
	//working locally
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//try on docker locally
	//clientOptions := options.Client().ApplyURI("mongodb://t800:t800t800@ds041248.mlab.com:41248/imageservice")
	//clientOptions := options.Client().ApplyURI("mongodb://admin:admin@127.0.0.1/GoLogin:27017")
	//clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0-4zfux.mongodb.net/GoLogin?retryWrites=true&w=majority")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func ConnectMongo() {

	var (
		client   *mongo.Client
		mongoURL = "mongodb://localhost:27017"
	)

	// Initialize a new mongo client with options
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))

	// Connect the mongo client to the MongoDB server
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("could not ping to mongo db service: %v\n", err)
		return
	}

	fmt.Println("connected to nosql database:", mongoURL)

	collection := client.Database("tasker").Collection("books")

	fmt.Println(collection)

}
