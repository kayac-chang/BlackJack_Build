import Service from '../service';
import { C2S, SEAT } from '../../models';
import store from '../../store';
import { EVENT } from '../types';

export default async function (service: Service, seat: SEAT) {
  const { game } = store.getState();

  service.send({
    cmd: C2S.CLIENT.LEAVE_SEAT,
    data: {
      id: game.room,
      no: seat,
    },
  });

  return new Promise((resolve) => {
    service.once(EVENT.UPDATE_SEAT, resolve);
  });
}
