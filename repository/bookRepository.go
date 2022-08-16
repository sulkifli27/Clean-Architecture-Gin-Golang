package repository

import (
	"github/com/sulkifli27/golang_api/model"

	"gorm.io/gorm"
)


type BookRepository interface {
	InsertBook(b model.Book) model.Book
	UpdateBook(b model.Book) model.Book
	DeleteBook(b model.Book)
    AllBook() []model.Book
    FindBookById(bookID uint64) model.Book
}

type bookConnection struct {
    connection *gorm.DB
}

func NewBookRepository(dbConn *gorm.DB) BookRepository {
    return &bookConnection{
        connection: dbConn,
    }
}

func (db *bookConnection) InsertBook(b model.Book) model.Book{
    db.connection.Save(&b)
    db.connection.Preload("User").Find(&b)
    return b
}

func (db *bookConnection) UpdateBook(b model.Book) model.Book{
    db.connection.Save(&b)
    db.connection.Preload("User").Find(&b)
    return b
}

func (db *bookConnection) DeleteBook(b model.Book) {
    db.connection.Delete(&b)  
}

func (db *bookConnection) FindBookById(bookID uint64) model.Book{
    var book model.Book
    db.connection.Preload("User").Find(&book, bookID)
    return book
}

func (db *bookConnection) AllBook() []model.Book{
    var books []model.Book
    db.connection.Preload("User").Find(&books)
    return books
}


