package set;
import(
	"net/http";
	"fmt";
    "io/ioutil";
    "text/template";
)
func Setpage(w http.ResponseWriter,file string,data map[string]string){
	cont,err:=ioutil.ReadFile(file);
	temp:=template.New("");
	if(err!=nil){
		fmt.Println("Write err:",err);
    	c404,_:=ioutil.ReadFile("./html/404.html");
	    c404str:=string(c404);
	    temp.Parse(c404str);
	    temp.Execute(w,nil);
        return;
	}
	contstr:=string(cont);
	temp.Parse(contstr);
	temp.Execute(w,data);
}
