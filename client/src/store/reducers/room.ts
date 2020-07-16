import { Room } from '../../models';
import { sort, ascend, prop, map } from 'ramda';
import { RoomAction, ROOM } from '../types';

const sortRoomAscByID = sort<Room>(ascend(prop('id')));

function update(newRoom: Room, rooms: Room[]) {
  return map((room) => (newRoom.id === room.id ? newRoom : room), rooms);
}

const initialState: Room[] = [];

export default function roomReducer(state = initialState, action: RoomAction): Room[] {
  const { type, payload } = action;

  if (type === ROOM.ADD) {
    const rooms = payload as Room[];

    return sortRoomAscByID(rooms);
  }

  if (type === ROOM.UPDATE) {
    const newRoom = payload as Room;

    return update(newRoom, state);
  }

  return state;
}
