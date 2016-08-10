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
