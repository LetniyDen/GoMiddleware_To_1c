package main

import (
	connector "GoMiddleware_To_1c/connector"
	handlers "GoMiddleware_To_1c/handler"
	rootsctuct "GoMiddleware_To_1c/rootdescription"
)

func main() {

	rootsctuct.LoggerConnV.InitLog()
	connector.ConnectorV.LoggerConn = rootsctuct.LoggerConnV

	rootsctuct.Global_settingsV.LoadSettingsFromDisk()
	err := connector.ConnectorV.SetSettings(rootsctuct.Global_settingsV)

	if err != nil {
		connector.ConnectorV.LoggerConn.ErrorLogger.Println(err.Error())
	}

	handlers.StratHandlers()
}
