package readIPs

import (
	"os"
	"bufio"
	"fmt"
)

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
			lines = append(lines, scanner.Text())
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
