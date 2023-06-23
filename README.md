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
### Generate a summary to prepare for the vacancy
```
go run cmd/interview/main.go -token="your token openID" prepare -role="Software Developer" -skills="Golang, Node e GCP" -lang="english"
```



## Todo list:
- [ ] Create Readme;
- [ ] Implement helper;
- [ ] Tests with mocks;
- [ ] Implement flags;
- [ ] Implement HTTP requests to OpenAI;
- [ ] Complete all tests;
- [ ] Implement a method to safe openai_key; 