# Server overload examples

There are three branches in this repository, each one with a different example 
of how to handle server overload.

## Branches

- [`main`](https://github.com/ivanlemeshev/serveroverload/tree/main): The branch contains an example of server overload.
- [`rate-limiting`](https://github.com/ivanlemeshev/serveroverload/tree/rate-limiting): The branch contains an example of how to use rate limiting to 
  handle server overload.
- [`load-shedding`](https://github.com/ivanlemeshev/serveroverload/tree/load-shedding): The branch contains an example of how to use load shedding to 
  handle server overload.

## Running the examples

```bash
$ git checkout <branch name>
$ make run-service
$ run-k6-tests
$ make stop-service
```
