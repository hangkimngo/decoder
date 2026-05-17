# Decoder

## Project Overview

Decoder is a CLI application written in Go that encodes and decodes text using a custom bracket-based compression format.

Example:

```text
ABC[10 D]EFG  
```

becomes:

```text
ABCDDDDDDDDDDEFG  
```  

**And:**  

```text
[5 #][5 -_]-[5 #]  
```

becomes:

```text
#####-_-_-_-_-_-#####  
```


The project supports single-line and multi-line processing, file input, and ASCII art encoding/decoding.

---

## Setup and Installation

Requirements:

* Go 1.20+

Clone the repository:

```bash
git clone https://github.com/hangkimngo/decoder.git
cd decoder
```


---

## Usage Guide

Show help:

```bash
go run . -h
```

Decode text:

```bash
go run . "ABC[10 D]EFG"
```

Returns:

```text
ABCDDDDDDDDDDEFG  
```

Encode text:

```bash
go run . -encode "aaaaabbbb"
```

Returns:

```text
[5 a][4 b]
```

Decode multi-line input:

```bash
go run . -multi
```

Encode multi-line input:

```bash
go run . -multi -encode
```

For multi-line mode, enter the input line by line and finish with:

```text
Ctrl + D
```

Use file input:

```bash
go run . -file input.txt
```

Encode file input:

```bash
go run . -encode -file art.txt
```

---

## Additional Features

* Multi-line decoder
* Encoder mode
* File input support
