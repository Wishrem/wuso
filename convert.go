package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/Wishrem/wuso/pkg/protocol"
)

func main() {
	to := flag.Int64("to", -1, "-to 123456")
	from := flag.Int64("from", -1, "-from 123456")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("input text:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		b, err := protocol.BuildMsgBytes(-1, *from, *to, input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("result:", base64.StdEncoding.EncodeToString(b))
	}
}
