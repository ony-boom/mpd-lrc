import React from "react";
import blessed from "blessed";
import { render } from "react-blessed";
import chalk from "chalk";
import { Lyric } from "lrc-kit";

interface HeaderProps {
  tittle: string;
}

class Header extends React.Component<HeaderProps> {
  render() {
    return (
      <blessed-box
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
      >
        {chalk.green(this.props.tittle)}
      </blessed-box>
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
      >
        {this.props.lyricsText}
      </box>
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
  });

  screen.key(["escape", "q", "C-c"], function () {
    return process.exit(0);
  });

  screen.title = screenTittle;
  return render(
    <App
      tittle={screenTittle}
      lyrics={lyrics}
      currentlyPlaying={screenTittle}
    />,
    screen
  );
};

export default renderScreen;
