import { Loader } from 'pixi.js';

import TABLE from './table';
import SEAT from './seat';
import POKER from './poker';
import CHIP from './chip';
import ICON from './icon';

const PKG = Object.freeze({
  ...CHIP,
  ...TABLE,
  ...SEAT,
  ...POKER,
  ...ICON,
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

function get(res: keyof typeof PKG) {
  const resource = loader.resources[res];

  if (!resource) {
    throw new Error(`Can not found resource [${res}]`);
  }

  return resource;
}

export default {
  load,
  get,
  PKG,
};
