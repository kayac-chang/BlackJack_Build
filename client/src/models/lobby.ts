import { Range } from './base';
import { DECISION } from './enum';

export interface User {
  name: string;
  balance: number;
  totalBet: number;
  decisions: DECISION[];
}

export type Token = string;

export interface Room {
  id: number;
  history: string[];
  bet: Range;
  people: number;
}
