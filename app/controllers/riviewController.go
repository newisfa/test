package controllers

import (
	"github.com/revel/revel"
	"github.com/newisfa/test/app/util"
	"log"
	"github.com/newisfa/test/app/encoders"
	"github.com/newisfa/test/app"
	"github.com/newisfa/test/app/models"
)

type RiviewController struct {
	*revel.Controller
}

func (c RiviewController) Create() revel.Result{
	var riview = encoders.EncodeRiview(c.Request.Body)

	if riview.Name == "" && riview.Address == "" {
		return c.RenderJson(util.ResponseError("Riview information not founded"))
	}
	//riview.BooksID,_ = strconv.ParseInt(c.Session["id"], 10, 0)
	if err := app.Db.Create(&riview).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Riview Creation Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(riview))
}

func (c RiviewController) Update() revel.Result  {
	var update = encoders.EncodeRiview(c.Request.Body)
	var riview models.Riview
	var id int
	c.Params.Bind(&id, "id")
	if rowsCout := app.Db.First(&riview, id).RowsAffected; rowsCout < 1 {
		return c.RenderJson(util.ResponseError("Riview Update information Fialed"))
	}

	if err := app.Db.Model(&riview).Updates(&update).Error; err != nil{
		return c.RenderJson(util.ResponseError("Riview updates Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(update))
}

func (c RiviewController) Delete() revel.Result {
	var (
		id int
		riview models.Riview
	)

	c.Params.Bind(&id, "id")
	if rowsCount := app.Db.First(&riview, id).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("Riview information Not Founded"))
	}

	if err := app.Db.Delete(&riview).Error; err != nil {
		return c.RenderJson(util.ResponseError("Riview Delete Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(riview))
}

func (c RiviewController) Get()revel.Result {
	var riviews []models.Riview
	var limitQuery = c.Request.URL.Query().Get("limt");
	if limitQuery == "" {
		limitQuery = "0"
	}
	var offsetQuery = c.Request.URL.Query().Get("offset");

	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&riviews).RowsAffected; founded < 1 {
		c.RenderJson(util.ResponseError("Not founded Riviews"))
	}
	return c.RenderJson(riviews)
}

func (c RiviewController) Find() revel.Result {
	var riview models.Riview
	var id int64
	c.Params.Bind(&id, "id")
	if err := app.Db.First(&riview, id).Error; err != nil {
		return c.RenderJson(util.ResponseError("Riview not founded"))
	}

	return c.RenderJson(util.ResponseSuccess(riview))
}
