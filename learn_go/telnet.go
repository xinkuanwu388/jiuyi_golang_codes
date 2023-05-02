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
	//s := []string{"5244", "5222"}
	tcpGather(readFile())

}
func tcpGather(ips []string) {

	//results := make(map[string]string)
	for _, address := range ips {
		//address := net.JoinHostPort(ip, port)
		//fmt.Println(address)
		// 3 秒超时
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			//results[ip] = "failed"
			fmt.Println("                          Fa---" + address)
			// todo log handler
		} else {
			if conn != nil {
				//results[ip] = "success"
				fmt.Println("ok---" + address)
				_ = conn.Close()
			} else {
				//results[ip] = "failed"
				fmt.Println("                          Fa---" + address)
			}
		}
	}
	//return results
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

		//lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines
	//fmt.Println("read lines:")
	//for _, line := range lines {
	//	fmt.Println(line)
	//}
}
