import { Range } from './base';
import { SEAT, PAIR, GAME_STATE, SUIT, RANK, CHIP, RESULT } from './enum';

export interface Turn {
  seat: SEAT;
  pair: PAIR;
}

export interface Game {
  room: number;
  round: string;
  bet: Range;
  state: GAME_STATE;
  countdown: number;
  turn?: Turn;
}

export type Seats = Record<SEAT, Seat>;

export interface Seat {
  player: string;
  bet: number;
  commited?: boolean;
  split: boolean;
  pays: Record<PAIR, number>;
  results: Record<PAIR, RESULT>;
}

export interface Card {
  suit: SUIT;
  rank: RANK;
}

export interface Bet {
  chip: CHIP;
  amount: number;
  seat?: SEAT;
  time?: Date;
}

export interface Hand {
  id: string;
  seat: SEAT;
  card: Card;
  points: number;
  pair: PAIR;
}
