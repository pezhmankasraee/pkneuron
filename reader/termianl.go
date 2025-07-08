package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
)

func Read() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\033[33;1mNOTE: use c:s / c:send at the end of the request)\033[0m")
	fmt.Println("\033[33mrequest > \033[0m")

	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			pklog.CreateLog(pklog.FatalError, err.Error())
			break
		}

		line = strings.TrimRight(line, "\r\n")
		if strings.EqualFold(line, "c:e") || strings.EqualFold(line, "c:exit") {
			fmt.Println("Exiting ...")
			os.Exit(0)
		}

		if strings.EqualFold(line, "c:send") || strings.EqualFold(line, "c:s") {
			break
		}

		lines = append(lines, line)
	}
	input := strings.Join(lines, "\n")

	return input
}
