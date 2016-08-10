package encoders

import (
	"io"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/newisfa/test/app/models"
)

func EncodeToken(body io.ReadCloser) (token models.Token)  {
	var data,_ = ioutil.ReadAll(body)
	if err := json.Unmarshal(data, &token); err != nil {
		log.Println(err);
		return
	}
	return
}
