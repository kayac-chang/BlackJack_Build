import { Seats } from '../../models';
import { Action } from 'redux';
import { Payload } from './base';

const PREFIX = `[SEAT]`;

export const SEAT = Object.freeze({
  UPDATE: `${PREFIX} UPDATE`,
  CLEAR: `${PREFIX} CLEAR`,
});

type UpdateSeatAction = Action<typeof SEAT.UPDATE> & Payload<Seats>;
type ClearSeatAction = Action<typeof SEAT.CLEAR> & Payload<undefined>;

export type SeatAction = UpdateSeatAction | ClearSeatAction;
