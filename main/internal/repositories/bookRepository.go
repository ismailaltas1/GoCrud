package repositories

import (
	"awesomeProject/main/internal/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"log"
)

const collectionName = "books"

type BooksRepository struct {
	db *mongo.Collection
}

//go:generate mockery --name IBookRepository
type IBookRepository interface {
	GetBooks(ctx context.Context) (b []models.Book, err error)
	CreateBooks(ctx context.Context, book models.Book) (err error)
	GetBookById(background context.Context, id string) (book models.Book, err error)
	UpdateBook(ctx context.Context, id string, book models.Book) (err error)
	DeleteBook(background context.Context, id string) (err error)
}

func NewBooksRepository(db *mongo.Database) IBookRepository {
	return &BooksRepository{
		db: db.Collection(collectionName),
	}
}

func (br *BooksRepository) GetBooks(ctx context.Context) (b []models.Book, err error) {

	cur, err := br.db.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal("Error get books connection")
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var books models.Book
		if err = cur.Decode(&books); err != nil {
			log.Fatal("Error get books connection")
		}
		b = append(b, books)
	}
	return b, err

}

func (br *BooksRepository) GetBookById(ctx context.Context, id string) (book models.Book, err error) {
	err = br.db.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	return book, err
}

func (br *BooksRepository) CreateBooks(ctx context.Context, book models.Book) (err error) {
	_, err = br.db.InsertOne(ctx, book)
	return
}

func (br *BooksRepository) UpdateBook(ctx context.Context, id string, book models.Book) (err error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}
	update := bson.M{
		"$set": bson.M{
			"title":  book.Title,
			"author": book.Author,
		},
	}
	filter := bson.M{"_id": objID}
	result, err := br.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no book found with the given id")
	}
	return nil

}
func (br *BooksRepository) DeleteBook(ctx context.Context, id string) (err error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}
	_, err = br.db.DeleteOne(ctx, bson.M{"_id": objID})
	return
}
