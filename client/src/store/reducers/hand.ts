import { Hand, SEAT } from '../../models';
import { HAND, HandAction, GAME } from '../types';
import { groupBy, mergeWith, concat } from 'ramda';

const initialState: Record<SEAT, Hand[]> = {
  [SEAT.DEALER]: [],
  [SEAT.A]: [],
  [SEAT.B]: [],
  [SEAT.C]: [],
  [SEAT.D]: [],
  [SEAT.E]: [],
};

const groupByID = groupBy(({ seat }: Hand) => String(seat));

export default function handReducer(state = initialState, action: HandAction): Record<SEAT, Hand[]> {
  const { type, payload } = action;

  if (type === HAND.DEAL) {
    const hands = payload as Hand[];

    const grouped = groupByID(hands);

    return mergeWith(concat, state, grouped);
  }

  if (type === HAND.UPDATE) {
    return payload as Record<SEAT, Hand[]>;
  }

  if (type === GAME.BET_START) {
    return initialState;
  }

  return state;
}
