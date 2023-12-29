rl
==

[![CI](https://github.com/nil-two/rl/actions/workflows/test.yml/badge.svg)](https://github.com/nil-two/rl/actions/workflows/test.yml)

Reverse lines.

```
$ echo Hello World | rl
dlroW olleH
```

Usage
-----

```
$ rl [<option(s)>] [<file(s)>]
reverse lines.

options:
  -s, --separator=SEP  separate lines by SEP
  -i, --keep-indent    keep the first indent as indents
      --help           print usage and exit
```

Requirements
------------

- Perl (5.8.0 or later)

Installation
------------

1. Copy `rl` into your `$PATH`.
2. Make `rl` executable.

### Example

```
$ curl -L https://raw.githubusercontent.com/nil-two/rl/master/rl > ~/bin/rl
$ chmod +x ~/bin/rl
```

Note: In this example, `$HOME/bin` must be included in `$PATH`.

Options
-------

### -s, --separator=\<sep\>

Separate a line by sep.
Default value is empty string.

```
$ echo foo,bar,baz | rl -s,
baz,bar,foo
```

### -i, --keep-indent

Keep the first indent as indents.
Default is disabled.

```
$ printf "%s\n" '  foo' 'bar' '    baz' | rl -i
  oof
  rab
  zab
```

### --help

Print usage and exit.

```
$ rl --help
(Print usage)
```

License
-------

MIT License

Author
------

nil2 <nil2@nil2.org>
