import * as events from "events";
import { PlaylistItem } from "mpc-js";
import { playLyric } from "./controller";

const player = new events.EventEmitter();

// prevent rerendering after pause
let cachedCurrent: number;

player.on("play", (currentSong: PlaylistItem) => {
  if (
    currentSong &&
    currentSong.path &&
    currentSong.title &&
    currentSong.artist &&
    currentSong.id !== cachedCurrent
  ) {
    playLyric(currentSong.path, currentSong.title, currentSong.artist);
  }
});

player.on("pause", (currentSong: PlaylistItem) => {
  cachedCurrent = currentSong.id!;
});

export default player;
