import { EVENT } from '../types';
import Service from '../service';
import { C2S } from '../../models';
import store from '../../store';

export default async function (service: Service) {
  const { game, user, seat } = store.getState();

  const bets = [];

  for (const [id, _seat] of Object.entries(seat)) {
    if (_seat.player === user.name) {
      bets.push({ no: Number(id), bet: _seat.bet });
    }
  }

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
