package controllers

import (
	"github.com/revel/revel"
	"github.com/newisfa/test/app/util"
	"log"
	"github.com/newisfa/test/app/encoders"
	"github.com/newisfa/test/app"
	"github.com/newisfa/test/app/models"
)

type PublisherController struct {
	*revel.Controller
}

func (c PublisherController) Create() revel.Result{
	var publisher = encoders.EncodePublisher(c.Request.Body)

	if publisher.Name == "" && publisher.Address == "" {
		return c.RenderJson(util.ResponseError("publisher information not founded"))
	}
	//publisher.AuthorID,_ = strconv.ParseInt(c.Session["id"], 10, 0)
	if err := app.Db.Create(&publisher).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("publisher Creation Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(publisher))
}

func (c PublisherController) Update() revel.Result  {
	var update = encoders.EncodePublisher(c.Request.Body)
	var publisher models.Publisher
	var id int
	c.Params.Bind(&id, "id")
	if rowsCout := app.Db.First(&publisher, id).RowsAffected; rowsCout < 1 {
		return c.RenderJson(util.ResponseError("Publisher Update information Fialed"))
	}

	if err := app.Db.Model(&publisher).Updates(&update).Error; err != nil{
		return c.RenderJson(util.ResponseError("Publisher updates Fialed"))
	}
	log.Println("@@@@@@@@@@@@@@@@@@@#####",update)
	return c.RenderJson(util.ResponseSuccess(update))
}

func (c PublisherController) Get()revel.Result {
	var publishers []models.Publisher
	var limitQuery = c.Request.URL.Query().Get("limt");
	if limitQuery == "" {
		limitQuery = "0"
	}
	var offsetQuery = c.Request.URL.Query().Get("offset");

	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&publishers).RowsAffected; founded < 1 {
		c.RenderJson(util.ResponseError("Not founded Publisher"))
	}
	return c.RenderJson(publishers)
}

func (c PublisherController) Delete() revel.Result {
	var (
		id int
		publisher models.Publisher
	)

	c.Params.Bind(&id, "id")
	if rowsCount := app.Db.First(&publisher, id).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("Publisher information Not Founded"))
	}

	if err := app.Db.Delete(&publisher).Error; err != nil {
		return c.RenderJson(util.ResponseError("Publisher Delete Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(publisher))
}

func (c PublisherController) Find() revel.Result {
	var publisher models.Publisher
	var id int64
	c.Params.Bind(&id, "id")
	if err := app.Db.First(&publisher, id).Error; err != nil {
		return c.RenderJson(util.ResponseError("Publisher not founded"))
	}
	return c.RenderJson(util.ResponseSuccess(publisher))
}
