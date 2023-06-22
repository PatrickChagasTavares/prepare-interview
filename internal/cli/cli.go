package cli

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/PatrickChagastavares/prepare-interview/internal/entity"
	"github.com/PatrickChagastavares/prepare-interview/pkg/chatgpt"
)

func Start() {
	cmds := entity.LoadComands()
	log := log.Default()

	switch cmds.Operation() {
	case entity.OperationVersion:
		fmt.Println(version)
		break

	case entity.OperationCompany:
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

	case entity.OperationPrepare:
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
