package rootdescription

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Global_settingsV Global_settings
var LoggerConnV LoggerConn

// структура настроек сети
type Global_settings struct {
	DB_centr  string
	DB_mobile string
}

type Customer_struct struct {
	Customer_id    string
	Customer_name  string
	Customer_type  string
	Customer_email string
	Address_Struct Address_Struct
}

type Address_Struct struct {
	Street string
	House  int
}

func (GlobalSettings *Global_settings) SaveSettingsOnDisk() {

	f, err := os.Create("./settings/config.json")
	if err != nil {
		log.Fatal(err)
	}

	JsonString, err := json.Marshal(GlobalSettings)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write(JsonString); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (GlobalSettings *Global_settings) LoadSettingsFromDisk() {

	file, err := os.OpenFile("./settings/config.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		file, err = os.CreateTemp("", "ConfigJsonlogConnector1C")
		if err != nil {
			fmt.Println(err)
		}
	}

	decoder := json.NewDecoder(file)
	Settings := Global_settings{}
	err = decoder.Decode(&Settings)
	if err != nil {
		fmt.Println(err)
	}

	//Global_settingsV = Settings
	*GlobalSettings = Settings

	if err := file.Close(); err != nil {
		fmt.Println(err)
	}
}

type LoggerConn struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func (LoggerConn *LoggerConn) InitLog() {

	file, err := os.OpenFile("./logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		file, err = os.CreateTemp("", "logCoonector1C")
		if err != nil {
			log.Fatal(err)
		}
	}

	LoggerConn.InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerConn.ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	LoggerConn.ErrorLogger.Println("Starting the application...")
}
