# outrun

### Summary

Outrun is a custom server for Sonic Runners, reverse engineered from the [Sonic Runners Revival](https://sonic.runner.es/) project.

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
5. Use `cd` ([Windows,](https://www.digitalcitizen.life/command-prompt-how-use-basic-commands) [Linux/macOS](https://www.macworld.com/article/2042378/master-the-command-line-navigating-files-and-folders.html)) to navigate to a directory of choice
6. Run `go build github.com/fluofoxxo/outrun` and wait until the build is complete
7. Run the produced executable (`outrun.exe` on Windows, `outrun` on Linux/macOS)

Binary releases may be provided in more mature stages of the repository.

### Misc.

Any pull requests deemed code improvements are strongly encouraged. Refactors may be merged into a different branch.

#### Palmbound (Public Development Server)

As of 1 September, 2019, an instance of this server named Palmbound is currently publicly accessible.

##### Installation

Android app file (from Windows):
1. Acquire a Sonic Runners Revival APK from the [Sonic Runners Revival website.](https://sonic.runner.es/) (`sr203_revival2.apk`)
2. Create a folder in a safe place (Ex. Desktop/PalmboundWorkspace)
3. Move the Revival APK to PalmboundWorkspace
4. Download xdelta3 (`xdelta3-3.1.0-x86_64.exe.zip`) from [jmacd/xdelta-gpl](https://github.com/jmacd/xdelta-gpl/releases)
5. Extract xdelta3 executable file (Ex. `xdelta3-3.1.0-x86_64.exe`) to PalmboundWorkspace
6. Download [`Revival_to_Palmbound_v1.vcdiff`](http://pbassets.fluofoxxo.pw:9002/Revival_to_Palmbound_v1.vcdiff) and move it to PalmboundWorkspace
7. Open PalmboundWorkspace in File Explorer, hold shift and right click on any location with visibly empty space in the folder, and click "Open PowerShell window here"
8. In the window that opens, enter `.\xdelta3-3.1.0-x86_64.exe -d -s sr203_revival2.apk Revival_to_Palmbound_v1.vcdiff palmbound.apk`
9. Move the newly created `palmbound.apk` to an Android device and install it

Android app file (from Android):
1. Acquire a Sonic Runners Revival APK from the [Sonic Runners Revival website.](https://sonic.runner.es/) (`sr203_revival2.apk`)
2. Download [`Revival_to_Palmbound_v1.vcdiff`](http://pbassets.fluofoxxo.pw:9002/Revival_to_Palmbound_v1.vcdiff)
3. Install [UniPatcher](https://play.google.com/store/apps/details?id=org.emunix.unipatcher) or any similar tool for xdelta3 patching
4. Set the patch file to `Revival_to_Palmbound_v1.vcdiff`
5. Set the "ROM" file to `sr203_revival2.apk` (Revival's APK)
6. Set the Output file to `palmbound.apk`
7. Install the new `palmbound.apk`

### Credits

Much thanks to:
  - **YPwn**, whose closest point of online social contact I do not know, for creating and running the Sonic Runners Revival server upon which this project bases much of its code upon.
  - **[@Sazpaimon](https://github.com/Sazpaimon)** for finding the encryption key I so desparately looked for but could not on my own.
  - **nacabaro** (nacabaro#2138 on Discord) for traffic logging and the discovery of **[DaGuAr](https://www.youtube.com/user/Gorila5)**'s asset archive.
