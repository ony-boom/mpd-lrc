# mpd-lrc

A simple terminal based **LRC** file player, made up with the [Music Player Daemon](https://www.musicpd.org/) and node.

[woozy.webm](https://user-images.githubusercontent.com/84435474/187069686-a1f3f2b7-0584-4ba6-b384-b8d98c3f7dd6.webm)

# Usage

> **_NOTE:_** For this, you should have mpd up and running.
> [hint](https://wiki.archlinux.org/title/Music_Player_Daemon) 💡

 You can just download the binary from [release](https://github.com/ony-boom/mpd-lrc/releases) and just configure it, or<br>
 <br>
Clone this repo:
```bash
git clone https://github.com/ony-boom/mpd-lrc.git && cd mpd-lrc
```

### Configuration
The configuration file is loaded from: <br>
`$HOME/.mpdlrcrc` or `$HOME/.config/mpdlrc/config`

```bash
mkdir ~/.config/mpdlrc
cp .config_example ~/.config/mpdlrc/config # or just ~/.mpdlrcrc
```

Install dependencies:
```bash
yarn
```
Then build, see [pkg](https://www.npmjs.com/package/pkg) for more info. <br>
Edit `build` scripts in `package.json`  for your need.(Default build is for linux)
```bash
yarn build
```
