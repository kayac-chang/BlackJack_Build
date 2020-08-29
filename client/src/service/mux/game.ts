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
  updateHand,
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

  store.dispatch(dealCard(hands));
}

function hasAce(hands: Hand[]) {
  return hands.filter(({ card }) => card.rank === RANK.ACE).length;
}

function onDeal(service: Service, prop: DealProp) {
  const newHand = toHand(prop);

  store.dispatch(dealCard([newHand]));

  const { hand } = store.getState();

  const hands = hand[newHand.seat];
  const latest = hands[hands.length - 1];

  if (hasAce(hands) && Number(latest.points) > 11) {
    latest.points = `${Number(latest.points) - 11} / ${latest.points}`;

    store.dispatch(updateHand(hand));
  }
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
