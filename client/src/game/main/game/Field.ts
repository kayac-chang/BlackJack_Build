import { Container, Graphics, Text } from "pixi.js";
import GameText from "../text";

export default class Field extends Container {
  field: Text;

  constructor(text: string | number = "") {
    super();

    const [width, height] = [108, 76];

    const background = new Graphics();
    background.beginFill(0x000, 0.7);
    background.drawRoundedRect(-0.5 * width, -0.5 * height, width, height, 16);
    background.endFill();
    this.addChild(background);

    const field = GameText();
    field.anchor.set(0.5);
    this.addChild(field);
    this.field = field;

    this.text = text;
  }

  get text() {
    return this.field.text;
  }
  set text(text: string | number) {
    this.visible = Boolean(text);

    this.field.text = String(text);
  }
}
