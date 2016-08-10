package app

import (
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"github.com/newisfa/test/app/models"
)
var Db *gorm.DB
func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		HeaderFilter,                  // Add some security based headers
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)

	revel.OnAppStart(func() {
		var err error
		Db, err = gorm.Open("sqlite3", "./database/books.db")

		if err != nil{
			log.Fatal(err)
		}
		log.Print(Db.DB().Stats());

		Db.AutoMigrate(&models.User{});
		Db.AutoMigrate(&models.Author{});
		Db.AutoMigrate(&models.Publisher{});
		Db.AutoMigrate(&models.Riview{});
		Db.AutoMigrate(&models.Books{});

		Db.DB().SetMaxIdleConns(10)
		Db.DB().SetMaxOpenConns(100)

		Db.SingularTable(false)
		Db.LogMode(true)
	})
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	//c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	//c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	//c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Set("Access-Control-Allow-Origin","*")
	c.Response.Out.Header().Set("Access-Control-Allow-Methods", "GET, POST,OPTIONS, PUT, DELETE")
	c.Response.Out.Header().Set("Access-Control-Allow-Headers", "Accept, Cache-Control, Content-Type,Content-Length,token,Accept-Encoding, X-CSRF-Token,Authorization, x-requested-with")

	if c.Request.Method == "OPTIONS"{
		return
	}
	fc[0](c, fc[1:]) // Execute the next filter stage.
}
