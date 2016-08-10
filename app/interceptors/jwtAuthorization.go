package interceptors

import (
	"github.com/revel/revel"
	"github.com/newisfa/test/app/util"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"strconv"
	"time"
	"log"
)

type JWTAuthorization struct {
	*revel.Controller
}

func (c JWTAuthorization) CheckUser() revel.Result  {
	var tokenString = c.Request.Header.Get("token")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{},error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("Unexpected Signing Method: %v",token.Header["alg"])
		}
		appSecret,_ := revel.Config.String("app.secret")
		return []byte(appSecret), nil
	})

	if err == nil {
		if claims,ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Session["email"] = claims["email"].(string)
			//c.Session["id"] = strconv.FormatFloat(claims["id"].(float64), 'f', 6, 64)
			c.Session["id"] = strconv.Itoa(int(claims["id"].(float64)))
			var expDate = time.Unix(int64(claims["exp"].(float64)),0)
			if expDate.Before(time.Now()){
				return c.RenderJson(util.ResponseError("Expired Token"))
			}
			return nil
		}
		return c.RenderJson(util.ResponseError("Invalid Token Key"))
	}else {
		log.Println("&&&&&&&&&&&&&&",err);
		return c.RenderJson(util.ResponseError("Not Founded Token Key"))
	}
}

func init() {
	revel.InterceptMethod(JWTAuthorization.CheckUser, revel.BEFORE)
}

