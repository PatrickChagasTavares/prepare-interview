package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type (
	httpChatgpt struct {
		baseURL string
		token   string
	}

	reqBody struct {
		Model       string `json:"model"`
		Prompt      string `json:"prompt"`
		MaxToken    uint   `json:"max_tokens"`
		Temperature uint   `json:"temperature"`
	}

	response struct {
		Choices []struct {
			Text string `json:"text"`
		}
	}
	errResponse struct {
		ErrorHttp struct {
			Message string `json:"message"`
		} `json:"error"`
	}
)

func (e *errResponse) Error() string {
	return e.ErrorHttp.Message
}

const gpt3TextDavinci003 = "text-davinci-003"

func NewHttpChatGPT(token string) Ichatgpt {
	return &httpChatgpt{
		baseURL: "https://api.openai.com/v1",
		token:   token,
	}
}

func (hc *httpChatgpt) CreateCompletion(ctx context.Context, input string) (output string, err error) {
	requestBody, err := json.Marshal(reqBody{
		Model:  gpt3TextDavinci003,
		Prompt: input,
		// https://platform.openai.com/docs/api-reference/completions/create#completions/create-max_tokens
		MaxToken: 3900,
		// https://platform.openai.com/docs/api-reference/completions/create#completions/create-temperature
		Temperature: 1,
	})
	if err != nil {
		return output, err
	}

	url := hc.baseURL + "/completions"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(requestBody))
	if err != nil {
		return output, err
	}
	req.Close = true

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+hc.token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return output, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode < 200 || res.StatusCode > 299 {
		var errHttp errResponse
		if err := decoder.Decode(&errHttp); err != nil {
			return output, err
		}

		return output, &errHttp
	}

	var data response
	err = decoder.Decode(&data)
	if err != nil {
		return output, err
	}

	if len(data.Choices) == 0 {
		return output, errors.New("gpt not return text")
	}

	return data.Choices[0].Text, nil
}

func (hc httpChatgpt) MakeInputInterview(role, skills, lang string) (output string) {

	switch lang {
	case "english":
		output = fmt.Sprintf(
			`I would like you to help me prepare for an interview for the position of "%s" in "%s".
			Finally generate this template linearly.`,
			role, skills)

	default:
		output = fmt.Sprintf(
			`Gostaria que você me ajudasse a me preparar para uma entrevista para o cargo de "%s" em "%s".
				Por fim gere esse template de forma linear.`,
			role, skills)
	}

	return
}
