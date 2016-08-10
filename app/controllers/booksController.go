package controllers

import (
	"github.com/revel/revel"
	"github.com/newisfa/test/app/util"
	"log"
	"github.com/newisfa/test/app/encoders"
	"github.com/newisfa/test/app"
	"github.com/newisfa/test/app/models"
	"strconv"
)

type BooksController struct {
	*revel.Controller
}

func (c BooksController) Create() revel.Result{
	var book = encoders.EncodeBooks(c.Request.Body)


	if book.Name == "" {
		return c.RenderJson(util.ResponseError("Fadlan buug cusub gali"))
	}

	//var id int
	//c.Params.Bind(&id, "id")
	book.AuthorID,_ = strconv.ParseInt(c.Session["id"], 10, 0)
	log.Println("AuthorID:", book.AuthorID)
	book.PublisherID,_ = strconv.ParseInt(c.Session["id"], 10, 0)
	log.Println("PublisherID:", book.PublisherID)
	book.RiviewID,_ = strconv.ParseInt(c.Session["id"], 10, 0)
	log.Println("RiviewID:", book.RiviewID)

	if err := app.Db.Create(&book).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Book Creation Fialed"))
	}

	return c.RenderJson(util.ResponseSuccess(book))
}

func (c BooksController) Update() revel.Result  {
	var update = encoders.EncodeBooks(c.Request.Body)
	var book models.Books
	var id int
	c.Params.Bind(&id, "id")
	if rowsCout := app.Db.First(&book, id).RowsAffected; rowsCout < 1 {
		return c.RenderJson(util.ResponseError("Book Update information Fialed"))
	}

	if err := app.Db.Model(&book).Updates(&update).Error; err != nil{
		return c.RenderJson(util.ResponseError("Book updates Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(update))
}

func (c BooksController) Delete() revel.Result {
	var (
		id int
		book models.Books
	)

	c.Params.Bind(&id, "id")
	if rowsCount := app.Db.First(&book, id).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("Author information Not Founded"))
	}

	if err := app.Db.Delete(&book).Error; err != nil {
		return c.RenderJson(util.ResponseError("Author Delete Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(book))
}

func (c BooksController) Get()revel.Result {
	var books []models.Books
	var limitQuery = c.Request.URL.Query().Get("limt");
	if limitQuery == "" {
		limitQuery = "0"
	}
	var offsetQuery = c.Request.URL.Query().Get("offset");


	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&books).RowsAffected; founded < 1 {
		c.RenderJson(util.ResponseError("Not founded authors"))
	}

	for i , book := range books {
		app.Db.Find(&books[i].Author, book.AuthorID);
		app.Db.Find(&books[i].Riview, book.RiviewID);
		app.Db.Find(&books[i].Publisher, book.PublisherID);
	}
	return c.RenderJson(books)
}

func (c BooksController) Find() revel.Result {
	var book models.Books
	var id int64
	c.Params.Bind(&id, "id")
	if err := app.Db.First(&book, id).Error; err != nil {
		return c.RenderJson(util.ResponseError("book not founded"))
	}
	app.Db.First(&book.Author, book.AuthorID)
	return c.RenderJson(util.ResponseSuccess(book))
}
