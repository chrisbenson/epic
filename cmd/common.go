package cmd

var adminUser string
var adminPassword string
var server string
var database string
//var epicUser string
var epicPassword string
var port string
var host string


func resetCreds() {
	adminUser = ""
	adminPassword = ""
	server = ""
	database = ""
	//epicUser = ""
	epicPassword = ""
	port = ""
	host = ""
}

var appName string
var appCode string