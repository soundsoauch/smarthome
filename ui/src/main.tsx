import * as React from "react";
import { createRoot, Root } from "react-dom/client";

import App from "./app";

class RootComponent extends HTMLElement {
  private _root: Root = createRoot(this);

  connectedCallback() {
    this._root.render(<App />);
  }
}

window.customElements.define("app-root", RootComponent);
