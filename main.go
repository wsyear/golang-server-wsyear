package main;
import(
	"net/http";
	"fmt";
	"./set";
)
const(
	page404=`
<head>
	<meta charset="utf-8">
	<title>404 Not Found</title>
</head>
<body style="display: flex;">
	<div style="margin: auto;">
		<div style="font-size: 10rem;text-align: center;font-weight: 1000;font-family: monaco;color:#3f3f3f">
			Oops!
		</div>
		<div style="font-size:2rem;text-align: center;font-family: monaco;">
			404 Not Found.<br>
			The web page seems to be lost
		</div>
	</div>
</body>
`;
)
func set404(w http.ResponseWriter){
    err:=set.Setpage(true,w,"./html/404.html",nil);
    if(err!=nil){
    	err=set.Setpage(false,w,page404,nil);
    }
}
func HandleServe(w http.ResponseWriter,r *http.Request){
	path:=r.URL.Path;
    err:=set.Setpage(true,w,"./html/"+path,nil);
    if(err!=nil){
    	err:=set.Setpage(true,w,"./html/"+path+"/index.html",nil);
    	if(err!=nil){
    		set404(w);
    	}
    }
}
func main(){
	http.HandleFunc("/", HandleServe);
	fmt.Println("运行于:127.0.0.1:8080");
	err:=http.ListenAndServe(":8080",nil);
	if(err!=nil){
		fmt.Println("服务器开启错误:\n",err);
	}
}
