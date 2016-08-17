package controllers

import (
	"github.com/revel/revel"
	"log"
	"github.com/newisfa/test/app/util"
	"github.com/newisfa/test/app"
	"github.com/newisfa/test/app/encoders"
	"github.com/dgrijalva/jwt-go"
	"github.com/newisfa/test/app/models"
	"time"
)

type UsersController struct {
	*revel.Controller
}

func (c UsersController) Create() revel.Result {
	var user = encoders.EncodeSingleUsers(c.Request.Body)

	if user.Email == "" || user.Password == "" {
		return c.RenderJson(util.ResponseError("User information is empty"))
	}

	if err := app.Db.Create(&user).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("User Creation Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(user))
}

func (c UsersController) Login() revel.Result {
	var user = encoders.EncodeSingleUsers(c.Request.Body)

	if user.Email == "" || user.Password == "" {
		return c.RenderJson(util.ResponseError("please insert correct email and password"))
	}

	if founded := app.Db.Where(&user).First(&user).RowsAffected; founded < 1 {
		return c.RenderJson(util.ResponseError("User Not Founded"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":user.ID,
		"email":user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	appSecret,_ := revel.Config.String("app.secret")
	tokenString, err := token.SignedString([]byte(appSecret))

	if err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("Token Generation Failed"))
	}

	var tokenModel  models.Token
	tokenModel.Email = user.Email
	tokenModel.Name = user.Name
	tokenModel.Token = tokenString

	return c.RenderJson(util.ResponseSuccess(tokenModel))
}

func (c UsersController) Get()revel.Result {
	var users []models.User
	var limitQuery = c.Request.URL.Query().Get("limt");
	if limitQuery == "" {
		limitQuery = "0"
	}
	var offsetQuery = c.Request.URL.Query().Get("offset");

	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&users).RowsAffected; founded < 1 {
		c.RenderJson(util.ResponseError("Not founded users"))
	}
	return c.RenderJson(users)
}

func (c UsersController) Update() revel.Result  {
	var update = encoders.EncodeSingleUsers(c.Request.Body)
	var user models.User
	var id int
	c.Params.Bind(&id, "id")
	if rowsCout := app.Db.First(&user, id).RowsAffected; rowsCout < 1 {
		return c.RenderJson(util.ResponseError("user Update information Fialed"))
	}

	if err := app.Db.Model(&user).Updates(&update).Error; err != nil{
		return c.RenderJson(util.ResponseError("user updates Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(update))
}

func (c UsersController) Delete() revel.Result {
	var (
		id int
		user models.User
	)

	c.Params.Bind(&id, "id")
	if rowsCount := app.Db.First(&user, id).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("Post information Not Founded"))
	}

	if err := app.Db.Delete(&user).Error; err != nil {
		return c.RenderJson(util.ResponseError("Post Delete Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(user))
}

func (c UsersController) Find() revel.Result {
	var user models.User
	var id int64
	c.Params.Bind(&id, "id")
	if err := app.Db.First(&user, id).Error; err != nil {
		return c.RenderJson(util.ResponseError("user not founded"))
	}
	return c.RenderJson(util.ResponseSuccess(user))
}

