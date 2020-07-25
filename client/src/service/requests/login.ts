import { EVENT } from '../types';
import Service from '../service';
import { C2S, User } from '../../models';

function clearURLParam() {
  return window.history.pushState(undefined, '', window.location.origin + window.location.pathname);
}

export default async function (service: Service): Promise<User> {
  service.send({
    cmd: C2S.CLIENT.LOGIN,
    data: undefined,
  });

  const user = await new Promise<User>((resolve) => service.once(EVENT.UPDATE_USER, resolve));

  clearURLParam();

  return user;
}
