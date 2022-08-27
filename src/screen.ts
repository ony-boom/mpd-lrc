import blessed from "blessed";

export const setBoxes = (screen: blessed.Widgets.Screen) => {
  const headerBox = blessed.box({
    top: "top",
    left: "center",
    width: "100%",
    height: "8%",
    border: "line",
    padding: {
      left: 2,
      right: 2,
    },
    tags: true,
    style: {
      fg: "green",
    },
    screen,
  });
  
  const lyricsBox = blessed.box({
    top: "center",
    left: "center",
    width: "100%",
    height: "92%",
    scrollable: true,
    tags: true,
    border: {
      type: "line",
    },
    padding: {
      left: 2,
      right: 2,
    },
    keys: true,
    vi: true,
    alwaysScroll: true,
    scrollbar: {
      style: {
        bg: "yellow",
      },
    },
    screen,
  });
  
  return { headerBox, lyricsBox };
};

export const setScreen = () => {
  const screen = blessed.screen({
    smartCSR: true,
    dockBorders: true,
    fullUnicode: true,
    fastCSR: true
  });
  
  screen.key(["escape", "q", "C-c"], function () {
    return process.exit(0);
  });
  return screen;
};
