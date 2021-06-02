import React, { useCallback } from 'react';
import { X, CornerUpLeft, RotateCw } from 'react-feather';
import Control from '../components/button/Control';
import styles from './Bet.module.scss';
import { useSelector, useDispatch } from 'react-redux';
import { AppState } from '../../store';
import { clearBet, undoBet, replaceBet } from '../../store/actions';
import services from '../../service';
import { throttleBy } from '../../utils';
import RES from '../../assets';

type Props = {
  enable: boolean;
  setCommited: (flag: boolean) => void;
};

export default function Controls ({ enable, setCommited }: Props) {
  const dispatch = useDispatch();
  const user = useSelector((state: AppState) => state.user);
  const seats = useSelector((state: AppState) => state.seat);
  const maxBet = useSelector((state: AppState) => state.game.bet.max);
  const countdown = useSelector((state: AppState) => state.game.countdown);

  const history = useSelector((state: AppState) => state.bet.history);
  const isDealable = history.length > 0 && enable;

  const previous = useSelector((state: AppState) => state.bet.previous);
  const isRepeatable = previous.length > 0 && enable;

  const onClear = useCallback(
    throttleBy(async function () {
      if (countdown <= 3) {
        return;
      }

      const tasks = Object.entries(seats)
        .filter(([, seat]) => seat.player === user.name && !seat.commited)
        .map(([id]) => services.leaveSeat(Number(id)));

      await Promise.all(tasks);

      dispatch(clearBet(user));
    }),
    [countdown, user]
  );

  const onUndo = useCallback(
    function () {
      if (countdown <= 2) {
        return;
      }

      const last = history[history.length - 1];

      last && dispatch(undoBet(last));
    },
    [countdown, dispatch, history]
  );

  const onDeal = useCallback(
    throttleBy(async function () {
      if (!enable || countdown < 2 || history.length <= 0) {
        return;
      }

      setCommited(true);

      await services.deal();
    }),
    [enable, history, countdown]
  );

  const onRepeat = useCallback(
    async function () {
      if (!enable) {
        return;
      }

      dispatch(clearBet(user));

      const tasks = previous.map(({ seat }) => {
        if (!seat) {
          return Promise.resolve();
        }

        if (seats[seat].player === user.name) {
          return Promise.resolve();
        }

        return services.joinSeat(seat);
      });

      await Promise.all(tasks);

      dispatch(replaceBet(previous));
    },
    [seats, dispatch, enable, user, previous]
  );

  const onDouble = useCallback(() => {
    if (countdown <= 2) {
      return;
    }

    const newBet = history.map(bet => ({ ...bet, amount: bet.amount * 2 }));

    const total = newBet.reduce((acc, { amount }) => acc + amount, 0);

    if (maxBet > user.totalBet + total) {
      dispatch(clearBet(user));
      dispatch(replaceBet(newBet));
    }
  }, [dispatch, countdown, user, history, maxBet]);

  return (
    <div className={styles.controls}>
      <Control title={'clear'} icon={<X />} onClick={onClear} enable={enable} />
      <Control
        title={'undo'}
        style={{ opacity: isDealable ? 1 : 0.3 }}
        icon={<CornerUpLeft />}
        onClick={onUndo}
        enable={isDealable}
      />
      <Control
        title={'deal'}
        icon={<img src={RES.getBase64('ICON_DEAL')} alt={'ICON_DEAL'} />}
        style={{ opacity: isDealable ? 1 : 0.3 }}
        onClick={onDeal}
        enable={isDealable}
      />
      <Control
        title={'repeat'}
        icon={<RotateCw />}
        style={{ opacity: isRepeatable ? 1 : 0.3 }}
        onClick={onRepeat}
        enable={isRepeatable}
      />
      <Control
        title={'double'}
        style={{ opacity: isDealable ? 1 : 0.3 }}
        icon={<h3>2x</h3>}
        onClick={onDouble}
        enable={isDealable}
      />
    </div>
  );
}
