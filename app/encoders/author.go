package encoders

import (
	"io"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/newisfa/test/app/models"
)

func EncodeAuthor(body io.ReadCloser) (author models.Author)  {
	var data,_ = ioutil.ReadAll(body)
	if err := json.Unmarshal(data, &author); err != nil {
		//log.Println("##########", author)
		log.Println(err);
		return
	}
	return
}
