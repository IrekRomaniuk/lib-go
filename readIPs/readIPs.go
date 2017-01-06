package readIPs

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"regexp"
)
/*
Input: file with list of single ip addresses per line
Output: slice with valid ip addresses
 */
func ReadIPs(target string) ([]string, error) {
	if !pathExists(target) {
		return nil, fmt.Errorf("File %s does not exist", target)
	}
	return readHosts(target)
}

func readHosts(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			// omit line if empty
			if validIP4(line) {
				lines = append(lines, scanner.Text())
			}

		}
	}
	return lines, scanner.Err()
}

func pathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil { return true }
	if os.IsNotExist(err) { return false }
	return true
}
//https://www.socketloop.com/tutorials/golang-validate-ip-address
func validIP4(ip string) bool {
	ip = strings.Trim(ip, " ")

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ip) {
		return true
	}
	return false
}
