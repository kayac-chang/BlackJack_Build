import { Container, Graphics, Text } from 'pixi.js';

export interface Field extends Container {
  text: string;
}

type IBackground = {
  color: number;
  alpha: number;
  width: number;
  height: number;
  radius: number;
};

type BackgroundProps = {
  color?: number;
  alpha?: number;
  radius?: number;
  padding?: [number, number];
};

export function GameText(text = '', styles = {}) {
  //
  return new Text(text, {
    fontWeight: 'bold',
    fontFamily: 'Arial',
    fill: 0xffffff,
    fontSize: 48,
    ...styles,
  });
}

function Background({ color, alpha, width, height, radius }: IBackground) {
  const it = new Graphics();

  it.beginFill(color, alpha);
  it.drawRoundedRect(0, 0, width, height, radius);
  it.endFill();

  return it;
}

function updateBGByField(it: Container, bgProps: BackgroundProps) {
  let bg: Graphics;

  return function update(field: Container) {
    if (bg) it.removeChild(bg);

    const [paddingW, paddingH] = bgProps.padding || [40, 24];

    bg = Background({
      color: 0x000000,
      alpha: 0.8,
      width: field.width + paddingW,
      height: field.height + paddingH,
      radius: 16,
      ...bgProps,
    });

    bg.pivot.set(bg.width / 2, bg.height / 2);
    it.addChildAt(bg, 0);
  };
}

type Props = {
  fontSize: number;
  background?: BackgroundProps;
};

export function createField({ fontSize, background = {} }: Props): Field {
  const it = new Container();

  const field = GameText('', { fontSize });
  field.anchor.set(0.5);
  it.addChild(field);

  const update = updateBGByField(it, background);

  return Object.defineProperties(it, {
    text: {
      get() {
        return field.text;
      },
      set(text: string | number) {
        it.visible = Boolean(text);

        field.text = String(text);

        update(field);
      },
    },
  });
}
