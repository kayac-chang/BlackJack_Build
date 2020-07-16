import { Sprite, AnimatedSprite, BitmapText, Container, Text, Point, Graphics } from 'pixi.js';

export function isContainer(element: any): element is Container {
  return element instanceof Container;
}

export function isSprite(element: any): element is Sprite {
  return element instanceof Sprite;
}

export function isAnimatedSprite(element: any): element is AnimatedSprite {
  return element instanceof AnimatedSprite;
}

export function isBitmapText(element: any): element is BitmapText {
  return element instanceof BitmapText;
}

export function isText(element: any): element is Text {
  return element instanceof Text;
}

export function isPoint(element: any): element is Point {
  return element instanceof Point;
}

export function isGraphics(element: any): element is Graphics {
  return element instanceof Graphics;
}
