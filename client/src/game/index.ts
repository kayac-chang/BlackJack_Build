import Game from './core/game';
import Main from './main';

export default async function (view: HTMLCanvasElement) {
  const app = Game(view);

  app.stage.addChild(Main(app));
}
