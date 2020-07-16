import { Hand, SEAT } from '../../models';
import { Action } from 'redux';
import { Payload } from './base';

const PREFIX = '[HAND]';

export const HAND = Object.freeze({
  DEAL: `${PREFIX} DEAL`,
  UPDATE: `${PREFIX} UPDATE`,
});

export type HandDealAction = Action<typeof HAND.DEAL> & Payload<Hand[]>;
export type HandUpdateAction = Action<typeof HAND.UPDATE> & Payload<Record<SEAT, Hand[]>>;

export type HandAction = HandDealAction | HandUpdateAction;
