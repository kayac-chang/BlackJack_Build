import { Container, Application } from 'pixi.js';
import Background from './background';
import Seats from './seats';
import Game from './game';
import Chips from './chips';
import { SEAT } from '../../models';
import Bets from './bets';

const seatMeta = [
  { id: SEAT.A, x: 15 / 100, y: 58 / 100 },
  { id: SEAT.B, x: 30 / 100, y: 75 / 100 },
  { id: SEAT.C, x: 50 / 100, y: 82 / 100 },
  { id: SEAT.D, x: 70 / 100, y: 75 / 100 },
  { id: SEAT.E, x: 85 / 100, y: 58 / 100 },
];

export default function Scene(app: Application): Container {
  const scene = new Container();
  scene.name = 'main';

  const background = Background();
  scene.addChild(background);

  const game = Game();
  scene.addChild(game);

  const seats = Seats(seatMeta);
  scene.addChild(seats);

  const chips = Chips(seatMeta);
  scene.addChild(chips);

  const bets = Bets(seatMeta);
  scene.addChild(bets);

  return scene;
}
