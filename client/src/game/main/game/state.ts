import { createMachine, interpret, assign, Interpreter, State } from 'xstate';
import store, { observe } from '../../../store';
import { Hand, SEAT, GAME_STATE, PAIR, RESULT, Turn } from '../../../models';
import { Container } from 'pixi.js';
import Poker from './Poker';
import { move, tween } from './anim';
import { origin, config } from './static';
import { Lose, Win, Bust } from './icon';
import { createField, Field } from '../field';

interface Context {
  history: Hand[];
  scores: Record<PAIR, number>;
  split: boolean;
}

type Event =
  | { type: 'INIT_NORMAL' }
  | { type: 'INIT_SPLIT'; hands: Hand[] }
  | { type: 'DEAL_NORMAL'; hand: Hand }
  | { type: 'DEAL_SPLIT'; hand: Hand }
  | { type: 'RESULT_NORMAL'; result: RESULT }
  | { type: 'RESULT_SPLIT'; result: RESULT }
  | { type: 'CLEAR' };

type Schema<T> =
  | { value: { normal: 'idle' }; context: T }
  | { value: { normal: 'deal' }; context: T }
  | { value: { normal: 'result' }; context: T }
  | { value: { split: 'idle' }; context: T }
  | { value: { split: 'deal' }; context: T }
  | { value: { split: 'result' }; context: T };

export type HandsService = Interpreter<Context, any, Event, Schema<Context>>;
export type HandsState = State<Context, Event, any, Schema<Context>>;

function toIcon(result: RESULT): Container | undefined {
  if (result === RESULT.LOSE) {
    return Lose();
  }

  if (result === RESULT.WIN) {
    return Win();
  }

  if (result === RESULT.BUST) {
    return Bust();
  }
}

type Elements = {
  handL: Container;
  handR: Container;
  fieldL: Field;
  fieldR: Field;
  results: Container;
};

