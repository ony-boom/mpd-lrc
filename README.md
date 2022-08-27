# mpd-lrc

A simple terminal based **LRC** file player, made up with the [Music Player Daemon](https://www.musicpd.org/) and node.

https://user-images.githubusercontent.com/84435474/186978513-bcdc2614-c4d2-4365-86fb-cfee745d87aa.mp4


# Usage

> **_NOTE:_** For this, you should have mpd up and running.
> [hint](https://wiki.archlinux.org/title/Music_Player_Daemon) 💡

For now, Just clone this repo:
```bash
git clone https://github.com/ony-boom/mpd-lrc.git && cd mpd-lrc
```

Configure, the configuration file is loaded from: <br>
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

# TODO 🗒️
- [ ] Refactor and improve the rendering process

