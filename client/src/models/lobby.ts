import { Range } from './base';
import { DECISION } from './enum';
import TABLES from '../assets/table';

export interface User {
  name: string;
  reward: number;
  balance: number;
  totalBet: number;
  decisions: DECISION[];
  table: keyof typeof TABLES;
}

export type Token = string;

export interface Room {
  id: number;
  history: string[];
  bet: Range;
  people: number;
}
