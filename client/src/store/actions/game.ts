import { Game } from '../../models';
import { GAME, GameAction } from '../types';

export function join(payload: Game): GameAction {
  return { type: GAME.JOIN, payload };
}

export function betStart(payload: Game): GameAction {
  return { type: GAME.BET_START, payload };
}

export function betEnd(payload: Game): GameAction {
  return { type: GAME.BET_END, payload };
}

export function settle(payload: Game): GameAction {
  return { type: GAME.SETTLE, payload };
}

export function turn(payload: Game): GameAction {
  return { type: GAME.TURN, payload };
}

export function countdown(payload: number): GameAction {
  return { type: GAME.COUNT_DOWN, payload };
}
