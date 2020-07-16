import { Graphics, Point } from 'pixi.js';

type Line = {
  start: Point;
  end: Point;
  color?: number;
  width?: number;
};
export function line({ start, end, color = 0xffffff, width = 5 }: Line, it: Graphics) {
  return (
    it
      //
      .lineStyle(width, color)
      .moveTo(start.x || 0, start.y || 0)
      .lineTo(end.x || 0, end.y || 0)
  );
}

export function circle(center: Point, it: Graphics) {
  const color = 0xffffff;
  const radius = 5;

  return (
    it
      //
      .beginFill(color)
      .drawCircle(center.x || 0, center.y || 0, radius)
      .endFill()
  );
}

export function bezierCurve(start: Point, controlA: Point, controlB: Point, end: Point, it: Graphics) {
  const color = 0xffffff;
  const width = 5;

  return (
    it
      //
      .lineStyle(width, color)
      .moveTo(start.x || 0, start.y || 0)
      .bezierCurveTo(
        //
        controlA.x || 0,
        controlA.y || 0,
        //
        controlB.x || 0,
        controlB.y || 0,
        //
        end.x || 0,
        end.y || 0
      )
  );
}
