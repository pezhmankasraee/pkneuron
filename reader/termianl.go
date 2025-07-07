package reader

import (
	"bufio"
	"os"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
)

func Read() string {

	reader := bufio.NewReader(os.Stdin)

	println("Enter the request> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		pklog.CreateLog(pklog.FatalError, err.Error())
	}

	input = strings.TrimSpace(input)
	return input
}
