import os from "os";
import fs from "fs";
import path from "path";
import { env } from "./config";
import { Lrc, Lyric } from "lrc-kit";
import { mpc } from "./index";
import { setBoxes, setScreen } from "./screen";
import { Widgets } from "blessed";

const MUSIC_PATH = path.join(os.homedir(), env.musicPath);
let lyrics: Lyric[] = [];
let oldScreen: Widgets.Screen;
const passedLyrics = new Map<number, string>();

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

const cleanup = () => {
  lyrics = [];
  passedLyrics.clear();
  oldScreen && oldScreen.destroy();
};

export const playLyric = (songPath: string, tittle: string, artist: string) => {
  // cleanup

  cleanup();

  lyrics = getLyricArray(songPath)!;
  let textContent = "";

  const screen = setScreen();
  screen.title = `${artist} - ${tittle}`;

  // with the cleanup(), this prevent screen tearing, i don't know really why though
  oldScreen = screen;

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

  // TODO: Find a better way to do sync lyrics
  setInterval(async () => {
    let elapsed = await getElapsedTime();
    elapsed = Math.floor(elapsed);
    let syncedLyrics = "";
    lyrics &&
      lyrics.map((lyric, idx) => {
        if (Math.floor(lyric.timestamp) === elapsed) {
          // this extra if check is for some optimization, i don't really how to explain it
          // but it just prevent some rerendering
          if (!passedLyrics.has(lyric.timestamp)) {
            const newText = `{yellow-fg}${lyric.content}{yellow-fg}\n{/}`;
            lyrics[idx].content = newText;
            syncedLyrics += newText;
            passedLyrics.set(lyric.timestamp, lyric.content);
            return newText;
          }
        }
        syncedLyrics += `${lyric.content}\n`;
        return lyric.content;
      });

    box.setContent(syncedLyrics);
    screen.render();
  }, 100);
};
