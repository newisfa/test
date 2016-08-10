package encoders

import (
	"io"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/newisfa/test/app/models"
)

func EncodeRiview(body io.ReadCloser) (riview models.Riview)  {
	var data,_ = ioutil.ReadAll(body)
	if err := json.Unmarshal(data, &riview); err != nil {
		log.Println("##########", riview)
		log.Println(err);
		return
	}
	return
}
