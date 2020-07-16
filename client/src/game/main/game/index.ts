import { Container } from 'pixi.js';
import { SEAT } from '../../../models';
import { createHandService } from './state';

export default function Game() {
  const container = new Container();
  container.name = 'game';

  const pokers = new Container();
  pokers.name = 'pokers';
  container.addChild(pokers);

  for (const id in SEAT) {
    if (isNaN(Number(id))) {
      continue;
    }

    const seatID = Number(id) as SEAT;
    createHandService(seatID, container).start();
  }

  return container;
}
