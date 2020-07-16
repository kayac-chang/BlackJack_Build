import { BetAction, BET } from '../types';
import { Bet, User } from '../../models';

export function choose(payload: Bet): BetAction {
  return { type: BET.CHOOSE, payload };
}

export function addBet(payload: Bet): BetAction {
  return { type: BET.ADD, payload };
}

export function clearBet(payload: User): BetAction {
  return { type: BET.CLEAR, payload };
}

export function undoBet(payload: Bet): BetAction {
  return { type: BET.UNDO, payload };
}

export function replaceBet(payload: Bet[]): BetAction {
  return { type: BET.REPLACE, payload };
}

export function commitBet(payload: Bet[]): BetAction {
  return { type: BET.COMMIT, payload };
}
