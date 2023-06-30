# Prepare you for the interview

## What is the idea of the repository?

Help you prepare for the interview, facilitate searching who the company is, and write about what you need to know or study about the vacancy.
I used OpenAI (ChatGPT) to generate that summary so you need to generate the api_key to use the commands.

[Here](https://elephas.app/blog/how-to-create-openai-api-keys-cl5c4f21d281431po7k8fgyol0) is a post to help to generate the API key.

## How to install:
```
go install github.com/PatrickChagastavares/prepare-interview/cmd/interview@latest
```

## How to use:

### Show version
```
interview version
```

### Show lists of commands and arguments
```
interview -help
```

### Search resume company
```
interview -token="your token openID" company -name="Linkedin" -lang="english"
```
![prepare_interview_company](https://github.com/PatrickChagastavares/prepare-interview/assets/49497853/483b41cc-2252-40ec-a0bb-a27410a8e68f)
### Generate a summary to prepare for the vacancy
```
interview -token="your token openID" prepare -role="Senior Software Engineer" -skills="Golang, Node And AWS" -lang="english"
```
![prepare_interview_role](https://github.com/PatrickChagastavares/prepare-interview/assets/49497853/74e35d19-6147-4d24-a45a-9e684a949beb)

## Todo list:
- [X] Create Readme;
- [X] Implement flags;
- [X] Implement HTTP requests to OpenAI;
