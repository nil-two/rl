rl
==

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

### compiled binary

See [releases](https://github.com/nil2nekoni/rl/releases)

### go get

```
go get github.com/nil2nekoni/rl
```

Options
-------

### -h, --help

Display a help message.

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

nil2 <nil2@nil2.org>
