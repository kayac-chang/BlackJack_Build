import { User } from '../../models';
import { Action } from 'redux';
import { Payload } from './base';

const PREFIX = '[USER]';

export const USER = Object.freeze({
  LOGIN: `${PREFIX} LOGIN`,
  UPDATE: `${PREFIX} UPDATE`,
});

export type LoginAction = Action<typeof USER.LOGIN> & Payload<User>;
export type UpdateAction = Action<typeof USER.UPDATE> & Payload<User>;

export type UserAction = LoginAction | UpdateAction;
