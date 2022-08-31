import React from "react";
import blessed from "blessed";
import { render } from "react-blessed";
import { Lyric } from "lrc-kit";
import chalk from "chalk";

interface HeaderProps {
  tittle: string;
}

class Header extends React.Component<HeaderProps> {
  render() {
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
        content={chalk.green(this.props.tittle)}
      />
    );
  }
}

interface LyricsBoxProps {
  lyricsText: string;
  currentlyPlaying: string;
}

class LyricsBox extends React.Component<LyricsBoxProps> {
  render() {
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
        content={this.props.lyricsText}
      />
    );
  }
}

type AppProps = HeaderProps &
  Omit<LyricsBoxProps, "lyricsText"> & {
    lyrics: Lyric[] | undefined;
  };

class App extends React.Component<AppProps> {
  render() {
    const { lyrics } = this.props;
    let lyricsText = "";

    if (lyrics) {
      for (const lyric of lyrics) {
        lyricsText += `${lyric.content}\n`;
      }
    } else {
      lyricsText = `No lyrics found. (Can't find LRC file in the directory of ${this.props.currentlyPlaying})`;
    }

    return (
      <element>
        <Header tittle={this.props.tittle} />
        <LyricsBox
          lyricsText={lyricsText}
          currentlyPlaying={this.props.currentlyPlaying}
        />
      </element>
    );
  }
}

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
