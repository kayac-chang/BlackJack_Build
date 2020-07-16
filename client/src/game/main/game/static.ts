import { SEAT, PAIR } from "../../../models";

export const origin = { x: 2515, y: 160 };

export const config = {
  //
  normal: {
    [SEAT.A]: { x: 443, y: 630 },
    [SEAT.B]: { x: 888, y: 880 },
    [SEAT.C]: { x: 1480, y: 980 },
    [SEAT.D]: { x: 2072, y: 880 },
    [SEAT.E]: { x: 2515, y: 630 },
    [SEAT.DEALER]: { x: 1480, y: 330 },
  },

  split: {
    [SEAT.A]: {
      [PAIR.L]: { x: 263, y: 630 },
      [PAIR.R]: { x: 623, y: 630 },
    },
    [SEAT.B]: {
      [PAIR.L]: { x: 708, y: 880 },
      [PAIR.R]: { x: 1068, y: 880 },
    },
    [SEAT.C]: {
      [PAIR.L]: { x: 1300, y: 980 },
      [PAIR.R]: { x: 1660, y: 980 },
    },
    [SEAT.D]: {
      [PAIR.L]: { x: 1892, y: 880 },
      [PAIR.R]: { x: 2252, y: 880 },
    },
    [SEAT.E]: {
      [PAIR.L]: { x: 2335, y: 630 },
      [PAIR.R]: { x: 2695, y: 630 },
    },
    [SEAT.DEALER]: {
      [PAIR.L]: { x: 1480, y: 330 },
      [PAIR.R]: { x: 1480, y: 330 },
    },
  },
};
