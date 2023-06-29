# Prepare you for the interview

## What is the idea of the repository?

Help you prepare for the interview, facilitate searching who the company is, and write about what you need to know or study about the vacancy.
I used OpenAI (ChatGPT) to generate that summary so you need to generate the api_key to use the commands.

[Here](https://elephas.app/blog/how-to-create-openai-api-keys-cl5c4f21d281431po7k8fgyol0) is a post to help to generate the API key.

## How to Use:

### search resume company
```
go run cmd/interview/main.go -token="your token openID" company -name="Linkedin"
```
![prepare_interview_company](https://github.com/PatrickChagastavares/prepare-interview/assets/49497853/5d0900e7-ba98-4994-9dac-e6cf4641f3d2)
### Generate a summary to prepare for the vacancy
```
go run cmd/interview/main.go -token="your token openID" prepare -role="Software Developer" -skills="Golang, Node e GCP" -lang="portuguese"
```
![prepare_interview_role](https://github.com/PatrickChagastavares/prepare-interview/assets/49497853/e6b888e8-1755-4a5e-860b-780f9b4cd066)

## Todo list:
- [X] Create Readme;
- [X] Implement flags;
- [X] Implement HTTP requests to OpenAI;
- [ ] Tests with mocks;
- [ ] Complete all tests;
