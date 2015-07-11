package IO
import (
    "testing"
    "fmt"
)
func TestNewDate(t *testing.T){
    var str string="c:/setup.log";
    ReadLine(str,output)

//    if t1:=NewDate(2015,5,12);t1.Year()!=2015||t1.Month()!=5||t1.Day()!=12{
//        t.Error("新建的时间不正确.")
//    }else{
//        t.Log("NewDate()通过")
//    }
}
func output(line string){
    fmt.Println(line)
}

