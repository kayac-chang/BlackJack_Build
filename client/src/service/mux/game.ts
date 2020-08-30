import { S2C, Hand, RANK } from '../../models';
import Service from '../service';

import store from '../../store';
import {
  betStart,
  betEnd,
  settle,
  countdown,
  clearBet,
  dealCard,
  turn,
  update,
  updateSeats,
  editRoom,
  replaceBet,
} from '../../store/actions';

import {
  GameProp,
  toGame,
  toGameState,
  CountDownProp,
  DealProp,
  toHand,
  TurnProp,
  toSeatNum,
  toPair,
  toDecision,
  toSeats,
} from '../types';
import { pipe } from 'ramda';
import { wait, looper } from '../../utils';
import { SEAT } from '../../models';

function updateHistory(data: GameProp) {
  const { room } = store.getState();

  const found = room.find(({ id }) => data.id === id);
  if (!found) {
    return;
  }

  store.dispatch(
    editRoom({
      ...found,
      history: data.history.map(String),
    })
  );
}

function onBetStart(service: Service, data: GameProp) {
  store.dispatch(
    betStart(
      toGame({
        ...data,
        state: [S2C.ROUND.BET_START],
      })
    )
  );

  updateHistory(data);

  store.dispatch(countdown(20));
  store.dispatch(updateSeats(toSeats(data.seats)));
}

function onCountDown(service: Service, { expire }: CountDownProp) {
  store.dispatch(countdown(expire));
}

function onBetEnd(service: Service, { state }: GameProp) {
  const { game, seat, bet, user } = store.getState();

  store.dispatch(
    betEnd({
      ...game,
      state: toGameState(state),
    })
  );

  const history = bet.history.filter((bet) => {
    if (!bet.seat) {
      return false;
    }

    if (seat[bet.seat].commited) {
      return true;
    }

    seat[bet.seat].bet = 0;

    return false;
  });

  store.dispatch(update(user));
  store.dispatch(updateSeats(seat));
  store.dispatch(replaceBet(history));
}

function onSettle(service: Service, data: GameProp) {
  const { game, user } = store.getState();

  for (const seat of data.seats) {
    if (Array.isArray(seat.piles) && seat.piles.length > 0) {
      seat.pay = seat.piles.reduce((acc, { pay }) => acc + pay, 0);
    }
  }

  store.dispatch(updateSeats(toSeats(data.seats)));

  store.dispatch(
    settle({
      ...game,
      state: toGameState([S2C.ROUND.SETTLE]),
    })
  );

  store.dispatch(clearBet(user));
}

function prefix(prop: DealProp) {
  const cards = [...(prop.cards || []), prop.card];

  return { ...prop, cards };
}

function onBegin(service: Service, prop: DealProp[]) {
  const hands = prop.map(pipe(prefix, toHand));

  const group: Record<SEAT, Hand[]> = {
    [SEAT.DEALER]: [],
    [SEAT.A]: [],
    [SEAT.B]: [],
    [SEAT.C]: [],
    [SEAT.D]: [],
    [SEAT.E]: [],
  };

  for (const hand of hands) {
    group[hand.seat].push(hand);
  }

  for (const [, _hands] of Object.entries(group)) {
    if (hasAce(_hands)) {
      _hands[_hands.length - 1].points = toPoints(_hands);
    }
  }

  store.dispatch(dealCard(hands));
}

function hasAce(hands: Hand[]) {
  return hands.filter(({ card }) => card.rank === RANK.ACE).length;
}

function toPoints(hands: Hand[]) {
  let min = 0;
  let max = 0;

  for (const hand of hands) {
    if (hand.card.rank === RANK.ACE) {
      min += 1;
      max += max > 10 ? 1 : 11;
    }

    if (
      [RANK.TWO, RANK.THREE, RANK.FOUR, RANK.FIVE, RANK.SIX, RANK.SEVEN, RANK.EIGHT, RANK.NINE].includes(hand.card.rank)
    ) {
      min += Number(hand.card.rank);
      max += Number(hand.card.rank);
    }

    if ([RANK.TEN, RANK.JACK, RANK.QUEEN, RANK.KING].includes(hand.card.rank)) {
      min += 10;
      max += 10;
    }
  }

  if (max > 21 || max === min) {
    return String(min);
  }

  return `${min} / ${max}`;
}

function onDeal(service: Service, prop: DealProp) {
  const { hand } = store.getState();

  const latest = toHand(prop);
  const hands = hand[latest.seat];

  const pair = [...hands, latest].filter((hand) => hand.pair === latest.pair);

  if (hasAce(pair)) {
    latest.points = toPoints(pair);
  }

  store.dispatch(dealCard([latest]));
}

let cancel: () => void;

function onTurn(service: Service, { no, pile }: TurnProp) {
  const { game } = store.getState();

  if (cancel) {
    cancel();
  }

  store.dispatch(
    turn({
      ...game,
      turn: {
        seat: toSeatNum(no),
        pair: toPair(pile),
      },
    })
  );
}

async function onAction(service: Service, { expire, options }: TurnProp) {
  const { user } = store.getState();

  store.dispatch(
    update({
      ...user,
      decisions: toDecision(options),
    })
  );

  if (cancel) {
    cancel();
  }

  cancel = looper(async () => {
    onCountDown(service, { expire });

    await wait(1000);

    expire -= 1;

    return expire > 0;
  });
}

export default {
  [S2C.ROUND.BET_START]: onBetStart,
  [S2C.ROUND.COUNT_DOWN]: onCountDown,
  [S2C.ROUND.BET_END]: onBetEnd,
  [S2C.ROUND.SETTLE]: onSettle,
  [S2C.ROUND.BEGIN]: onBegin,
  [S2C.ROUND.DEAL]: onDeal,
  [S2C.ROUND.TURN]: onTurn,
  [S2C.ROUND.ACTION]: onAction,
};
