import { Card, SUIT, RANK } from '../../models';

export function toSuit(suit: string): SUIT {
  const mapping: { [key: string]: SUIT } = {
    A: SUIT.SPADE,
    B: SUIT.HEART,
    C: SUIT.DIAMOND,
    D: SUIT.CLUB,
  };

  if (mapping[suit]) {
    return mapping[suit];
  }

  throw new Error(`Not support suit type: ${suit}`);
}

export function toRank(rank: string): RANK {
  const mapping: { [key: string]: RANK } = {
    1: RANK.ACE,
    2: RANK.TWO,
    3: RANK.THREE,
    4: RANK.FOUR,
    5: RANK.FIVE,
    6: RANK.SIX,
    7: RANK.SEVEN,
    8: RANK.EIGHT,
    9: RANK.NINE,
    A: RANK.TEN,
    B: RANK.JACK,
    D: RANK.QUEEN,
    E: RANK.KING,
  };

  if (mapping[rank]) {
    return mapping[rank];
  }

  throw new Error(`Not support rank type: ${rank}`);
}

export function toCard([suit, rank]: string): Card {
  return {
    suit: toSuit(suit),
    rank: toRank(rank),
  };
}
