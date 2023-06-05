# mpd-lrc

A simple terminal based **LRC** file player, made up with the [Music Player Daemon](https://www.musicpd.org/) and node.

[mdp-lrc-demo.webm](https://user-images.githubusercontent.com/84435474/187713102-868838df-1761-4fe6-abd1-772153041cb9.webm)


# Usage

> **_NOTE:_** For this, you should have mpd up and running.
> [hint](https://wiki.archlinux.org/title/Music_Player_Daemon) 💡

Clone this repo:
```bash
git clone https://github.com/ony-boom/mpd-lrc.git && cd mpd-lrc
```

Install dependencies:
```bash
yarn
```
Then build, see [pkg](https://www.npmjs.com/package/pkg) for more info. <br>
Edit `build` scripts in `package.json`  for your need.
```bash
yarn build
```


### Configuration
The configuration file is loaded from: <br>
`$HOME/.mpdlrcrc` or `$HOME/.config/mpdlrc/config`

```bash
mkdir ~/.config/mpdlrc
cp .config_example ~/.config/mpdlrc/config # or just ~/.mpdlrcrc
```
