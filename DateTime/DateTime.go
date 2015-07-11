//go的日期转换, 类似与c#的将字符串转换成日期. 输入年月日转换成日期
//go的日期转换比较麻烦, 例如我想将"2015-5-12" 那么就就要写一串"2006-1-02"的字符串来转换, 注意多一位少一位
//如果传入"2015.5.12"那么就要写"2006.1.02", 如果传入"2015.05.12" 那么就要"2006.01.02"
//还有用Parse是转换成0时区的,必须使用ParseInLocation才是当前时区的.
package DateTime
import "time"

var layout string = "2006-01-02 15:04:05"
func Parse(timestring string)(time.Time,error){
    switch len(timestring) {
        case 10:
        timestring=timestring+" 00:00:00"
        case 8:
        timestring=time.Now().Format("2006-01-02")+" "+ timestring
    }
   return time.ParseInLocation(layout, timestring, time.Local)
}
func NewDate(year , month , day int)time.Time{
    m:=time.Month(month)
    t := time.Date(year , m , day, 0, 0, 0, 0, time.Local)
    return t
}

func add(x, y int) int {
    return x + y
}