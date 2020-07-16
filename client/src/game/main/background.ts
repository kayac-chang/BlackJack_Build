import { Container, Sprite } from 'pixi.js';
import RES from '../assets';

export default function Background() {
  const it = new Container();
  it.name = 'background';

  const table = new Sprite(RES.get('TABLE_BLUE').texture);
  table.name = 'table';

  const title = new Sprite(RES.get('TABLE_TITLE').texture);
  title.name = 'title';

  title.x = table.width / 2;
  title.y = table.height / 3;
  title.anchor.set(0.5);

  it.addChild(table, title);

  return it;
}
