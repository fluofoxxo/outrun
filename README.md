# outrun

### Summary

Outrun is a custom server for Sonic Runners, reverse engineered from the [Sonic Runners Revival](https://sonicrunners.com/) (Now uses Outrun) project.

### Goals

  - Implement all non debug endpoints as fulfilled by the original server
  - Maintain speed throughout runtime
  - Consume little resources for conservative deployment

### Current functionality

Notable:
  - Timed Mode
  - Story Mode
  - Ring/Red Star Ring keeping
  - Functional shop
  - Character/Chao equipping
  - Character leveling and progression
  - Item/Chao roulette functionality
  - Events
  - Basic ranking

Functional:
  - Android and iOS support
  - High score keeping
  - In game notices
  - Deep configuration options
  - Powerful RPC control functions
  - Ticker notices
  - Small database size and memory footprint
  - Low CPU usage
  - Analytics support

### Building

1. [Download and install Go 1.15.8](https://golang.org/dl/) (**NOTE:** Go 1.16+ currently break the ability to build some or all of Outrun, such as RPC programs. A fix should hopefully be coming soon!)
2. [Download and install Git](https://git-scm.com/downloads) (for `go get`)
3. Set your [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) environment variable
4. Open a terminal/command prompt
5. Use `cd` ([Windows,](https://www.digitalcitizen.life/command-prompt-how-use-basic-commands) [Linux/macOS](https://www.macworld.com/article/2042378/master-the-command-line-navigating-files-and-folders.html)) to navigate to a directory of choice
6. Run `go get github.com/fluofoxxo/outrun` and wait until the command line returns
7. Run `go build github.com/fluofoxxo/outrun` and wait until the build is complete
8. Run the produced executable (`outrun.exe` on Windows, `outrun` on Linux/macOS)

Binary releases can be found [in the releases tab.](https://github.com/fluofoxxo/outrun/releases)

#### Modifying an APK to connect to your instance (from Windows)

1. Install [dnSpy](https://github.com/0xd4d/dnSpy/releases) (dnSpy-netcore-win64.zip)
2. Install [7-Zip](https://www.7-zip.org/download.html)
3. Install [ZipSigner](https://www.apkmirror.com/apk/ken-ellinwood/zipsigner/zipsigner-3-4-release/zipsigner-3-4-android-apk-download/) on an Android device or emulator
4. Open a Sonic Runners v2.0.3 APK file with 7-Zip
5. Navigate to assets/bin/Data/Managed and extract all the DLL files to their own folder
6. Open Assembly-CSharp.dll in dnSpy
7. Open the class `NetBaseUtil`, and find the variable `mActionServerUrlTable `
8. Edit every string in the `mActionServerUrlTable` array to `http://<IP>:<PORT>/` where `<IP>` is replaced by the IP for your instance and `<PORT>` is replaced by the port for your instance (Default: 9001)
9. Repeat step 7 for `mSecureActionServerUrlTable`
10. If you have an assets server, use its IP and port to replace the values in `mAssetURLTable` and `mInformationURLTable` to `http://<IP>:<PORT>/assets/` and `http://<IP>:<PORT>/information/` respectively
11. Click File -> Save Module... and save the DLL file
12. Drag the newly saved Assembly-CSharp.dll back into assets/bin/Data/Managed in 7-Zip, confirming to overwrite if asked
13. Transfer the APK to an Android device and use ZipSigner to sign it
14. Install the APK


### Misc.

Any pull requests deemed code improvements are strongly encouraged. Refactors may be merged into a different branch.

#### Palmbound (Public Development Server)

As of 1 September, 2019, an instance of this server named Palmbound is currently publicly accessible. App file downloads can be acquired by joining the Palmbound Download Discord server: https://discord.gg/eeQAe8R

### Credits

Much thanks to:
  - **YPwn**, whose closest point of online social contact I do not know, for creating and running the Sonic Runners Revival server upon which this project bases much of its code upon.
  - **[@Sazpaimon](https://github.com/Sazpaimon)** for finding the encryption key I so desparately looked for but could not on my own.
  - **nacabaro** (nacabaro#2138 on Discord) for traffic logging and the discovery of **[DaGuAr](https://www.youtube.com/user/Gorila5)**'s asset archive.

#### Additional assistance
  - Story Mode items
    - lukaafx (Discord @Kalu04#3243)
    - [TemmieFlakes](https://twitter.com/pictochat3)
    - SuperSonic893YT
