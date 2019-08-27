# outrun

### Summary

`outrun` ('ORN' for short) is a custom server for Sonic Runners, reverse engineered from the [Sonic Runners Revival](https://sonic.runner.es/) project.

### Goals

  - Achieve 1:1 parity with the original server functionality
  - Maintain speed throughout runtime
  - Consume little resources for conservative deployment

### Current functionality

Notable:
  - Timed Mode
  - Ring/Red Star Ring keeping
  - Character/Chao equipping

Functional:
  - High score keeping
  - Notifications (pending revamp)
  - Ticker notices

### Building

1. [Download and install Go 1.12](https://golang.org/dl/) (project tested on Go 1.12.4)
2. [Download and install Git](https://git-scm.com/downloads) (for `go get`)
3. Set your [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) environment variable
4. Open a terminal/command prompt
5. Use `cd` ([Windows,](https://www.digitalcitizen.life/command-prompt-how-use-basic-commands) [Linux/macOS](https://www.macworld.com/article/2042378/master-the-command-line-navigating-files-and-folders.html) to navigate to a directory of choice
6. Run `go build github.com/fluofoxxo/outrun` and wait until the build is complete
7. Run the produced executable (`outrun.exe` on Windows, `outrun` on Linux/macOS)

Binary releases may be provided in more mature stages of the repository.

### Misc.

Any pull requests deemed code improvements are strongly encouraged. Refactors may be merged into a different branch.

### Credits

Much thanks to:
  - **YPwn**, whose closest point of online social contact I do not know, for creating and running the Sonic Runners Revival server upon which this project bases much of its code upon.
  - **[@Sazpaimon](https://github.com/Sazpaimon)** for finding the encryption key I so desparately looked for but could not on my own.
  - **nacabaro** (nacabaro#2138 on Discord) for traffic logging and the discovery of **[DaGuAr](https://www.youtube.com/user/Gorila5)**'s asset archive.
