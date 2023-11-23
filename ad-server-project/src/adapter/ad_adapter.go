package adapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)


type PCTRResponse struct {
	PCTR []float64 `json:"pctr"`
}

// SendCtrPredictionServer ctr-prediction-server 에 요청하여 광고의 CTR 조회
func SendCtrPredictionServer(userId int, idList []int) (*PCTRResponse, error) {
	idListLength := len(idList)
	adIdsString := ""
	if idListLength > 0{
		adIdsString = strconv.Itoa(idList[0])
		for i := 1; i < idListLength; i++ {
			adIdsString += "," + strconv.Itoa(idList[i])
		}
	}

	url := fmt.Sprintf("https://predict-ctr-pmj4td4sjq-du.a.run.app/?user_id=%d&ad_campaign_ids=%s", userId, adIdsString)

	response, err := http.Get(url)
	if err != nil {
		log.Printf("[Error] get CTR from ctr-prediction-server: %v", err)
		return nil, err
	}
	defer response.Body.Close()

	// 응답 결과를 읽음
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var pctrResponse PCTRResponse
	if err := json.Unmarshal(body, &pctrResponse); err != nil {
		return nil, err
	}

	return &pctrResponse, nil
}