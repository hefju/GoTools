package jutil
//该文件是对go文件的扩展, go逐行读取数据比较麻烦, 在这里写好, 以后调用就可以了
import (
	"bufio"
	"io"
	"os"
	"strings"
)

//逐行读取数据, 默认utf8, windows建立的文件会显示乱码
func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
