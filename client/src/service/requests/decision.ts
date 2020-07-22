import { EVENT } from '../types';
import Service from '../service';
import { C2S, DECISION } from '../../models';
import store from '../../store';

export default async function (service: Service, decision: DECISION) {
  const { game, user, seat } = store.getState();

  if (!game.turn) {
    throw new Error(`No turn existed...`);
  }

  if (seat[game.turn.seat].player !== user.name) {
    throw new Error(`Not user's turn ...`);
  }

  service.send({
    cmd: C2S.CLIENT.DECISION,
    data: {
      id: game.room,
      no: game.turn.seat,

      action: {
        [DECISION.DOUBLE]: 'dbl',
        [DECISION.SURRENDER]: 'gvp',
        [DECISION.HIT]: 'hit',
        [DECISION.INSURANCE]: 'ins',
        [DECISION.PAY]: 'pay',
        [DECISION.SPLIT]: 'spt',
        [DECISION.STAND]: 'sty',
      }[decision],

      pile: {
        L: 0,
        R: 1,
      }[game.turn.pair],
    },
  });

  return new Promise((resolve) => service.once(EVENT.DECISION, resolve));
}
