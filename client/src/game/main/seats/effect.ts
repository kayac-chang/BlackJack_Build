import { Sprite, Container } from "pixi.js";
import RES from "../../../assets";
import gsap from "gsap";

export function Effect() {
  const it = new Container();

  const effect = new Sprite(RES.getTexture("SEAT_NORMAL"));

  effect.tint = 0xf0aa0a;
  effect.anchor.set(0.5);

  gsap.fromTo(
    //
    effect,
    { pixi: { alpha: 1, scale: 1 } },
    { pixi: { alpha: 0, scale: 1.5 }, duration: 1, repeat: -1 }
  );

  it.addChild(effect);

  return it;
}
