# Solutions

## Ruby

Initial Solution. Reads whole file in, creates a collection of hashes
and runs over those via Array#map to normalize the data. Works (sort of,
DST bug, and not structured well for testing).

Tested with system ruby on:

```
      System Version: macOS 11.6 (20G165)
      Kernel Version: Darwin 20.6.0

  ruby 2.6.3p62 (2019-04-16 revision 67580) [universal.x86_64-darwin20]
```

If in the current directory, example run:

```shell
  ./normalizer.rb < ../sample.csv > output.csv
```

## Go

Follow up to initial solution. This one is a lot better, more conducive to
testing, will run much much faster, and processes line by line, versus
reading the whole file in. Therefore this one is much more scalable. It's also
a little more pleasing to the senses and just has a better "feel," Lastly,
the DST bug is not here.

Tested on the same mac as above, go version:

```
  go version go1.16.2 darwin/amd64
```

Example run (from this directory):

```shell
  go run ./normalizer.go < ../sample.csv > output.csv
```
