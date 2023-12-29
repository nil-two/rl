#!/usr/bin/env bats

cmd=$BATS_TEST_DIRNAME/../rl
tmpdir=$BATS_TEST_DIRNAME/../tmp
stdout=$BATS_TEST_DIRNAME/../tmp/stdout
stderr=$BATS_TEST_DIRNAME/../tmp/stderr
exitcode=$BATS_TEST_DIRNAME/../tmp/exitcode

setup() {
  mkdir -p -- "$tmpdir"
}

teardown() {
  rm -rf -- "$tmpdir"
}

check() {
  printf "%s\n" "" > "$stdout"
  printf "%s\n" "" > "$stderr"
  printf "%s\n" "0" > "$exitcode"
  "$@" > "$stdout" 2> "$stderr" || printf "%s\n" "$?" > "$exitcode"
}

@test 'rl: reverse lines per character if no arguments passed' {
  src=$(printf "%s\n" $'
  Hello, world!
  100
  100,200
  100,200,300
  100,200,300,400
  ' | sed -e '1d' -e 's/^  //')
  dst=$(printf "%s\n" $'
  !dlrow ,olleH
  001
  002,001
  003,002,001
  004,003,002,001
  ' | sed -e '1d' -e 's/^  //')
  check "$cmd" <<< "$src"
  [[ $(cat "$exitcode") == 0 ]]
  [[ $(cat "$stdout") == $dst ]]
}

@test 'rl: trim lines' {
  src=$(printf "%s\n" $'
    Hello, world!  
    100,200,300  
    100,200,300
  100,200,300  
  100,200,  300
  100  ,200,300
  ' | sed -e '1d' -e 's/^  //')
  dst=$(printf "%s\n" $'
  world!,Hello
  300,200,100
  300,200,100
  300,200,100
  300,200,100
  300,200,100
  ' | sed -e '1d' -e 's/^  //')
  check "$cmd" -s, <<< "$src"
  [[ $(cat "$exitcode") == 0 ]]
  [[ $(cat "$stdout") == $dst ]]
}

@test 'rl: separate lines if -s passed' {
  src=$(printf "%s\n" $'
  Hello, world!
  100
  100,200
  100,200,300
  100,200,300,400
  ' | sed -e '1d' -e 's/^  //')
  dst=$(printf "%s\n" $'
  world!,Hello
  100
  200,100
  300,200,100
  400,300,200,100
  ' | sed -e '1d' -e 's/^  //')
  check "$cmd" -s, <(printf "%s\n" "$src")
  [[ $(cat "$exitcode") == 0 ]]
  [[ $(cat "$stdout") == $dst ]]
}

@test 'rl: separate lines if --separator passed' {
  src=$(printf "%s\n" $'
  Hello, world!
  100
  100,200
  100,200,300
  100,200,300,400
  ' | sed -e '1d' -e 's/^  //')
  dst=$(printf "%s\n" $'
  world!,Hello
  100
  200,100
  300,200,100
  400,300,200,100
  ' | sed -e '1d' -e 's/^  //')
  check "$cmd" --separator=, <(printf "%s\n" "$src")
  [[ $(cat "$exitcode") == 0 ]]
  [[ $(cat "$stdout") == $dst ]]
}

@test 'rl: keep first indent if -i passed' {
  src=$(printf "%s\n" $'
    100
  200
      300
  ' | sed -e '1d' -e 's/^  //')
  dst=$(printf "%s\n" $'
    001
    002
    003
  ' | sed -e '1d' -e 's/^  //')
  check "$cmd" -i <(printf "%s\n" "$src")
  [[ $(cat "$exitcode") == 0 ]]
  [[ $(cat "$stdout") == $dst ]]
}

@test 'rl: keep first indent if --keep-indent passed' {
  src=$(printf "%s\n" $'
    100
  200
      300
  ' | sed -e '1d' -e 's/^  //')
  dst=$(printf "%s\n" $'
    001
    002
    003
  ' | sed -e '1d' -e 's/^  //')
  check "$cmd" --keep-indent <(printf "%s\n" "$src")
  [[ $(cat "$exitcode") == 0 ]]
  [[ $(cat "$stdout") == $dst ]]
}


@test 'rl: print usage if --help passed' {
  check "$cmd" --help
  [[ $(cat "$exitcode") == 0 ]]
  [[ $(cat "$stdout") =~ ^"usage: rl" ]]
}

@test 'rl: print error if nonexistent file passed' {
  check "$cmd" ctcE4_S_c4IsW5JZaxtuaahC7sLb1cWGT9lslCRn
  [[ $(cat "$exitcode") == 1 ]]
  [[ $(cat "$stderr") =~ ^'rl: Can'"'"'t open' ]]
}

@test 'rl: print error if unkown option passed' {
  check "$cmd" --test
  [[ $(cat "$exitcode") == 1 ]]
  [[ $(cat "$stderr") =~ ^'rl: Unknown option' ]]
}

# vim: ft=bash
