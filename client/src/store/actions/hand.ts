import { HandAction, HAND } from '../types';
import { Hand, SEAT } from '../../models';

export function dealCard(payload: Hand[]): HandAction {
  return { type: HAND.DEAL, payload };
}

export function updateHand(payload: Record<SEAT, Hand[]>): HandAction {
  return { type: HAND.UPDATE, payload };
}
