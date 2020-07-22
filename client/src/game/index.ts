import Res from "../assets";
import Game from "./core/game";
import Main from "./main";

export default async function (view: HTMLCanvasElement) {
  const app = Game(view);

  await Res.load();

  app.stage.addChild(Main(app));
}
