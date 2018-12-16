[![Build Status](https://travis-ci.org/Sepuka/gonerator.svg?branch=master)](https://travis-ci.org/Sepuka/gocademy)
[![Go Report Card](https://goreportcard.com/badge/github.com/Sepuka/gonerator)](https://goreportcard.com/report/github.com/Sepuka/gonerator)
This package useful to generate some sequences which you can use for unit tests, your programs, etc...
Usage:
=====
`
fmt.Println(Uint8(100))
[153 26 73 201 240 0 209 249 253 212 112 218 240 96 206 28 169 235 82 34 217 79 75 40 34 214 248 184 40 190 83 207 221 144 160 134 74 152 230 70 156 51 41 197 159 87 154 196 217 148 68 107 93 149 39 56 138 20 49 146 208 83 136 247 93 176 33 77 8 109 62 221 206 254 179 63 242 171 27 198 90 144 54 173 232 176 15 141 15 17 57 19 88 255 137 147 221 54 21]
`
You get the sequence from 0 to 255 (8 bytes) has length 100 numbers (arg0)