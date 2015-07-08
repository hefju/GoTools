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

func add(x, y int) int {
    return x + y
}