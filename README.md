# outrun

### Summary

`outrun` ('ORN' for short) is a custom server for Sonic Runners, reverse engineered from the [Sonic Runners Revival](https://sonic.runner.es/) project.

### Goals

  - Achieve 1:1 parity with the original server functionality
  - Maintain speed throughout runtime
  - Consume little resources for conservative deployment

### Current functionality

Most of the login sequence functions. However, some of the functions which should retain player data do not retain such data. These are typically signed with in-code comments.

### State

This code should not be considered usable by most people. There are multiple dependencies undocumented here, much of the code is improperly formatted/documented, and the repository is unsuitable for `go get`.

If you wish to run this code, clone the repository and build/run `mess1.go` in the root directory.

### Misc.

Any pull requests deemed code improvements are strongly encouraged. Refactors may be merged into a different branch.

In closed testing, the server consumes around 4 MB of memory. This will probably be increased as the `bbolt` database increases in size.

### Credits

Much thanks to YPwn, whose closest point of online social contact I do not know, for creating and running the Sonic Runners Revival server upon which this project bases much of its code upon.

Much thanks as well to [@Sazpaimon](https://github.com/Sazpaimon) for finding the encryption key I so desparately looked for but could not on my own.
