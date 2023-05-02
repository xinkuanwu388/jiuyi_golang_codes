package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	tcpGather(readFile())

}
func tcpGather(ips []string) {

	for _, address := range ips {
		// 3 秒超时
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			fmt.Println("                          Fa---" + address)
			// todo log handler
		} else {
			if conn != nil {
				fmt.Println("ok---" + address)
				_ = conn.Close()
			} else {
				fmt.Println("                          Fa---" + address)
			}
		}
	}
}

func readFile() []string {

	file, err := os.Open("xx.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// 这是我们的缓冲区
	var lines []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			address := strings.TrimSpace(line[9:29]) + ":" + strings.TrimSpace(line[29:])
			lines = append(lines, address)
		}
	}
	return lines
}
