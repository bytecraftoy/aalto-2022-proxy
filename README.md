# APIKeyProxy Go

## Build & Run

### Running

```sh
go run main.go
```

### Building

```sh
go build
```

## Usage

Using this software: Once the server is running, send a POST request to http://localhost:8080/ containing JSON file with proper OpenAI-format, and wait for a moment to receive the answer.

This requires that you have OpenAI Apikey in your environment variable, stored in "OpenAI_apikey". Include thea Bearer token (without the "Bearer", just the apikey).

For example:

```json
{
  "model": "text-davinci-002",
  "prompt": "List 10 science fiction books:",
  "temperature": 0.5,
  "max_tokens": 200,
  "top_p": 1.0,
  "frequency_penalty": 0.52,
  "presence_penalty": 0.5
}
```
