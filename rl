#!/usr/bin/perl
use strict;
use warnings;
use utf8;
use Encode qw(decode_utf8);
use File::Basename qw(basename);                                             
use Getopt::Long qw(:config posix_default gnu_compat permute bundling no_ignore_case);
use Data::Dumper;

use open IO => ":encoding(UTF-8)";
binmode STDIN, ":encoding(UTF-8)";
binmode STDOUT, ":encoding(UTF-8)";
binmode STDERR, ":encoding(UTF-8)";

my $cmd_name  = basename($0);
my $cmd_usage = <<EOL;
usage: $cmd_name [<option(s)>] [<file(s)>]
reverse lines.

options:
  -s, --separator=SEP  separate each line by SEP
  -i, --keep-indent    keep indents of each line
      --help           print usage and exit
EOL

sub print_error_and_abort {
    my ($error) = @_;
    chomp $error;
    print STDERR "$cmd_name: $error\n";
    exit 1;
}

sub read_lines_from_argf {
    my $lines = [];
    while (my $line = <>) {
        chomp $line;
        push @$lines, $line;
    }
    return $lines;
}

sub reverse_lines {
    my ($lines, $separator, $keep_indent) = @_;
    my $reversed_lines     = [];
    my $separator_pattern  = qr/@{[quotemeta($separator)]}/;
    for my $line (@$lines) {
        my $prefix  = "";
        my $postfix = "";
        if ($keep_indent) {
            $prefix  = $line =~ s/^(\s*).*/$1/r;
            $postfix = $line =~ s/.*?(\s*)$/$1/r;
            $line    = $line =~ s/^\s*|\s*$//gr;;
        }
        my $cells = [split($separator_pattern, $line)];
        for (my $i = 0; $i < @$cells / 2; $i++) {
            ($cells->[$i], $cells->[$#$cells - $i]) = ($cells->[$#$cells - $i], $cells->[$i]);
        }
        my $reversed_line = join($separator, @$cells);
        if ($keep_indent) {
          $reversed_line = $prefix . $reversed_line . $postfix;
        }
        push @$reversed_lines, $reversed_line;
    }
    return $reversed_lines;
}

sub dump_lines {
    my ($lines) = @_;
    for my $line (@$lines) {
        print "$line\n";
    }
}

sub main {
    local $SIG{__WARN__} = \&print_error_and_abort;

    my $separator   = "";
    my $keep_indent = 0;
    my $is_help     = 0;
    foreach (@ARGV) {
        $_ = decode_utf8($_);
    }
    GetOptions(
        "s|separator=s" => \$separator,
        "i|keep-indent" => \$keep_indent,
        "help"          => \$is_help,
    );
    if ($is_help) {
        print $cmd_usage;
        return;
    }

    eval {
        my $lines          = read_lines_from_argf();
        my $reversed_lines = reverse_lines($lines, $separator, $keep_indent);
        dump_lines($reversed_lines);
    };
    if ($@) {
        print_error_and_abort($@);
    }
}
main;
