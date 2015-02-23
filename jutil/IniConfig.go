package jutil
//这个文件是用来从ini读取配置信息的
import (
	"log"
	"strings"
)

//用于测试时候调用, 正常使用是在main方法里面调用jutil.Initconfiguration()
func TestIni_configuration(){
	Initconfiguration()
	log.Println("测试:",IniConfiger["server"]["ip"])
}
//初始化配置文件
func Initconfiguration(){
	ParseIni("conf.ini")
}

//存储Ini配置的数据
var IniConfiger map[string]map[string]string

func ParseIni(filename string) {
	IniConfiger = make(map[string]map[string]string)
	ReadLine(filename, handle) //从指定文件中逐行读取字符串, 第二参数是每行字符的处理方法
}

var section string //上次读取的section

func handle(line string) {
	length := len(line)
	switch {
	case length == 0: //没有字符不处理
	case line[0] == ';': //注释不处理

	case line[0] == '[' && line[length-1] == ']': //处理section
		section = line[1 : length-1]
		IniConfiger[section] = make(map[string]string)

	default: //处理key-value
		i := strings.Index(line, "=")
		key := strings.TrimSpace(line[0:i])
		value := strings.TrimSpace(line[i+1 : length])
		if len(section) == 0 {
			log.Fatalln("IniConfig.go, section is empty.")
		}
		IniConfiger[section][key] = value
	}
}
