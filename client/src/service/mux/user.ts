import { S2C, PAIR, SEAT } from '../../models';
import Service from '../service';
import { EVENT, LoginProp, UpdateProp, ActionProp, toSeatNum } from '../types';
import store from '../../store';
import { login, update, commitBet, updateSeats, updateHand } from '../../store/actions';

function onLogin(service: Service, { user_name }: LoginProp) {
  const { user } = store.getState();

  store.dispatch(
    login({
      ...user,
      name: String(user_name),
      balance: 0,
      totalBet: 0,
      decisions: [],
    })
  );
}

function onUpdate(service: Service, { name, balance }: UpdateProp) {
  const { user } = store.getState();

  const res = store.dispatch(
    update({
      ...user,
      name: String(name),
      balance: Number(balance),
    })
  );

  service.emit(EVENT.UPDATE_USER, res.payload);
}

function onBet(service: Service, data: any) {
  const { bet } = store.getState();

  store.dispatch(commitBet(bet.history));

  return service.emit(EVENT.BET);
}

function onSplit(seatID: SEAT) {
  const { seat, hand } = store.getState();

  const [L, R] = hand[seatID];

  store.dispatch(
    updateHand({
      ...hand,
      [seatID]: [L, { ...R, pair: PAIR.R }],
    })
  );

  store.dispatch(updateSeats({ ...seat, [seatID]: { ...seat[seatID], split: true } }));
}

function onAction(service: Service, data?: ActionProp) {
  if (!data) {
    return;
  }

  const { user } = store.getState();

  store.dispatch(
    update({
      ...user,
      decisions: [],
    })
  );

  if (data.action === 'spt') {
    onSplit(toSeatNum(data.no));
  }

  return service.emit(EVENT.DECISION, data.action);
}

export default {
  [S2C.USER.LOGIN]: onLogin,
  [S2C.USER.UPDATE]: onUpdate,
  [S2C.USER.BET]: onBet,
  [S2C.USER.ACTION]: onAction,
};
