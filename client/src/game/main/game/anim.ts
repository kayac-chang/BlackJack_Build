import { DisplayObject } from 'pixi.js';
import gsap from 'gsap';

export function tween<T extends DisplayObject>(element: T, options: gsap.TweenVars = {}) {
  return gsap.to(element, { duration: 0.8, ease: 'expo.out', ...options });
}

export function move(pokers: DisplayObject[], end: { x: number; y: number }) {
  const offset = 50;
  const mid = pokers.length === 1 ? 0 : Math.round(pokers.length / 2) - 0.5;

  return pokers.map((child, index) => tween(child, { x: end.x + offset * (index - mid), y: end.y, alpha: 1 }));
}
