import { User, Bet } from '../../models';
import { UserAction, USER, BET, BetAction } from '../types';
import { v4 } from 'uuid';

const initialState: User = {
  name: v4(),
  reward: 0,
  balance: 0,
  totalBet: 0,
  decisions: [],
  table: 'TABLE_BLUE',
};

export default function userReducer(state = initialState, action: UserAction | BetAction): User {
  const { type, payload } = action;

  if ([USER.LOGIN, USER.UPDATE].includes(type)) {
    const user = payload as User;

    return { ...state, ...user };
  }

  if (type === BET.ADD) {
    const { amount } = payload as Bet;

    return {
      ...state,
      balance: state.balance - amount,
      totalBet: state.totalBet + amount,
    };
  }

  if (type === BET.REPLACE) {
    const bets = payload as Bet[];

    const totalBet = bets.reduce((acc, { amount }) => acc + amount, 0);

    return {
      ...state,
      balance: state.balance + state.totalBet - totalBet,
      totalBet,
    };
  }

  if (type === BET.UNDO) {
    const { amount } = payload as Bet;

    return {
      ...state,
      balance: state.balance + amount,
      totalBet: state.totalBet - amount,
    };
  }

  if (type === BET.CLEAR) {
    return {
      ...state,
      balance: state.balance + state.totalBet,
      totalBet: 0,
    };
  }

  return state;
}
