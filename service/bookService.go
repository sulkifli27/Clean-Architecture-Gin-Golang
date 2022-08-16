package service

import (
	"fmt"
	"github/com/sulkifli27/golang_api/dto"
	"github/com/sulkifli27/golang_api/model"
	"github/com/sulkifli27/golang_api/repository"
	"log"

	"github.com/mashingan/smapping"
)


type BookService interface {
	Insert(b dto.BookCreateDTO) model.Book
	Update(b dto.BookUpdateDTO) model.Book
    Delete(b model.Book)
    All() []model.Book
    FindByID(bookID uint64) model.Book
    IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
    bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
    return &bookService{
        bookRepository : bookRepo,
    }
}

func (service *bookService) Insert(b dto.BookCreateDTO) model.Book {
    book := model.Book{}
    err := smapping.FillStruct(&book, smapping.MapFields(&b))
    if err != nil {
        log.Fatal("Failed map %v: ", err )
    }
    res := service.bookRepository.InsertBook(book)
    return res  
}

func (service *bookService) Update(b dto.BookUpdateDTO) model.Book {
    book := model.Book{}
    err := smapping.FillStruct(&book, smapping.MapFields(&b))
    if err != nil {
        log.Fatal("Failed map %v: ", err )
    }
    res := service.bookRepository.UpdateBook(book)
    return res  
}

func (service *bookService) Delete(b model.Book){
    service.bookRepository.DeleteBook(b)
}

func (service *bookService) All() []model.Book {
    return service.bookRepository.AllBook()
}

func (service *bookService) FindByID(bookID uint64) model.Book {
    return service.bookRepository.FindBookById(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
    b := service.bookRepository.FindBookById(bookID)
    id := fmt.Sprintf("%v", b.UserID)
    return userID == id
}
