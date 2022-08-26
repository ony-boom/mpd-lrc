import { config } from "dotenv";

config();

export const env = {
  host: process.env.MPD_HOST || "127.0.0.1",
  port: process.env && process.env.MPD_PORT ? +process.env.MPD_PORT : 6600,
  musicPath: process.env.MUSIC_PATH || "Music"
};

