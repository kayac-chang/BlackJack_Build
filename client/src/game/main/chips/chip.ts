import { CHIP } from '../../../models';
import RES from '../../assets';
import { Sprite } from 'pixi.js';

function mapping(type: CHIP) {
  switch (type) {
    case CHIP.RED:
      return 'CHIP_RED';
    case CHIP.GREEN:
      return 'CHIP_GREEN';
    case CHIP.BLUE:
      return 'CHIP_BLUE';
    case CHIP.BLACK:
      return 'CHIP_BLACK';
    case CHIP.PURPLE:
      return 'CHIP_PURPLE';
    case CHIP.YELLOW:
      return 'CHIP_YELLOW';
  }
}

export default function Chip(type: CHIP) {
  //
  const chip = new Sprite(RES.get(mapping(type)).texture);

  chip.anchor.set(0.5);
  chip.scale.set(0.6);

  return chip;
}
