import { UserAction, USER } from '../types';
import { User } from '../../models';

export function login(user: User): UserAction {
  return {
    type: USER.LOGIN,
    payload: user,
  };
}

export function update(user: User): UserAction {
  return {
    type: USER.UPDATE,
    payload: user,
  };
}
