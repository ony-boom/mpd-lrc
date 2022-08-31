import { MPC } from "mpc-js";
import { env } from "./config";

import player from "./Event";

export const mpc = new MPC();

(async () => {
  await mpc.connectTCP(env.host, env.port);

  if (mpc.isReady) {
    console.log("Connected");
    
    mpc.on("", () => {
    
    })

    const currentSong = await mpc.status.currentSong();
    let currentSongId = currentSong.id;

    let status = await mpc.status.status();

    if (status.state === "play") {
      player.emit("play", currentSong);
    }

    mpc.on("changed-player", async () => {
      status = await mpc.status.status();
      const currentPlaying = await mpc.status.currentSong();
  
      if (currentSongId! < currentPlaying.id!) {
        player.emit("next", +currentSongId!);
        currentSongId = currentPlaying.id;
      }
  
      if (currentSongId! > currentPlaying.id!) {
        player.emit("previous", +currentSongId!);
        currentSongId = currentPlaying.id;
      }

      if (status.state === "play") {
        player.emit("play", currentPlaying);
      }

      if (status.state === "pause") {
        player.emit("pause", currentPlaying);
      }
    });
  }
})();
