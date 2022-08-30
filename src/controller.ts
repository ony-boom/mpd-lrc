import os from "os";
import fs from "fs";
import path from "path";
import { env } from "./config";
import { Lrc } from "lrc-kit";
import { mpc } from "./index";
import renderScreen from "./screen";

const MUSIC_PATH = path.join(os.homedir(), env.musicPath);

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

export const playLyric = (
  songPath: string,
  tittle: string,
  artist: string,
  duration: number
) => {
  const screenTittle = `${artist} - ${tittle}`;
  const lyric = getLyricArray(songPath);
  renderScreen(screenTittle, lyric);
};