function createHandMachine(id: SEAT, { handL, handR, fieldL, fieldR, results }: Elements) {
  //
  return createMachine<Context, Event, Schema<Context>>(
    {
      type: 'parallel',

      context: {
        history: [],
        scores: {
          [PAIR.L]: 0,
          [PAIR.R]: 0,
        },
        split: false,
      },

      states: {
        //
        normal: {
          initial: 'idle',

          states: {
            idle: {
              on: {
                INIT_NORMAL: { target: 'deal' },
              },
            },

            deal: {
              on: {
                DEAL_NORMAL: { target: 'deal', actions: ['updateHistory', 'updatePokers', 'updateScores'] },
                RESULT_NORMAL: { target: 'result', actions: 'updateResults' },
              },
            },

            result: {
              on: {
                CLEAR: {
                  target: 'idle',
                  actions: ['updateHistory', 'updatePokers', 'updateScores', 'updateResults', 'updateSplit'],
                },
              },
            },
          },
        },

        split: {
          initial: 'idle',

          states: {
            idle: {
              on: {
                INIT_SPLIT: { target: 'deal', actions: ['updateHistory', 'updateSplit'] },
              },
            },

            deal: {
              on: {
                DEAL_SPLIT: { target: 'deal', actions: ['updateHistory', 'updatePokers', 'updateScores'] },
                RESULT_SPLIT: { target: 'result', actions: 'updateResults' },
              },
            },

            result: {
              on: {
                CLEAR: {
                  target: 'idle',
                  actions: ['updateHistory', 'updatePokers', 'updateScores', 'updateResults', 'updateSplit'],
                },
              },
            },
          },
        },
      },
    },
    //
    {
      actions: {
        updateHistory: assign({
          history: (context, event) => {
            if (event.type === 'DEAL_NORMAL' || event.type === 'DEAL_SPLIT') {
              return [...context.history, event.hand];
            }

            if (event.type === 'INIT_SPLIT') {
              return event.hands;
            }

            if (event.type === 'CLEAR') {
              return [];
            }

            return context.history;
          },
        }),

        updateSplit: assign({
          split: (context, event) => {
            const mapping = {
              [PAIR.L]: handL,
              [PAIR.R]: handR,
            };

            if (event.type === 'INIT_SPLIT') {
              [...handL.children, ...handR.children].forEach((poker) => {
                const target = event.hands.find(({ id }) => id === poker.name);
                if (!target) {
                  return;
                }

                mapping[target.pair].addChild(poker);
              });

              for (const [pair, hand] of Object.entries(mapping)) {
                move(hand.children, config['split'][id][pair as PAIR]);
              }

              return true;
            }

            if (event.type === 'CLEAR') {
              return false;
            }

            return context.split;
          },
        }),

        updatePokers: (context, event) => {
          const mapping = {
            [PAIR.L]: handL,
            [PAIR.R]: handR,
          };

          if (id === SEAT.DEALER && event.type === 'DEAL_NORMAL') {
            const { card, pair } = event.hand;

            const fold = mapping[pair].children.find((poker) => !(poker as Poker).faceUp);
            if (fold) {
              const poker = new Poker(card.suit, card.rank);

              poker.faceUp = false;
              poker.position.set(fold.position.x, fold.position.y);
              mapping[pair].addChild(poker);
              mapping[pair].removeChild(fold);
              poker.flip();

              return;
            }

            const poker = new Poker(card.suit, card.rank);
            poker.name = event.hand.id;
            poker.alpha = 0;
            poker.position.set(origin.x, origin.y);
            mapping[pair].addChild(poker);

            if (context.history.length === 1) {
              const fold = new Poker(card.suit, card.rank);

              fold.faceUp = false;
              fold.alpha = 0;
              fold.position.set(origin.x, origin.y);
              mapping[pair].addChild(fold);
            }

            const pos = !context.split ? config['normal'][id] : config['split'][id][pair];
            move(mapping[pair].children, pos);

            return;
          }

          if (event.type === 'DEAL_NORMAL') {
            const { card, pair } = event.hand;

            const poker = new Poker(card.suit, card.rank);
            poker.name = event.hand.id;
            poker.alpha = 0;
            poker.position.set(origin.x, origin.y);
            mapping[pair].addChild(poker);

            const pos = !context.split ? config['normal'][id] : config['split'][id][pair];
            move(mapping[pair].children, pos);

            return;
          }

          if (event.type === 'DEAL_SPLIT') {
            const { card, pair } = event.hand;

            const poker = new Poker(card.suit, card.rank);
            poker.name = event.hand.id;
            poker.alpha = 0;
            poker.position.set(origin.x, origin.y);
            mapping[pair].addChild(poker);

            move(mapping[pair].children, config['split'][id][pair]);

            return;
          }

          if (event.type === 'CLEAR') {
            mapping[PAIR.L].removeChildren();
            mapping[PAIR.R].removeChildren();

            return;
          }
        },

        updateScores: assign({
          scores: (context, event) => {
            const mapping = {
              [PAIR.L]: fieldL,
              [PAIR.R]: fieldR,
            };

            const offsetY = 70;

            if (event.type === 'DEAL_NORMAL') {
              const { pair, points } = event.hand;

              const pos = !context.split ? config['normal'][id] : config['split'][id][pair];
              mapping[pair].position.set(pos.x, pos.y + offsetY);
              mapping[pair].text = String(points);
              mapping[pair].parent.addChild(mapping[pair]);

              return {
                ...context.scores,
                [pair]: points,
              };
            }

            if (event.type === 'DEAL_SPLIT') {
              const { pair, points } = event.hand;

              const pos = !context.split ? config['normal'][id] : config['split'][id][pair];
              mapping[pair].position.set(pos.x, pos.y + offsetY);
              mapping[pair].text = String(points);
              mapping[pair].parent.addChild(mapping[pair]);

              return {
                ...context.scores,
                [pair]: points,
              };
            }

            if (event.type === 'CLEAR') {
              mapping[PAIR.L].text = '';
              mapping[PAIR.R].text = '';

              return {
                [PAIR.L]: 0,
                [PAIR.R]: 0,
              };
            }

            return context.scores;
          },
        }),

        updateResults: (context, event) => {
          if (id === SEAT.DEALER) {
            return;
          }

          const offsetY = -200;

          if (event.type === 'RESULT_NORMAL') {
            const icon = toIcon(event.result);
            if (!icon) {
              return;
            }

            const pos = !context.split ? config['normal'][id] : config['split'][id][PAIR.L];
            icon.position.set(pos.x, pos.y);

            results.addChild(icon);

            tween(icon, { y: pos.y + offsetY });
          }

          if (event.type === 'RESULT_SPLIT') {
            const icon = toIcon(event.result);
            if (!icon) {
              return;
            }

            const pos = config['split'][id][PAIR.R];
            icon.position.set(pos.x, pos.y);

            results.addChild(icon);

            tween(icon, { y: pos.y + offsetY });
          }

          if (event.type === 'CLEAR') {
            results.removeChildren();
          }
        },
        //
      },
    }
  );
}

