package handler

import (
	connector "GoMiddleware_To_1c/connector"
	rootsctuct "GoMiddleware_To_1c/rootdescription"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/beevik/etree"
	"github.com/gorilla/mux"
)

/*
	 func api_json(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":

			resp, err := connector.ConnectorV.GET()

			if err != nil {
				connector.ConnectorV.LoggerConn.ErrorLogger.Println(err.Error())
				fmt.Fprintf(w, err.Error())
				return
			}

			resp, err = connector.ConnectorV.POST(w)
			if err != nil {
				connector.ConnectorV.LoggerConn.ErrorLogger.Println(err.Error())
				fmt.Fprintf(w, err.Error())
				return
			}
			// fmt.Fprintf(w, string(JsonString))

		case "POST":

			// нам отправили запрос
			// перенаправляем запрос в 1с базу
			// отправляем ответ назад, что все нормик или нет
			resp, err := connector.ConnectorV.POSTto(w, sender)

			if err != nil {
				connector.ConnectorV.LoggerConn.ErrorLogger.Println(err.Error())
				fmt.Fprintf(w, err.Error())
				return
			}

		default:

			fmt.Fprintf(w, r.Method+" - This method is not implemented")

		}

}
*/
func Settings(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/settings.html", "templates/header.html")
	if err != nil {
		connector.ConnectorV.LoggerConn.ErrorLogger.Println(err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}

	tmpl.ExecuteTemplate(w, "settings", connector.ConnectorV.Global_settings)

}

func Api_xml(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		customer_map_s, err := connector.ConnectorV.GetAllCustomer()

		if err != nil {
			connector.ConnectorV.LoggerConn.ErrorLogger.Println(err.Error())
			fmt.Fprintf(w, err.Error())
			return
		}

		doc := etree.NewDocument()

		Custromers := doc.CreateElement("Custromers")

		for _, p := range customer_map_s {
			Custromer := Custromers.CreateElement("Custromer")
			Custromer.CreateAttr("value", p.Customer_id)

			id := Custromer.CreateElement("Customer_id")
			id.CreateAttr("value", p.Customer_id)
			name := Custromer.CreateElement("Customer_name")
			name.CreateAttr("value", p.Customer_name)
			type1 := Custromer.CreateElement("Customer_type")
			type1.CreateAttr("value", p.Customer_type)
			email := Custromer.CreateElement("Customer_email")
			email.CreateAttr("value", p.Customer_email)
		}

		doc.Indent(2)
		XMLString, _ := doc.WriteToString()

		fmt.Fprintf(w, XMLString)

	case "POST":

		body, err := io.ReadAll(r.Body)
		if err != nil {
			connector.ConnectorV.LoggerConn.ErrorLogger.Println(err.Error())
			fmt.Fprintf(w, err.Error())
		}

		response := ""
		connector.ConnectorV.PostTo(body, response)

		//fmt.Fprintf(w, resp)

		/* if err != nil {
			connector.ConnectorV.LoggerConn.ErrorLogger.Println(err.Error())
			fmt.Fprintf(w, err.Error())
		}
		*/
		doc := etree.NewDocument()
		if err := doc.ReadFromBytes(body); err != nil {
			panic(err)
		}

		var customer_map_xml = make(map[string]rootsctuct.Customer_struct)

		Custromers := doc.SelectElement("Custromers")

		for _, Custromer := range Custromers.SelectElements("Custromer") {

			Customer_struct := rootsctuct.Customer_struct{}

			if Customer_id := Custromer.SelectElement("Customer_id"); Customer_id != nil {
				value := Customer_id.SelectAttrValue("value", "unknown")
				Customer_struct.Customer_id = value
			}
			if Customer_name := Custromer.SelectElement("Customer_name"); Customer_name != nil {
				value := Customer_name.SelectAttrValue("value", "unknown")
				Customer_struct.Customer_name = value
			}
			if Customer_type := Custromer.SelectElement("Customer_type"); Customer_type != nil {
				value := Customer_type.SelectAttrValue("value", "unknown")
				Customer_struct.Customer_type = value
			}

			if Customer_email := Custromer.SelectElement("Customer_email"); Customer_email != nil {
				value := Customer_email.SelectAttrValue("value", "unknown")
				Customer_struct.Customer_email = value
			}

			customer_map_xml[Customer_struct.Customer_id] = Customer_struct
		}

		/* 	for _, p := range customer_map_xml {
			err := connector.ConnectorV.AddChangeOneRow(connector.ConnectorV.DataBaseType, p, rootsctuct.Global_settingsV)
			if err != nil {
				connector.ConnectorV.LoggerConn.ErrorLogger.Println(err.Error())
				fmt.Println(err.Error())
			}
		} */

		fmt.Fprintf(w, string(body))

	default:

		fmt.Fprintf(w, r.Method+" - This method is not implemented")

	}
}

func StratHandlers() {

	router := mux.NewRouter()

	router.HandleFunc("/", Settings)
	router.HandleFunc("/settings", Settings)

	//router.HandleFunc("/api_json", Api_json)
	router.HandleFunc("/api_xml", Api_xml)

	//router.HandleFunc("/list_customer", List_customer)

	http.Handle("/", router)
	fmt.Println("Server is listening...")

	http.ListenAndServe(":8181", nil)
}
