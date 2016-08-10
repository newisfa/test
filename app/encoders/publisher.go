package encoders

import (
	"io"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/newisfa/test/app/models"
)

func EncodePublisher(body io.ReadCloser) (publish models.Publisher)  {
	var data,_ = ioutil.ReadAll(body)

	if err := json.Unmarshal(data, &publish); err != nil {
		log.Println(err);
		return
	}
	return
}
