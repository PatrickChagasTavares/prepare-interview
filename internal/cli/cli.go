package cli

import (
	"context"
	"flag"
	"log"

	"github.com/PatrickChagastavares/prepare-interview/internal/entity"
	"github.com/PatrickChagastavares/prepare-interview/pkg/chatgpt"
)

func Start() {
	cmds := entity.LoadComands()
	log := log.Default()

	switch cmds.Operation() {
	case entity.OperationHelper:
		flag.Usage()
		break
	case entity.OperationVersion:
		log.Println(version)
		break
	default:
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
	}
}
