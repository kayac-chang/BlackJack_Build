import { Application } from 'pixi.js';
import { getSize, resize } from './screen';

import * as PIXI from 'pixi.js';
window.PIXI = PIXI;

let app: Application | undefined = undefined;

export default function (view: HTMLCanvasElement) {
  //
  if (app) {
    app.destroy();
  }

  app = new Application({
    view,
    ...getSize(),
    resolution: 1,
  });

  app.ticker.add(resize(app));

  return app;
}
