package DateTime
import (
    "testing"
)

func TestParse(t *testing.T){

    if dt,e:=Parse("2015-01-01");e!=nil||dt.Year()!=2015||dt.Month()!=1||dt.Day()!=1{
        t.Error("时间转换出错.")
    }else{
        t.Log("第一个测试通过")
    }
}

func TestAdd(t *testing.T) {
    if add(1, 2) != 3 {
        t.Error("test foo:Addr failed")
    } else {
        t.Log("test foo:Addr pass")
    }
}