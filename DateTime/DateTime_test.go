package DateTime
import (
    "testing"
   // "time"
)

func TestParse(t *testing.T){

    if dt,e:=Parse("2015-01-07");e!=nil||dt.Year()!=2015||dt.Month()!=1||dt.Day()!=7{
        t.Error("时间转换出错.")
    }else{
        t.Log("第一个测试通过")
    }
}
func TestNewDate(t *testing.T){
    if t1:=NewDate(2015,5,12);t1.Year()!=2015||t1.Month()!=5||t1.Day()!=12{
        t.Error("新建的时间不正确.")
    }else{
        t.Log("NewDate()通过")
    }
}

func TestAdd(t *testing.T) {
    if add(1, 2) != 3 {
        t.Error("test foo:Addr failed")
    } else {
        t.Log("test foo:Addr pass")
    }
}