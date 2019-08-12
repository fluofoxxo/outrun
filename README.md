# outrun

### Summary

`outrun` ('ORN' for short) is a custom server for Sonic Runners, reverse engineered from the [Sonic Runners Revival](https://sonic.runner.es/) project.

### Goals

  - Achieve 1:1 parity with the original server functionality
  - Maintain speed throughout runtime
  - Consume little resources for conservative deployment

### Current functionality

The game functions "normally" all through timed mode. The story mode (campaign) starts, but the server sends values that break the game flow (i.e. "score to next point" value being ridiculously low). Ancillary menu functions such as viewing your ID may not work as intended, although changing your username does work.

### State

This code should not be considered usable by most people. There are multiple dependencies undocumented here, much of the code is improperly formatted/documented, and the repository is unsuitable for `go get`.

If you wish to run this code, clone the repository and build/run `mess1.go` in the root directory.

### Misc.

Any pull requests deemed code improvements are strongly encouraged. Refactors may be merged into a different branch.

In closed testing, the server consumes around 4 MB of memory. This will probably be increased as the `bbolt` database increases in size.

### Credits

Much thanks to YPwn, whose closest point of online social contact I do not know, for creating and running the Sonic Runners Revival server upon which this project bases much of its code upon.

Much thanks as well to [@Sazpaimon](https://github.com/Sazpaimon) for finding the encryption key I so desparately looked for but could not on my own.
