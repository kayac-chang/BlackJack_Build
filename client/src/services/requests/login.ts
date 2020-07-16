import { EVENT } from '../types';
import Service from '../service';
import { C2S, User } from '../../models';

export default async function (service: Service): Promise<User> {
  console.groupCollapsed('Login');

  service.send({
    cmd: C2S.CLIENT.LOGIN,
    data: undefined,
  });

  const user = await new Promise<User>((resolve) => service.once(EVENT.UPDATE_USER, resolve));

  console.log(user);

  console.groupEnd();

  return user;
}
