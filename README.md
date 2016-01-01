rl
==

[![Build Status](https://travis-ci.org/kusabashira/rl.svg?branch=master)](https://travis-ci.org/kusabashira/rl)

Reverse lines.

```
$ echo Hello World | rl
dlroW olleH
```

Usage
-----

```
$ rl [OPTION]... [FILE]...

Options:
  -d, --delimiter=DELIM    delimit line by DELIM
  -h, --help               display this help text and exit
  -v, --version            output version information and exit
```

Installation
------------

### go get

```
go get github.com/kusabashira/rl
```

Options
-------

### -h, --help

Display the usage and exit.

### -v, --version

Output the version of rl.

### -d, --delimiter=DELIM

Reverse lines delimited by DELIM.

```sh
$ cat nums
1,20,300,4000
10,20,30,40,50

$ cat nums | rl -d,
4000,300,20,1
50,40,30,20,10
```

License
-------

MIT License

Author
------

kusabashira <kusabashira227@gmail.com>
