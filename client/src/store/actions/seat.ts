import { Seats } from '../../models';
import { SeatAction, SEAT } from '../types';

export function updateSeats(payload: Seats): SeatAction {
  return { type: SEAT.UPDATE, payload };
}

export function clearSeats(): SeatAction {
  return { type: SEAT.CLEAR, payload: undefined };
}
