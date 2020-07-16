import { Container, Sprite, Texture } from 'pixi.js';
import RES from '../../assets';
import GameText from '../text';

type Props = {
  texture: Texture;
  msg: string;
  style?: {};
  textPos: { x: number; y: number };
};

function Icon({ texture, msg, style = {}, textPos }: Props) {
  const it = new Container();

  const icon = new Sprite(texture);
  icon.anchor.set(0.5);

  const text = GameText(msg, style);
  text.anchor.set(0.5);
  text.position.set(textPos.x, textPos.y);

  it.addChild(icon, text);

  it.scale.set(0.8);

  return it;
}

export function Win() {
  return Icon({
    texture: RES.get('ICON_WIN').texture,
    msg: 'WIN',
    style: { fill: 0x000000 },
    textPos: { x: 0, y: 76 },
  });
}

export function Lose() {
  return Icon({
    texture: RES.get('ICON_LOSE').texture,
    msg: 'LOSE',
    textPos: { x: 0, y: 65 },
  });
}

export function Bust() {
  return Icon({
    texture: RES.get('ICON_BUST').texture,
    msg: 'BUST',
    textPos: { x: 0, y: 65 },
  });
}
