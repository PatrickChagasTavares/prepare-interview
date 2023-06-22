package chatgpt

import "context"

type (
	Ichatgpt interface {
		CreateCompletion(ctx context.Context, input string) (output string, err error)
		MakeInputInterview(role, skills, lang string) (output string)
		MakeInputCompany(name, lang string) (output string)
	}
)
