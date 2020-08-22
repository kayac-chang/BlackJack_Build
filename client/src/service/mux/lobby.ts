import { S2C } from '../../models';
import Service from '../service';

import store from '../../store';
import { addRoom, editRoom, join, updateSeats } from '../../store/actions';
import { GameProp, toGame, toSeats, EVENT, RoomProp, toRoom } from '../types';

function onLobby (service: Service, data: RoomProp[]) {
  const rooms = data.map(toRoom);

  store.dispatch(addRoom(rooms));
}

function onUpdate (service: Service, data: RoomProp) {
  const room = toRoom(data);

  store.dispatch(editRoom(room));
}

function onJoin (service: Service, data: GameProp) {
  const { room } = store.getState();

  const action = store.dispatch(join(toGame(data)));

  store.dispatch(updateSeats(toSeats(data.seats)));

  service.emit(EVENT.JOIN_ROOM, action.payload);

  const found = room.find(({ id }) => data.id === id);
  if (!found) {
    return;
  }

  store.dispatch(
    editRoom({
      ...found,
      history: data.history.map(String),
    })
  );
}

export default {
  [S2C.ROOM.ADD]: onLobby,
  [S2C.ROOM.UPDATE]: onUpdate,
  [S2C.ROOM.JOIN]: onJoin,
};
