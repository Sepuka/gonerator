[![Build Status](https://travis-ci.org/Sepuka/gonerator.svg?branch=master)](https://travis-ci.org/Sepuka/gocademy)
[![Go Report Card](https://goreportcard.com/badge/github.com/Sepuka/gonerator)](https://goreportcard.com/report/github.com/Sepuka/gonerator)

This package useful to generate some sequences which you can use for unit tests, your programs, etc...

Usage:
=====
```
fmt.Println(SeqUint8(2500))
[153 26 73 201 240 0 209 249 253 212 112 218 240 96 206 28 169 235 82 ... 147 221 54 21]
```
You've get the sequence from 0 to 255 (8 bytes) has length 2500 numbers (arg0)

Below contents distribution of points
![unsigned integer 8 bytes](https://user-images.githubusercontent.com/825371/50472420-c33ad500-09c9-11e9-91d4-a1aca780244e.png)