package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	olapi "github.com/ollama/ollama/api"
)

func main() {
	var (
		flagVerbose bool
		flagStdin   bool
		flagModel   string // model
	)

	flag.BoolVar(&flagVerbose, "v", false, "verbose")
	flag.BoolVar(&flagStdin, "i", false, "enable stdin")
	flag.StringVar(&flagModel, "m", os.Getenv("UAI_MODEL"), "model")
	flag.Parse()

	cli, err := olapi.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	prompt := strings.Join(flag.Args(), " ")
	if flagStdin {
		p, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		prompt += "\n---\n" + string(p) + "\n---\n"
	}

	if flagVerbose {
		fmt.Fprint(os.Stderr, "The prompt:\n", prompt)
		fmt.Fprint(os.Stderr, "\n\n-- response\n")
	}

	req := &olapi.GenerateRequest{
		Model:   flagModel,
		Prompt:  prompt,
		Context: []int{},
	}
	err = cli.Generate(context.Background(), req, func(res olapi.GenerateResponse) error {
		fmt.Print(res.Response)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
