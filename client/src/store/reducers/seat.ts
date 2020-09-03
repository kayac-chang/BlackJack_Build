import { SEAT as SEAT_ID, Bet, Seats, User, PAIR, RESULT } from '../../models';
import { SeatAction, SEAT, BET, BetAction, GAME } from '../types';
import { v4 } from 'uuid';

const dealer = {
  player: v4(),
  bet: 0,
  commited: false,
  split: false,

  pays: {
    [PAIR.L]: 0,
    [PAIR.R]: 0,
  },

  results: {
    [PAIR.L]: RESULT.LOSE,
    [PAIR.R]: RESULT.LOSE,
  },
};

const Seat = () => ({
  player: '',
  bet: 0,
  commited: false,
  split: false,

  pays: {
    [PAIR.L]: 0,
    [PAIR.R]: 0,
  },

  results: {
    [PAIR.L]: RESULT.LOSE,
    [PAIR.R]: RESULT.LOSE,
  },
});

const initialState: Seats = {
  [SEAT_ID.DEALER]: dealer,
  [SEAT_ID.A]: Seat(),
  [SEAT_ID.B]: Seat(),
  [SEAT_ID.C]: Seat(),
  [SEAT_ID.D]: Seat(),
  [SEAT_ID.E]: Seat(),
};

export default function seatReducer(state = initialState, action: SeatAction | BetAction): Seats {
  const { type, payload } = action;

  if (type === SEAT.UPDATE) {
    const seats = payload as Seats;

    const newState = {} as Seats;

    for (const [id, seat] of Object.entries(seats)) {
      const seatID = Number(id) as SEAT_ID;

      newState[seatID] = { ...state[seatID], ...seat };
    }

    return { ...state, ...newState };
  }

  if (type === GAME.SETTLE) {
    const newState = {} as Seats;

    for (const [id, seat] of Object.entries(state)) {
      newState[Number(id) as SEAT_ID] = { ...seat, commited: false, split: false };
    }

    return newState;
  }

  if (type === SEAT.CLEAR) {
    //
    return initialState;
  }

  if (type === BET.ADD) {
    const { seat, amount } = payload as Bet;

    if (seat === undefined) {
      return state;
    }

    const target = state[seat];

    return {
      ...state,
      [seat]: { ...target, bet: target.bet + amount },
    };
  }

  if (type === BET.UNDO) {
    const { seat, amount } = payload as Bet;

    if (seat === undefined) {
      return state;
    }

    const target = state[seat];

    return {
      ...state,
      [seat]: { ...target, bet: target.bet - amount },
    };
  }

  if (type === BET.COMMIT) {
    const bets = payload as Bet[];

    const newState = {} as Seats;

    for (const { seat } of bets) {
      if (seat === undefined) continue;

      if (!newState[seat]) {
        newState[seat] = { ...state[seat] };
      }
    }

    return {
      ...state,
      ...newState,
    };
  }

  if (type === BET.CLEAR) {
    const { name } = payload as User;

    const newState = {} as Seats;
    for (const [id, seat] of Object.entries(state)) {
      if (seat.player !== name) {
        continue;
      }

      if (seat.commited) {
        continue;
      }

      const seatID = Number(id) as SEAT_ID;

      if (!newState[seatID]) {
        newState[seatID] = { ...state[seatID] };
      }

      newState[seatID].bet = 0;
    }

    return {
      ...state,
      ...newState,
    };
  }

  if (type === BET.REPLACE) {
    const bets = payload as Bet[];

    const newState = {} as Seats;

    for (const { seat, amount } of bets) {
      if (seat === undefined) continue;

      if (!newState[seat]) {
        newState[seat] = { ...state[seat], bet: 0 };
      }

      newState[seat].bet += amount;
    }

    return {
      ...state,
      ...newState,
    };
  }

  return state;
}
