package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/AmazonScraper/ScraperAPI/models"
)

type productCaller struct {
	URL string
}

func NewCaller(URL string) Caller {
	return &productCaller{URL: URL}
}

func (pc *productCaller) Call(document *models.ProductDocument) error {
	body, err := json.Marshal(document)
	if err != nil {
		log.Println("Error in marshaling document: ", err.Error())
		return err
	}

	log.Println("BODY ", string(body))

	req, err := http.NewRequest(http.MethodPost, pc.URL, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error in creating Request: ", err.Error())
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error in calling service: ", err.Error())
		return err
	} else if resp.StatusCode != 201 {
		log.Println("Error in calling service")
		return errors.New("error in calling service")
	}

	return nil
}
