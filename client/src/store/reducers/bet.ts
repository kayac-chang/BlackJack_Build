import { Bet } from '../../models';
import { BetAction, BET } from '../types';

type BetState = {
  chosen?: Bet;
  previous: Bet[];
  history: Bet[];
};

const initialState: BetState = {
  previous: [],
  history: [],
};

export default function betReducer(state = initialState, action: BetAction): BetState {
  const { type, payload } = action;

  if (type === BET.CHOOSE) {
    const bet = payload as Bet;

    return { ...state, chosen: bet };
  }

  if (type === BET.ADD) {
    const bet = payload as Bet;

    return { ...state, history: [...state.history, bet] };
  }

  if (type === BET.UNDO) {
    const { time } = payload as Bet;

    const history = state.history.filter((record) => record.time !== time);

    return { ...state, history };
  }

  if (type === BET.CLEAR) {
    //
    return { ...state, history: [] };
  }

  if (type === BET.COMMIT) {
    const bets = payload as Bet[];

    return { ...state, previous: bets };
  }

  if (type === BET.REPLACE) {
    const bets = payload as Bet[];

    return { ...state, history: bets };
  }

  return state;
}
