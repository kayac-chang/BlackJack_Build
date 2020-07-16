import { Container } from "pixi.js";
import { SEAT, Turn } from "../../../models";
import { createSeat, Seat } from "./seat";
import { observe } from "../../../store";
import { Effect } from "./effect";

type Props = {
  id: SEAT;
  x: number;
  y: number;
};

function updateEffect(seats: Seat[]) {
  const effect = Effect();
  effect.visible = false;
  effect.scale.set(0.75);

  return function (turn?: Turn) {
    if (!turn) {
      effect.visible = false;
      return;
    }

    const found = seats.find(({ name }) => name === SEAT[turn.seat]);
    if (!found) {
      effect.visible = false;
      return;
    }

    effect.visible = true;
    found.addChild(effect);
  };
}

export default function Seats(meta: Props[]) {
  const it = new Container();
  it.name = "seats";

  it.once("added", function onInit({ width, height }: Container) {
    //
    const seats = meta.map(({ id, x, y }) =>
      createSeat({
        id,
        x: width * x,
        y: height * y,
      })
    );

    it.addChild(...seats);

    observe((state) => state.game.turn, updateEffect(seats));
  });

  return it;
}
