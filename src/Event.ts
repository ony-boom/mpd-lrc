import * as events from "events";
import { PlaylistItem } from "mpc-js";
import { playLyric } from "./controller";
import { Widgets } from "blessed";

const player = new events.EventEmitter();

let oldScreen: Widgets.Screen;

// prevent re-rendering after pause
let cachedCurrent: number;

const handleChange = (newSong: number) => {
  cachedCurrent = newSong;
  oldScreen.children.forEach((c) => c.destroy());
  oldScreen.destroy();
};

player.on("play", (currentSong: PlaylistItem) => {
  if (
    currentSong &&
    currentSong.path &&
    currentSong.title &&
    currentSong.artist &&
    currentSong.id !== cachedCurrent
  ) {
    oldScreen = playLyric(
      currentSong.path,
      currentSong.title,
      currentSong.artist
    );
  }
});

player.on("pause", (currentSong: PlaylistItem) => {
  cachedCurrent = currentSong.id!;
});

player.on("next", handleChange);
player.on("previous", handleChange);

export default player;
