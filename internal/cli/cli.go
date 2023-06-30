package cli

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/PatrickChagastavares/prepare-interview/internal/config"
	"github.com/PatrickChagastavares/prepare-interview/pkg/chatgpt"
)

func Start() {
	cmds := config.LoadComands()
	log := log.Default()

	switch cmds.Operation() {
	case config.OperationVersion:
		fmt.Println(version)
		break

	case config.OperationCompany:
		if !cmds.HasToken() {
			log.Fatalln("ERROR: token is required")
			break
		}
		chatGPT := chatgpt.NewHttpChatGPT(cmds.Token)

		input := chatGPT.MakeInputCompany(cmds.Name, cmds.Lang)

		output, err := chatGPT.CreateCompletion(context.TODO(), input)
		if err != nil {
			log.Fatalf("ERROR: problem request openAI -> %s", err.Error())
			break
		}

		log.Println(output)

	case config.OperationPrepare:
		if !cmds.HasToken() {
			log.Fatalln("ERROR: token is required")
			break
		}

		chatGPT := chatgpt.NewHttpChatGPT(cmds.Token)

		input := chatGPT.MakeInputInterview(cmds.Role, cmds.Skills, cmds.Lang)

		output, err := chatGPT.CreateCompletion(context.TODO(), input)
		if err != nil {
			log.Fatalf("ERROR: problem request openAI -> %s", err.Error())
			break
		}

		log.Println(output)

	default:
		flag.Usage()
		break

	}
}
