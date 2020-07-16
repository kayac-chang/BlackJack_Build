import { Container, Sprite } from 'pixi.js';
import { observe } from '../../../store';
import { Bet, SEAT, Seats } from '../../../models';
import { without } from 'ramda';
import Chip from './chip';
import gsap from 'gsap';

type Props = {
  id: SEAT;
  x: number;
  y: number;
};

function Group(id: SEAT, x: number, y: number) {
  //
  return Object.assign(new Container(), { name: SEAT[id], x, y });
}

function transIn(chip: Sprite) {
  chip.y -= 50;

  gsap.to(chip, { y: 0, alpha: 1 });
}

function findGroupBySeatID(groups: Container[], seat: SEAT) {
  //
  return groups.find(({ name }) => name === SEAT[seat]);
}

function updateChip(groups: Container[]) {
  let pre: Bet[] = [];

  function addChip({ seat, chip }: Bet) {
    //
    if (seat === undefined) {
      return;
    }

    const group = findGroupBySeatID(groups, seat);
    if (!group) {
      return;
    }

    const _chip = Chip(chip);
    group.addChild(_chip);
    transIn(_chip);
  }

  return function onUpdate(history: Bet[]) {
    //
    if (history.length > pre.length) {
      without(pre, history).forEach(addChip);
    }

    pre = [...history];
  };
}

function updateByBet(seats: Container[]) {
  return function (state: Seats) {
    //
    for (const [id, seat] of Object.entries(state)) {
      if (seat.bet > 0) continue;

      const found = findGroupBySeatID(seats, Number(id) as SEAT);

      if (found) {
        found.removeChildren();
      }
    }
  };
}

function init(container: Container, meta: Props[]) {
  //
  return function onInit({ width, height }: Container) {
    //
    const seats = meta.map(({ id, x, y }) => Group(id, width * x, height * y));

    container.addChild(...seats);

    observe((state) => state.bet.history, updateChip(seats));
    observe((state) => state.seat, updateByBet(seats));
  };
}

export default function Chips(meta: Props[]) {
  const chips = new Container();
  chips.name = 'chips';
  chips.once('added', init(chips, meta));

  return chips;
}
