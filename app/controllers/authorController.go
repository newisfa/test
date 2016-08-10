package controllers

import (
	"github.com/revel/revel"
	"github.com/newisfa/test/app/util"
	"log"
	"github.com/newisfa/test/app/encoders"
	"github.com/newisfa/test/app"
	"github.com/newisfa/test/app/models"
)

type AuthorController struct {
	*revel.Controller
}

func (c AuthorController) Create() revel.Result{
	var author = encoders.EncodeAuthor(c.Request.Body)

	if author.Name == "" && author.Address == "" {
		return c.RenderJson(util.ResponseError("Author information not founded"))
	}

	if err := app.Db.Create(&author).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Author Creation Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(author))
}

func (c AuthorController) Update() revel.Result  {
	var update = encoders.EncodeAuthor(c.Request.Body)
	var author models.Author
	var id int
	c.Params.Bind(&id, "id")
	if rowsCout := app.Db.First(&author, id).RowsAffected; rowsCout < 1 {
		return c.RenderJson(util.ResponseError("Author Update information Fialed"))
	}

	if err := app.Db.Model(&author).Updates(&update).Error; err != nil{
		return c.RenderJson(util.ResponseError("Author updates Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(update))
}

func (c AuthorController) Delete() revel.Result {
	var (
		id int
		author models.Author
	)

	c.Params.Bind(&id, "id")
	if rowsCount := app.Db.First(&author, id).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("Author information Not Founded"))
	}

	if err := app.Db.Delete(&author).Error; err != nil {
		return c.RenderJson(util.ResponseError("Author Delete Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(author))
}

func (c AuthorController) Get()revel.Result {
	var authors []models.Author
	var limitQuery = c.Request.URL.Query().Get("limt");
	if limitQuery == "" {
		limitQuery = "0"
	}
	var offsetQuery = c.Request.URL.Query().Get("offset");

	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&authors).RowsAffected; founded < 1 {
		c.RenderJson(util.ResponseError("Not founded authors"))
	}
	return c.RenderJson(authors)
}

func (c AuthorController) Find() revel.Result {
	var author models.Author
	var id int64
	c.Params.Bind(&id, "id")
	if err := app.Db.First(&author, id).Error; err != nil {
		return c.RenderJson(util.ResponseError("Author not founded"))
	}
	return c.RenderJson(util.ResponseSuccess(author))
}

