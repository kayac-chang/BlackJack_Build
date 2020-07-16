import { Sprite, Container } from 'pixi.js';
import { SEAT } from '../../../models';
import { createField, Field } from './field';
import { observe } from '../../../store';
import { createSeatService, SeatService, SeatState } from './state';
import { pipe } from 'ramda';
import RES from '../../assets';

interface Prop {
  id: SEAT;
  x: number;
  y: number;
}

function updateSeat(seat: Sprite) {
  return function onChange(state: SeatState) {
    if (state.matches('empty')) {
      seat.texture = RES.get('SELECT_SEAT_NORMAL').texture;
    }

    if (state.matches({ occupy: 'normal' })) {
      seat.texture = RES.get('SEAT_NORMAL').texture;
    }

    if (state.matches({ occupy: 'betting' })) {
      seat.texture = RES.get('SEAT_ENABLE').texture;
    }

    return state;
  };
}

function updateField(field: Field) {
  return function onChange(state: SeatState) {
    if (state.matches('empty')) {
      field.visible = false;
    }

    if (state.matches('occupy')) {
      field.visible = true;

      field.text = state.context.owner;
    }

    return state;
  };
}

export interface Seat extends Container {
  service: SeatService;
}

export function createSeat({ id, x, y }: Prop): Seat {
  const it = new Container();
  it.name = SEAT[id];
  it.buttonMode = true;
  it.interactive = true;
  it.x = x;
  it.y = y;

  const sprite = new Sprite();
  sprite.anchor.set(0.5);
  sprite.scale.set(0.75);
  it.addChild(sprite);

  const field = createField();
  field.y = 130;
  it.addChild(field);

  const service = createSeatService(id);
  service.onTransition(
    pipe(
      //
      updateSeat(sprite),
      updateField(field)
    )
  );

  it.on('pointerdown', service.send);
  it.on('added', () => service.start());
  it.on('removed', () => service.stop());

  observe(
    (state) => state.seat[id].player,
    (user) => service.send({ type: user ? 'JOIN' : 'LEAVE', user })
  );
  observe(
    (state) => state.game.state,
    (state) => service.send({ type: 'STATE' })
  );

  return Object.assign(it, { service });
}
