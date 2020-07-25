import { Container, Text } from 'pixi.js';
import { SEAT, Seats } from '../../../models';
import { observe } from '../../../store';
import { createField } from '../field';

type Props = {
  id: SEAT;
  x: number;
  y: number;
};

function findGroupBySeatID(groups: Container[], seat: SEAT) {
  return groups.find(({ name }) => name === SEAT[seat]);
}

function updateSeat(groups: Container[]) {
  //
  function bet(group: Container, totalBet: number) {
    const field = group.getChildByName('field') as Text;

    field.text = totalBet ? String(totalBet) : '';
  }

  function show(group: Container, totalBet: number) {
    group.visible = totalBet > 0;
  }

  return function onUpdate(seats: Seats) {
    //
    for (const [id, seat] of Object.entries(seats)) {
      const group = findGroupBySeatID(groups, Number(id) as SEAT);

      if (!group) {
        continue;
      }

      bet(group, seat.bet);
      show(group, seat.bet);
    }
  };
}

function Bet(id: SEAT, x: number, y: number) {
  const it = new Container();
  it.name = SEAT[id];
  it.position.set(x, y + 70);
  it.visible = false;

  const field = createField({ fontSize: 48, background: { color: 0xf0aa0a, alpha: 1, padding: [24, 4], radius: 24 } });
  field.name = 'field';
  it.addChild(field);

  return it;
}

function init(container: Container, meta: Props[]) {
  //
  return function onInit({ width, height }: Container) {
    //
    const seats = meta.map(({ id, x, y }) => Bet(id, width * x, height * y));

    container.addChild(...seats);

    observe((state) => state.seat, updateSeat(seats));
  };
}

export default function (meta: Props[]) {
  const it = new Container();
  it.name = 'bets';
  it.once('added', init(it, meta));

  return it;
}
