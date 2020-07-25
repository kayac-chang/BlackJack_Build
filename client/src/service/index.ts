import Service from './service';
import {
  joinRoom as _joinRoom,
  joinSeat as _joinSeat,
  leaveSeat as _leaveSeat,
  leaveRoom as _leaveRoom,
  deal as _deal,
  decision as _decision,
} from './requests';
import { Token, SEAT, DECISION } from '../models';

const service = new Service(process.env.REACT_APP_BACKEND || '');

function init(token?: Token) {
  return service.connect(token);
}

function joinRoom(roomID: number) {
  return _joinRoom(service, roomID);
}

function joinSeat(seat: SEAT) {
  return _joinSeat(service, seat);
}

function leaveSeat(seat: SEAT) {
  return _leaveSeat(service, seat);
}

function leaveRoom() {
  return _leaveRoom(service);
}

function deal() {
  return _deal(service);
}

function decision(decision: DECISION) {
  return _decision(service, decision);
}

export default { init, joinRoom, joinSeat, leaveSeat, leaveRoom, deal, decision };
