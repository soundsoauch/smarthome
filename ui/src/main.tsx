import * as React from "react";
import * as ReactDOM from "react-dom";

import { App } from "./app";

class RootComponent extends HTMLElement {
  connectedCallback() {
    ReactDOM.render(<App />, this);
  }
}

window.customElements.define("app-root", RootComponent);
