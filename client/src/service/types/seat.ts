import { SEAT, Seats, PAIR, RESULT } from '../../models';
import { SeatProp } from './prop';

export function toSeatNum(no: number): SEAT {
  if (no in SEAT) {
    return no as SEAT;
  }

  throw new Error(`Not support seat type: ${no}`);
}

export function isPlayerExist({ player }: SeatProp) {
  return Boolean(player);
}

function toResult(result = 20): RESULT {
  if (result === 21) {
    return RESULT.WIN;
  }

  if (result === 22) {
    return RESULT.PAID;
  }

  if (result === 23) {
    return RESULT.TIE;
  }

  if (result === 24) {
    return RESULT.LOSE;
  }

  if (result === 25) {
    return RESULT.SURRENDER;
  }

  return RESULT.LOSE;
}

export function toSeats(seats: SeatProp[]): Seats {
  //
  return seats.reduce((config, { no, player, total_bet, pay, piles }) => {
    return {
      ...config,
      [toSeatNum(no)]: {
        player: String(player),
        bet: Number(total_bet),
        split: false,
        pays: {
          [PAIR.L]: piles?.[0]?.pay || 0,
          [PAIR.R]: piles?.[1]?.pay || 0,
        },
        results: {
          [PAIR.L]: toResult(piles?.[0]?.result),
          [PAIR.R]: toResult(piles?.[0]?.result),
        },
      },
    };
  }, {} as Seats);
}
