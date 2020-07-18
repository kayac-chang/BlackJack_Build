import { Loader } from "pixi.js";

import TABLE from "./table";
import SEAT from "./seat";
import POKER from "./poker";
import CHIP from "./chip";
import ICON from "./icon";
import LOBBY from "./lobby";

import { toBase64 } from "../utils";

const PKG = Object.freeze({
  ...CHIP,
  ...TABLE,
  ...SEAT,
  ...POKER,
  ...ICON,
  ...LOBBY,
});

const loader = new Loader();

async function load() {
  for (const [name, url] of Object.entries(PKG)) {
    if (loader.resources[name]) {
      continue;
    }

    loader.add(name, url);
  }

  return new Promise((resolve, reject) => {
    loader.load(resolve);

    loader.onError.add(reject);
  });
}

const cache: Record<string, string> = {};

function getBase64(res: keyof typeof PKG) {
  const resource = loader.resources[res];

  if (!cache[res]) {
    cache[res] = toBase64(resource.data);
  }

  return cache[res];
}

function getTexture(res: keyof typeof PKG) {
  const resource = loader.resources[res];

  return resource.texture;
}

export default {
  load,
  getBase64,
  getTexture,
  PKG,
};
