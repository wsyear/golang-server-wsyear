package main;
import(
	"net/http";
	"fmt";
	"./set";
)
func HandleServe(w http.ResponseWriter,r *http.Request){
    set.Setpage(w,"./"+r.URL.Path,nil);
}
func main(){
	http.HandleFunc("/", HandleServe);
	err:=http.ListenAndServe(":8080",nil);
	if(err!=nil){
		fmt.Println("服务器开启错误:\n",err);
	}
}