import { Sprite, Container } from 'pixi.js';
import RES from '../../assets';
import GameText from '../text';

export interface Field extends Container {
  text: string;
}

export function createField(): Field {
  const it = new Container();

  const background = new Sprite(RES.get('FIELD').texture);
  background.anchor.set(0.5);
  background.alpha = 0.8;
  background.tint = 0x000000;
  it.addChild(background);

  const field = GameText('');
  field.anchor.set(0.5);
  it.addChild(field);

  return Object.defineProperties(it, {
    text: {
      set: (value) => (field.text = value),
    },
  });
}
