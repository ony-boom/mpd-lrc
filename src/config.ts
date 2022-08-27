import { config } from "dotenv";
import path from "path";
import os from "os";
import fs from "fs";

const CONFIG_PATH = [
    path.join(os.homedir(), ".mpdlrcrc"),
    path.join(os.homedir(), ".config/mpdlrc/config")
]

const getConfig = () => {
  let configFilePath: string = CONFIG_PATH[0];
  
  for (const path of CONFIG_PATH) {
    if (fs.existsSync(path)) {
      configFilePath = path;
      break
    }
  }
  
  return configFilePath;
}

config({
  path: getConfig()
});

export const env = {
  host: process.env.MPD_HOST || "127.0.0.1",
  port: process.env && process.env.MPD_PORT ? +process.env.MPD_PORT : 6600,
  musicPath: process.env.MUSIC_PATH || "Music"
};

