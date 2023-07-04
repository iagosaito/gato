# Gato

## What is it?

Gato (which is a Brazilian name for a cat) is an alternative to the "cat" Linux command. It was developed based on a coding challenge created by John Crickett (https://codingchallenges.fyi/challenges/challenge-cat).

It's available for Mac, Linux or Windows
## How is it works?

Gato provides a simple interface for displaying file contents. It can read directly from a file or from stdin. For example:

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

It will execute all the test cases provided in the challenge.

## License

[MIT](https://choosealicense.com/licenses/mit/)

