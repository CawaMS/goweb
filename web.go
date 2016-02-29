package main

import(
"fmt"
"net/http"
"io/ioutil"
"os"
)

func handler (w http.ResponseWriter, r *http.Request){
	
	host_name, err := os.Hostname()
	if err != nil{
		fmt.Fprintf(w, "%s", err)
	}
	
	fmt.Fprintf(w, "my first update")
	fmt.Fprintf(w, "<h1>Web Client Host Name</h1> <div>%s</div>", string(host_name))
	
	res_mn, err := http.Get("http://restservice-demo-azure.marathon.mesos:5000/api/info/machinename")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}else{
		defer res_mn.Body.Close()
		contents_mn, err := ioutil.ReadAll(res_mn.Body)
		if err != nil{
			fmt.Fprintf(w, "%s", err)
		}
		fmt.Fprintf(w, "<h1>REST Machine Host Name </h1> <div>%s</div>", string(contents_mn))
	}

	res_ct, err := http.Get("http://restservice-demo-azure.marathon.mesos:5000/api/info/currenttime")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}else{
		defer res_ct.Body.Close()
		contents_ct, err := ioutil.ReadAll(res_ct.Body)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
		}
		fmt.Fprintf(w, "<h1>REST Access Time </h1> <div>%s</div>", string(contents_ct))
	}

	
	
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}