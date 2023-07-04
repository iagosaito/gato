# Gato

## What is it?

Gato (it's a brazilian name for cat) is a stupid alternative for cat linux command. It's based on coding challenge created by John Crickett (https://codingchallenges.fyi/challenges/challenge-cat).

It's available for Mac, Linux or Windows
## How is it works?

Gato provide a simple interface in order to display file contents. It can read directly from file or from stdin. For example: 

```bash
  ./gato test1.txt (Directly from file)
  head -n1 test1.txt | ./gato - (Read from stdin)
```

To see all options available, run: 

```bash
  ./gato -h
```
## Running Tests

To run tests, run the following command

```bash
  make test
```

It will run all the test cases provided in the challenge

## License

[MIT](https://choosealicense.com/licenses/mit/)

