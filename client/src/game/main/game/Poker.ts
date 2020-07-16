import Res from '../../assets';
import POKER from '../../assets/poker';
import { Sprite, SimpleMesh, Container } from 'pixi.js';
import gsap from 'gsap';
import { Power0 } from 'gsap/gsap-core';
import { SUIT, RANK } from '../../../models';
import { nextFrame } from '../../../utils';

export default class Poker extends Container {
  //
  duration = 0.5;

  constructor(suit: SUIT, rank: RANK) {
    super();

    const back = new Face('BACK');
    const front = new Face(`${suit}_${rank}` as keyof typeof POKER);
    this.addChild(back, front);
  }

  _faceUp = true;
  get faceUp() {
    return this._faceUp;
  }
  set faceUp(flag: boolean) {
    const [front, back] = this.children;
    this.swapChildren(back, front);
    this._faceUp = flag;
  }

  private fliping = false;
  async flip() {
    if (this.fliping) return;

    this.fliping = true;

    const [front, back] = this.children as SimpleMesh[];

    await nextFrame();
    await flip(back, front, this.duration);

    this.faceUp = !this.faceUp;
    this.fliping = false;
  }
}

class Face extends SimpleMesh {
  //
  constructor(res: keyof typeof POKER) {
    super();

    const sprite = new Sprite(Res.get(res).texture);
    this.addChild(sprite);
    sprite.anchor.set(0.5);

    requestAnimationFrame(() => {
      this.vertices = (sprite as any)['vertexData'];
    });
  }
}

function flip(back: SimpleMesh, front: SimpleMesh, duration: number) {
  //
  const origin = Array.from(back.vertices);
  const [x1, y1, x2, y2, , y3, ,] = origin;

  const width = x2 - x1;
  const height = y3 - y2;

  // points
  const a = [x1 + width * 0.5, y1 + height * 0.1];
  const b = [x1 + width * 0.5, y1 + height * -0.1];
  const c = [x1 + width * 0.5, y1 + height * 1];
  const d = [x1 + width * 0.5, y1 + height * 0.9];

  Object.assign(front.vertices, [...b, ...a, ...d, ...c]);

  return gsap
    .timeline()
    .add(tween(back, [...a, ...b, ...c, ...d], duration / 2))
    .add(tween(front, origin, duration / 2))
    .then();

  function tween(card: SimpleMesh, vertices: number[], duration: number) {
    //
    const state = vertices as any;

    return gsap.to(card.vertices, {
      ...state,
      duration,
      ease: Power0.easeIn,
    });
  }
}
