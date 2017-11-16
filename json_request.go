package bidder

import (
	"fmt"
	"strings"
	"github.com/gin-gonic/gin/json"
)

type StatisticsForIfa struct {
	Device struct {
		Ifa string `json:"ifa"`
	} `json:"device"`
}


type Statistics struct {
	App struct {
		App string `json:"bundle"`
	} `json:"app"`

	Device struct {
		Geo struct {
			Country string `json:"country"`
		} `json:"geo"`
		Platform string `json:"os"`
	} `json:"device"`
}

type ReturnStatistics struct {
	Country string `redis:"Country"`
	App string `redis:"App"`
	Platform string `redis:"Platform"`
	Count int
}


func (st *Statistics) GetStBytes() ([]byte, error) {
	b, e :=json.Marshal(st)
	if e != nil{
		return nil, e
	}
	return b, nil

}


func (st *Statistics) GetMarshalStruct() []byte {
	b, _ := json.Marshal(&ReturnStatistics{
		Country:st.Device.Geo.Country,
		App:st.App.App,
		Platform:st.Device.Platform,
	})
	return b
}

func (st *Statistics) GetFormattedJson() *ReturnStatistics {
	return &ReturnStatistics{
		Country:st.Device.Geo.Country,
		App:st.App.App,
		Platform:st.Device.Platform,
	}
}



func (st *Statistics) GetFormattedStatString() string {
	return strings.ToLower(fmt.Sprintf("%s:%s:%s", st.Device.Geo.Country, st.App.App, st.Device.Platform))
}

