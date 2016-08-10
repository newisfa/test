package encoders

import (
	"io"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/newisfa/test/app/models"
)

func EncodeBooks(body io.ReadCloser) (book models.Books)  {
	var data,_ = ioutil.ReadAll(body)
	if err := json.Unmarshal(data, &book); err != nil {
		log.Println("##########", book)
		log.Println(err);
		return
	}
	return
}
