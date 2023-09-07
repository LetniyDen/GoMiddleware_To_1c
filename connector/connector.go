package connector

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"

	rootsctuct "GoMiddleware_To_1c/rootdescription"
)

// ConnectorV центральная переменная (своего рода движок)
var ConnectorV Connector

// Connector центральная структура (своего рода движок)
type Connector struct {
	Global_settings rootsctuct.Global_settings
	LoggerConn      rootsctuct.LoggerConn
	DemoDBmap       map[string]rootsctuct.Customer_struct
	Mutex           *sync.Mutex
}

func (Connector *Connector) GetAllCustomer() (map[string]rootsctuct.Customer_struct, error) {

	resp, err := http.Get("http://localhost/REST_test/hs/exchange/custom_json")
	if err != nil {
		Connector.LoggerConn.ErrorLogger.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Connector.LoggerConn.ErrorLogger.Println(err.Error())
		return nil, err
	}

	var customer_map_json = make(map[string]rootsctuct.Customer_struct)

	err = json.Unmarshal(body, &customer_map_json)
	if err != nil {
		Connector.LoggerConn.ErrorLogger.Println(err.Error())
		return nil, err
	}

	return customer_map_json, nil

}

func (Connector *Connector) SetSettings(Global_settings rootsctuct.Global_settings) error {

	Connector.Global_settings = Global_settings

	return nil

}

func (Connector *Connector) PostTo(content []byte, response string) {

	resp, err := http.Post("http://localhost/Test_golang_1/hs/exchange/custom_json", "application/json", bytes.NewBuffer(content))
	if err != nil {
		Connector.LoggerConn.ErrorLogger.Println(err.Error())
	}

	defer resp.Body.Close()
	body_response, _ := io.ReadAll(resp.Body)
	response = string(body_response)
}
