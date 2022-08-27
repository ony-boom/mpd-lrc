import os from "os";
import fs from "fs";
import path from "path";
import { env } from "./config";
import { Lrc, Lyric } from "lrc-kit";
import { mpc } from "./index";
import { clearInterval } from "timers";
import { setBoxes, setScreen } from "./screen";

const MUSIC_PATH = path.join(os.homedir(), env.musicPath);
let lyrics: Lyric[] = [];
let interval: NodeJS.Timer;

const getLrcPath = (songPath: string): string => {
  return path.dirname(path.join(MUSIC_PATH, songPath));
};

const getLrcFile = (songPath: string): string | undefined => {
  const lrcDir = getLrcPath(songPath);
  const songName = songPath.substring(
    songPath.lastIndexOf("/") + 1,
    songPath.lastIndexOf(".")
  );

  const files = fs.readdirSync(lrcDir);
  const lrcFile = files.find((file) => {
    const filename = file.substring(0, file.lastIndexOf("."));

    if (file.endsWith(".lrc") && filename === songName) {
      return file;
    }
  });

  if (lrcFile) return path.join(lrcDir, lrcFile);
};

const parsedLrc = (songPath: string): Lrc | undefined => {
  const lrcFile = getLrcFile(songPath);

  if (!lrcFile) return;

  const lrcString = fs.readFileSync(lrcFile, "utf-8");

  return Lrc.parse(lrcString);
};

const getElapsedTime = async () => {
  return (await mpc.status.status()).elapsed!;
};

const getLyricArray = (songPath: string) => {
  const lrc = parsedLrc(songPath);

  console.clear();

  if (!lrc) return;

  const { lyrics: lyric } = lrc;
  return lyric;
};



export const playLyric = (songPath: string, tittle: string, artist: string) => {
  if (interval) {
    clearInterval(interval);
  }

  // cleanup the lyrics array
  lyrics = [];
  lyrics = getLyricArray(songPath)!;

  let textContent = "";

  const screen = setScreen();
  screen.title = `${artist} - ${tittle}`;

  const { headerBox, lyricsBox: box } = setBoxes(screen);
  if (!lyrics) {
    textContent = `No Lyrics. (No LRC found in the directory of: "${tittle} by ${artist}")`;
    box.setContent(textContent);
    screen.append(box);
    screen.render();
    return;
  }

  for (const lyric of lyrics) {
    if (lyric) textContent += `${lyric.content}\n`;
  }

  box.setContent(textContent);
  headerBox.setContent(`{bold}${artist} - ${tittle}{/bold}`);
  screen.append(box);
  screen.insertBefore(headerBox, box);

  interval = setInterval(async () => {
    let elapsed = await getElapsedTime();
    elapsed = Math.floor(elapsed);

    const syncedLyrics = lyrics && lyrics.map((lyric) => {
      if (Math.floor(lyric.timestamp) === elapsed) {
        return `{yellow-fg}${lyric.content}{yellow-fg}\n{/}`;
      }
      return lyric.content;
    });

    box.setContent(syncedLyrics && syncedLyrics.join("\n"));

    // preserve lyrics state.
    // keep passed lyrics color to yellow
    // TODO: find a better way to do it🥲
    lyrics = lyrics && lyrics.map((l, idx) => ({
      content: syncedLyrics[idx],
      timestamp: l.timestamp,
    }));

    screen.render();
  }, 100); // 100 seems to be the perfect one
};
