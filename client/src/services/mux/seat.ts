import { S2C } from '../../models';
import Service from '../service';

import store from '../../store';
import { updateSeats } from '../../store/actions';
import { EVENT, SeatProp, toSeats, toSeatNum } from '../types';

function onUpdate(service: Service, data: SeatProp[]) {
  const { user, seat } = store.getState();

  data = data.map(({ no, player, total_bet, ...props }) => {
    if (player === user.name) {
      total_bet = seat[toSeatNum(no)].bet;
    }

    return { no, player, total_bet, ...props };
  });

  const action = store.dispatch(updateSeats(toSeats(data)));

  service.emit(EVENT.UPDATE_SEAT, action.payload);
}

export default {
  [S2C.SEAT.UPDATE]: onUpdate,
};
