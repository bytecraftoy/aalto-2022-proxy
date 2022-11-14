<<<<<<< HEAD
# APIKeyProxy #

## Build & Run ##

```sh
$ cd apikeyproxy
$ sbt
> Jetty / start
```

If you want to quit:
```
> Jetty / stop
```

For restart upon changing files:
```
> ~Jetty / start
```


Using this software: Once the Jetty runner is running, send a POST request to http://localhost:8080/ containing JSON file with proper OpenAI-format, and wait for a moment to receive the answer.


This requires that you have OpenAI Apikey in your environment variable, stored in "OpenAI_apikey". Include thea Bearer token (without the "Bearer", just the apikey).

For example:

```
{
    "model": "text-davinci-002",
    "prompt": "List 10 science fiction books:",
    "temperature": 0.5,
    "max_tokens": 200,
    "top_p": 1.0,
    "frequency_penalty": 0.52,
    "presence_penalty": 0.5,
    "stop": [
        "11."
    ]
}```


Authors:
Antti Halava (Bytecraft)
=======
# aalto-2022-proxy
>>>>>>> 9cae11927ba50f01ddfb0d2c212be61420f27417
