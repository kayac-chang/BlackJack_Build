import { EVENT } from '../types';
import Service from '../service';
import { C2S } from '../../models';

export default async function (service: Service, id: number) {
  service.send({
    cmd: C2S.CLIENT.JOIN_ROOM,
    data: { id },
  });

  return new Promise((resolve) => {
    service.once(EVENT.JOIN_ROOM, resolve);
  });
}
