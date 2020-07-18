import { CHIP } from "../../../models";
import RES from "../../../assets";
import { Sprite } from "pixi.js";

function mapping(type: CHIP) {
  switch (type) {
    case CHIP.RED:
      return "CHIP_FLAT_RED";
    case CHIP.GREEN:
      return "CHIP_FLAT_GREEN";
    case CHIP.BLUE:
      return "CHIP_FLAT_BLUE";
    case CHIP.BLACK:
      return "CHIP_FLAT_BLACK";
    case CHIP.PURPLE:
      return "CHIP_FLAT_PURPLE";
    case CHIP.YELLOW:
      return "CHIP_FLAT_YELLOW";
  }
}

export default function Chip(type: CHIP) {
  //
  const chip = new Sprite(RES.getTexture(mapping(type)));

  chip.anchor.set(0.5);
  chip.scale.set(0.6);

  return chip;
}
