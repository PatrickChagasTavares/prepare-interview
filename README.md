# Prepare you for the interview

## What is the idea of the repository?

Help you prepare for the interview, facilitate searching who the company is, and write about what you need to know or study about the vacancy.
I used OpenAI (ChatGPT) to generate that summary so you need to generate the api_key to use the commands.

[Here](https://elephas.app/blog/how-to-create-openai-api-keys-cl5c4f21d281431po7k8fgyol0) is a post to help to generate the API key.

## How to Use:

### search resume company
```
go run cmd/interview/main.go -token="your token openID" company -name="Facebook"
```
![prepare_interview_company](https://github.com/PatrickChagastavares/prepare-interview/assets/49497853/314b8918-1d8c-4020-9681-e5cc21be3af6)
### Generate a summary to prepare for the vacancy
```
go run cmd/interview/main.go -token="your token openID" prepare -role="Software Developer" -skills="Golang, Node e GCP" -lang="english"
```
![prepare_interview_company](https://github.com/PatrickChagastavares/prepare-interview/assets/49497853/9de17435-48f4-461d-9317-d58f73d63e2f)
## Todo list:
- [X] Create Readme;
- [X] Implement flags;
- [X] Implement HTTP requests to OpenAI;
- [ ] Tests with mocks;
- [ ] Complete all tests;
