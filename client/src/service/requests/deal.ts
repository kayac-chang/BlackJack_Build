import { EVENT } from '../types';
import Service from '../service';
import { C2S, SEAT as TSEAT } from '../../models';
import store from '../../store';
import { updateSeats } from '../../store/actions';

export default async function (service: Service) {
  const { game, user, seat } = store.getState();

  const bets = [];
  const _seat = { ...seat };

  for (const [id, s] of Object.entries(seat)) {
    if (s.player === user.name) {
      bets.push({ no: Number(id), bet: s.bet });

      _seat[Number(id) as TSEAT].commited = true;
    }
  }

  store.dispatch(updateSeats(_seat));

  service.send({
    cmd: C2S.CLIENT.BET,
    data: {
      id: game.room,
      bets,
    },
  });

  return new Promise((resolve) => {
    service.once(EVENT.BET, resolve);
  });
}
