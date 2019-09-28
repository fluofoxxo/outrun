# outrun

### Summary

Outrun is a custom server for Sonic Runners, reverse engineered from the [Sonic Runners Revival](https://sonic.runner.es/) project.

### Goals

  - Implement all non debug endpoints as fulfilled by the original server
  - Maintain speed throughout runtime
  - Consume little resources for conservative deployment

### Current functionality

Notable:
  - Timed Mode
  - Story Mode (Functional, in progress)
  - Ring/Red Star Ring keeping
  - Functional shop
  - Character/Chao equipping
  - Character leveling and progression
  - Item roulette functionality

Functional:
  - Android and iOS support
  - High score keeping
  - Notifications (pending revamp)
  - Ticker notices
  - Small database size and memory footprint
  - Low CPU usage
  - Descriptive error reporting and handling

### Building

1. [Download and install Go 1.12](https://golang.org/dl/) (project tested on Go 1.12.4)
2. [Download and install Git](https://git-scm.com/downloads) (for `go get`)
3. Set your [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) environment variable
4. Open a terminal/command prompt
5. Use `cd` ([Windows,](https://www.digitalcitizen.life/command-prompt-how-use-basic-commands) [Linux/macOS](https://www.macworld.com/article/2042378/master-the-command-line-navigating-files-and-folders.html)) to navigate to a directory of choice
6. Run `go get github.com/fluofoxxo/outrun` and wait until the command line returns
7. Run `go build github.com/fluofoxxo/outrun` and wait until the build is complete
8. Run the produced executable (`outrun.exe` on Windows, `outrun` on Linux/macOS)

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

iOS app file (from Windows):
1. Acquire a Sonic Runners Revival IPA from the [Sonic Runners Revival website.](https://sonic.runner.es/) (`sr203_revival.ipa`)
2. Create a folder in a safe place (Ex. Desktop/PalmboundWorkspace)
3. Move the Revival IPA to PalmboundWorkspace
4. Download xdelta3 (`xdelta3-3.1.0-x86_64.exe.zip`) from [jmacd/xdelta-gpl](https://github.com/jmacd/xdelta-gpl/releases)
5. Extract xdelta3 executable file (Ex. `xdelta3-3.1.0-x86_64.exe`) to PalmboundWorkspace
6. Download [`Revival_iOS_to_Palmbound_v2_iOS.vcdiff`](http://pbassets.fluofoxxo.pw:9002/Revival_iOS_to_Palmbound_v2_iOS.vcdiff) and move it to PalmboundWorkspace
7. Open PalmboundWorkspace in File Explorer, hold shift and right click on any location with visibly empty space in the folder, and click "Open PowerShell window here"
8. In the window that opens, enter `.\xdelta3-3.1.0-x86_64.exe -d -s sr203_revival2.ipa Revival_iOS_to_Palmbound_v2_iOS.vcdiff Palmbound.ipa`
9. Use the newly created `Palmbound.ipa` file in the common instructions below. If you are using a jailed device (your device is not jailbroken), use **Common iOS Installation (Windows, jailed)**. If you are using a jailbroken device, use **Common iOS Installation (Windows, jailbroken)**.

Common iOS Installation (Windows, jailed):
1. Sign into your [Apple ID](https://appleid.apple.com) online
2. Navigate to the Security section of your profile
3. Under App Specific Password, click Generate Password
4. Enter 'Palmbound' as the label, and click Create
5. Copy the code somewhere safe for later use
6. Install [iTunes](https://support.apple.com/downloads/itunes) and sign into it within the application
7. Download [Cydia Impactor](http://www.cydiaimpactor.com/), extract the downloaded archive somewhere, and run the included Impactor.exe
8. Plug the iOS device into the computer using an appropriate USB cable, and unlock the device and, if prompted to trust the computer, tap Trust
9. Drag `Palmbound.ipa` onto the Impactor window
10. Enter your Apple ID's email, click OK, then enter your App Specific Password, then click OK, and wait for completion
11. (Potentially unrequired) Go to Settings -> Profiles or Device Management -> Tap the entry with your Apple ID email -> Tap 'Trust'
12. Launch the game!

Cydia Impactor errors can be investigated [here.](https://cydiaimpactor.online/cydia-impactor-errors/)

If the game no longer works after a period of time, please refer to '**iOS App Recertification**'.

Common iOS Installation (Windows, jailbroken):
1. Through Cydia, install [AppSync Unified](https://cydia.akemi.ai/?page/net.angelxwind.appsyncunified)
2. Install [Filza](https://filza.net/download/) on the device
3. Transfer the file to the device using a local web server, Dropbox, or other method of transferring files to an iOS device
4. [Export the file and select Copy to Filza](https://kubadownload.com/news/appsync-unified)
5. Open the file and tap Install
6. Launch the game!

iOS App Recertification:
1. Open iTunes and plug your iOS device into the computer using an appropriate USB cable, and unlock the device
2. Click on your device in the iTunes device panel
3. Click File Sharing
4. Click on the game
5. Click on `ifrn.game` and save it in a safe location; do the same for `sfrn.game`
6. On the device, tap and hold the app on the home screen and delete it
7. Install the game as per the instructions in **Common iOS Installation (Windows, jailed)**, but do not open the game
8. Add the `ifrn.game` and `sfrn.game` files into the app through iTunes
9. Launch the game

### Credits

Much thanks to:
  - **YPwn**, whose closest point of online social contact I do not know, for creating and running the Sonic Runners Revival server upon which this project bases much of its code upon.
  - **[@Sazpaimon](https://github.com/Sazpaimon)** for finding the encryption key I so desparately looked for but could not on my own.
  - **nacabaro** (nacabaro#2138 on Discord) for traffic logging and the discovery of **[DaGuAr](https://www.youtube.com/user/Gorila5)**'s asset archive.
