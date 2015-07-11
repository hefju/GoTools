package IO
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