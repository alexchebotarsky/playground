# Wavy Word Generator

> Disclaimer: I did not come up with this challenge, but the solution is uniquely mine.

## Challenge description

A wavy word is a string that stretches up and down across multiple lines to form waves. A word's waviness is how many lines away from the baseline the word can stretch. Your program should take a word as a string and output its wavy form.

A wavy word always starts out going down.

If waviness is 0, the output should be the same as the input word.

Empty lines must be included.

## Input description

1. A string word to make wavy.
2. An integer waviness that decides how wavy to make the word.

## Output

A wavy word spanning exactly `waviness * 2 + 1` lines. Each line should be the same length, including spaces.

Example:

```
   y   d
w v w r
 a   o
```
