import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";

import Game from "./game";
import { App, UI } from "./ui";
import { i18n, gsap } from "./plugins";
import service from "./services";
import "./index.scss";

import store from "./store";
import { getURLParam } from "./utils";

import RES from "./assets";

async function main() {
  await Promise.all([
    i18n.init(),
    gsap.init(),
    service.init(getURLParam("token")),
    RES.load(),
  ]);

  const Root = (
    <React.StrictMode>
      <Provider store={store}>
        <App game={Game} ui={<UI />} />
      </Provider>
    </React.StrictMode>
  );
  ReactDOM.render(Root, document.getElementById("root"));
}

main();
