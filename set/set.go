package set;
import(
	"net/http";
    "text/template";
)
func Setpage(op bool,w http.ResponseWriter,file string,data map[string]string)(error){
	if(op){
		t,err:=template.ParseFiles(file);
		if(err!=nil){
			return err;
		}
		t.Execute(w,data);
		return nil;
	}else{
		t:=template.New("page");
		t,err:=t.Parse(file);
		if(err!=nil){
			return err;
		}
		t.Execute(w,data);
		return nil;
	}
}
