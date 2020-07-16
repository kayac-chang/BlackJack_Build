import { createMachine, interpret, State, Interpreter } from 'xstate';
import store from '../../../store';
import services from '../../../services';
import { SEAT, GAME_STATE } from '../../../models';
import { addBet } from '../../../store/actions';

interface Context {
  id: SEAT;
  owner: string;
}

type Event = { type: 'JOIN'; user: string } | { type: 'LEAVE' } | { type: 'pointerdown' } | { type: 'STATE' };

type Schema<T> =
  | { value: 'empty'; context: T }
  | { value: 'fetching'; context: T }
  | { value: 'occupy'; context: T }
  | { value: { occupy: 'normal' }; context: T }
  | { value: { occupy: 'betting' }; context: T };

export type SeatState = State<Context, Event, any, Schema<Context>>;
export type SeatService = Interpreter<Context, any, Event, Schema<Context>>;

function updateOwner(context: Context, event: Event) {
  //
  if (event.type === 'JOIN') {
    context.owner = event.user;
  }

  if (event.type === 'LEAVE') {
    context.owner = '';
  }

  return context.owner;
}

function canJoin(context: Context, event: Event) {
  const { user, seat } = store.getState();

  return Object.values(seat).filter(({ player }) => player === user.name).length < 3;
}

function join(context: Context, event: Event) {
  services.joinSeat(context.id);
}

function canBet(context: Context, event: Event) {
  const { user, game } = store.getState();

  const isBetting = game.state === GAME_STATE.BETTING;

  const name = event.type === 'JOIN' ? event.user : context.owner;

  return isBetting && name === user.name;
}

function placeBet(context: Context, event: Event) {
  const { game, bet, user, seat } = store.getState();

  if (!bet.chosen) {
    return;
  }

  if (game.countdown <= 2 || seat[context.id].commited) {
    return;
  }

  if (user.totalBet + bet.chosen.amount > game.bet.max) {
    return;
  }

  store.dispatch(
    addBet({
      ...bet.chosen,
      time: new Date(),
      seat: context.id,
    })
  );
}

export function createSeatService(id: SEAT) {
  //
  const machine = createMachine<Context, Event, Schema<Context>>(
    {
      initial: 'empty',

      context: {
        id,
        owner: '',
      },

      states: {
        //
        empty: {
          on: {
            pointerdown: {
              target: 'fetching',
              cond: 'canJoin',
              actions: 'join',
            },
          },
        },

        fetching: {
          on: {
            JOIN: [
              //
              { target: 'occupy.betting', cond: 'canBet' },
              { target: 'occupy.normal' },
            ],
          },
        },

        occupy: {
          initial: 'normal',

          entry: 'updateOwner',

          states: {
            //
            normal: {},
            betting: {
              on: {
                pointerdown: {
                  actions: 'placeBet',
                },
              },
            },
          },

          on: {
            LEAVE: 'empty',
            STATE: [
              //
              { target: '.betting', cond: 'canBet' },
              { target: '.normal' },
            ],
          },
        },
      },
    },
    {
      guards: { canBet, canJoin },
      actions: { updateOwner, join, placeBet },
    }
  );

  return interpret(machine);
}
