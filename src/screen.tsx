import React from "react";
import blessed from "blessed";
import { render } from "react-blessed";
import { Lyric } from "lrc-kit";
import chalk from "chalk";
import { mpc } from "./index";

interface HeaderProps {
  tittle: string;
}

const Header: React.FC<HeaderProps> = ({ tittle }) => {
  return (
    <box
      left="center"
      width="100%"
      height="12%"
      border={{ type: "line" }}
      padding={{
        left: 1,
        right: 1,
        top: 0,
        bottom: 0,
      }}
      content={chalk.green(tittle)}
    />
  );
};

interface LyricsBoxProps {
  lyricsText: string;
  currentlyPlaying: string;
}

const LyricsBox: React.FC<LyricsBoxProps> = ({ lyricsText }) => {
  return (
    <box
      top="center"
      left="center"
      width="100%"
      height="92%"
      border="line"
      padding={{
        left: 1,
        right: 1,
        top: 0,
        bottom: 0,
      }}
      scrollable={true}
      alwaysScroll={true}
      mouse={true}
      keys={true}
      vi={true}
      tags={true}
      content={lyricsText}
    />
  );
};

type AppProps = HeaderProps &
  Omit<LyricsBoxProps, "lyricsText"> & {
    lyrics: Lyric[] | undefined;
  };


let oldInterval: NodeJS.Timer | null;
const App: React.FC<AppProps> = ({ lyrics, tittle, currentlyPlaying }) => {
  let lyricsText = "";
  const [elapsedTime, setElapsedTime] = React.useState(0);

  React.useEffect(() => {
    oldInterval && clearInterval(oldInterval);
    const interval = setInterval(async () => {
      const status = await mpc.status.status();
      if (status.elapsed) setElapsedTime(status.elapsed);
    }, 100);
    oldInterval = interval;
    return () => clearInterval(interval);
  }, []);

  if (lyrics) {
    for (let i = 0; i < lyrics.length; i++) {
      const currentLyric = lyrics[i];
      let nextLyric = lyrics[i + 1];

      if (!nextLyric) {
        nextLyric = {
          ...currentLyric,
          timestamp: currentLyric.timestamp + 100,
        };
      }

      if (
        Math.ceil(elapsedTime) < Math.ceil(nextLyric.timestamp) &&
        Math.ceil(elapsedTime) >= Math.ceil(currentLyric.timestamp)
      ) {
        lyricsText += `${chalk.bgYellow(
          chalk.black(`${currentLyric.content}\n{/}`)
        )}`;
      } else {
        lyricsText += `${currentLyric.content}\n`;
      }
    }
  } else {
    lyricsText = `No lyrics found. (Can't find LRC file in the directory of ${currentlyPlaying})`;
   if (oldInterval) {
     clearInterval(oldInterval);
     oldInterval = null;
   }
  }

  return (
    <element>
      <Header tittle={tittle} />
      <LyricsBox lyricsText={lyricsText} currentlyPlaying={currentlyPlaying} />
    </element>
  );
};

const renderScreen = (screenTittle: string, lyrics: Lyric[] | undefined) => {
  const screen = blessed.screen({
    smartCSR: true,
    title: "lrc",
    dockBorders: true,
    autoPadding: true,
    fullUnicode: true,
    fastCSR: true,
  });

  screen.key(["escape", "q", "C-c"], function () {
    return process.exit(0);
  });

  render(
    <App
      tittle={screenTittle}
      lyrics={lyrics}
      currentlyPlaying={screenTittle}
    />,
    screen
  );

  return screen;
};

export default renderScreen;
