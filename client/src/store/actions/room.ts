import { Room } from '../../models';
import { RoomAction, ROOM } from '../types';

export function addRoom(payload: Room[]): RoomAction {
  return {
    type: ROOM.ADD,
    payload,
  };
}

export function editRoom(payload: Room): RoomAction {
  return {
    type: ROOM.UPDATE,
    payload,
  };
}
