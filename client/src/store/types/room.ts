import { Room } from '../../models';
import { Action } from 'redux';
import { Payload } from './base';

const PREFIX = '[ROOM]';

export const ROOM = Object.freeze({
  ADD: `${PREFIX} ADD`,
  UPDATE: `${PREFIX} UPDATE`,
});

export type AddRoomAction = Action<typeof ROOM.ADD> & Payload<Room[]>;
export type UpdateRoomAction = Action<typeof ROOM.UPDATE> & Payload<Room>;

export type RoomAction = AddRoomAction | UpdateRoomAction;
