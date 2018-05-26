package ming

import (
	"net/http"
)



func HttpAction()  {
	
	 client := http.Client{

	 }	
	 res ,_ := http.NewRequest("GET","http://127.0.0.1:9080/backend",nil)
	 res.Header.Add("If-None-Match", `W/"wyzzy"`)
	 client.Do(res)

}