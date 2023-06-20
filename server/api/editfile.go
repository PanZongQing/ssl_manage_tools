package api

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReplaceSslFile(filenname string) {
	filepath := "./confpath/" + filenname
	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ssl_certificate") {
			newLine := strings.Replace(line, "/etc/nginx/conf.d/www.polixir.site.pem", "/etc/nginx/conf.d/www.polixir.site.pem", 1)
			_, err := file.Seek(-int64(len(line)), os.SEEK_CUR)
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = file.WriteString(newLine)
			if err != nil {
				fmt.Println(err)
				return
			}
			break

		}
	}
	file.Sync()

}
