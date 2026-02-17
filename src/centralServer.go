package main

import (
	"fmt"
	"log"
	"net/http"
)
// serverConfigs struct{
// 	serverType string
// 	port int
// 	mode string
// 	dbUserName string
// 	dbPass string
// 	logFilePath string
// 	notifEmail string
// }



//User to Server Requests
//Accessing the index page
func indexPage(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path!="/"){
		http.ServeFile(w,r,"./html/404.html")
	}
	http.ServeFile(w,r,"./html/index.html")
}

//Accessing the login page
func loginPage(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"./html/login.html")
}

//accessing the admin page
func adminPage(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"./html/admin.html")
}
//accessing the analyst page
func analystPage(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"./html/analyst.html")
}

//logging into the analyst/admin page


//requesting data from the server



//Relay Server requests
//inventory look up
//post end of day reports
//post stock take
func main() {
	fs:=http.FileServer(http.Dir("./html/static"))
	http.Handle("/static/",http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/login",loginPage)
	http.HandleFunc("/admin",adminPage)
	http.HandleFunc("/analyst",analystPage)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}