function onGameStateChange(service: HandsService, id: SEAT) {
  let hasJoin = false;

  return function (state: GAME_STATE) {
    const { seat } = store.getState();

    const canJoin = id === SEAT.DEALER || (seat[id].player && seat[id].bet);

    if (state === GAME_STATE.BETTING && hasJoin) {
      hasJoin = false;

      service.send({ type: 'CLEAR' });

      return;
    }

    if (state === GAME_STATE.DEALING && canJoin) {
      hasJoin = true;

      service.send({ type: 'INIT_NORMAL' });

      return;
    }

    if (state === GAME_STATE.SETTLE) {
      service.send({ type: 'RESULT_SPLIT', result: seat[id].results.R });
      service.send({ type: 'RESULT_NORMAL', result: seat[id].results.L });

      return;
    }
  };
}

function onHandsChange(service: HandsService, id: SEAT) {
  let last: Hand[] = [];

  return function (hands: Hand[]) {
    const latest = hands.slice(last.length);
    last = hands;

    if (latest.length <= 0) {
      return;
    }

    for (const hand of latest) {
      if (hand.seat !== id) {
        continue;
      }

      if (hand.pair === PAIR.L) {
        service.send({ type: 'DEAL_NORMAL', hand });
      }

      if (hand.pair === PAIR.R) {
        service.send({ type: 'DEAL_SPLIT', hand });
      }
    }
  };
}

function onSplit(service: HandsService, id: SEAT) {
  return function (split: boolean) {
    if (!split) {
      return;
    }

    const { hand } = store.getState();
    service.send({ type: 'INIT_SPLIT', hands: hand[id] });
  };
}

function onTurnChange(service: HandsService, id: SEAT, { handL, handR }: { handL: Container; handR: Container }) {
  const mapping = {
    [PAIR.L]: handL,
    [PAIR.R]: handR,
  };

  return function (turn?: Turn) {
    if (!turn || turn.seat !== id) {
      tween(handL, { alpha: 1 });
      tween(handR, { alpha: 1 });

      return;
    }

    for (const [pair, hand] of Object.entries(mapping)) {
      const alpha = turn.pair === pair ? 1 : 0.5;

      tween(hand, { alpha });
    }
  };
}

export function createHandService(id: SEAT, container: Container): HandsService {
  const handL = new Container();
  const handR = new Container();
  container.addChild(handL, handR);

  const fieldL = createField({ fontSize: 48 });
  const fieldR = createField({ fontSize: 48 });
  container.addChild(fieldL, fieldR);

  const results = new Container();
  container.addChild(results);

  const service = interpret(createHandMachine(id, { handL, handR, fieldL, fieldR, results }));

  observe((state) => state.game.state, onGameStateChange(service, id));
  observe((state) => state.game.turn, onTurnChange(service, id, { handL, handR }));
  observe((state) => state.hand[id], onHandsChange(service, id));
  observe((state) => state.seat[id].split, onSplit(service, id));

  service.onTransition((state) => {
    if (!state.changed) {
      return;
    }

    if (state.context.scores.L > 21) {
      service.send({ type: 'RESULT_NORMAL', result: RESULT.BUST });
    }

    if (state.context.scores.R > 21) {
      service.send({ type: 'RESULT_SPLIT', result: RESULT.BUST });
    }
  });

  return service;
}
