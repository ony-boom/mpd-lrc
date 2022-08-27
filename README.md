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
Edit the .env file
```bash
cp .env_example .env
```

Install dependencies:
```bash
yarn
```
Then build and just start it
```bash
yarn build && yarn start
```

# TODO 🗒️
- [ ] Refactor and improve the rendering process
- [ ] Make an executable (.bin, .exe, ...)
- [ ] Better Readme
