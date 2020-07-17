import React, { useState, useEffect } from 'react';
import Timer from '../components/timer';
import styles from './Bet.module.scss';
import { GAME_STATE } from '../../models';
import { useSelector } from 'react-redux';
import { AppState } from '../../store';
import Controls from './Controls';
import Chips from './Chips';
import { animated, useSpring } from 'react-spring';

export default function Bet() {
  const user = useSelector((state: AppState) => state.user);
  const seats = useSelector((state: AppState) => state.seat);
  const { state, countdown } = useSelector((state: AppState) => state.game);

  const isBetting = state === GAME_STATE.BETTING && countdown > 1;
  const isUserJoin = Object.values(seats).some(({ player }) => user.name === player);

  const [hasCommited, setCommited] = useState(false);
  const [style, setOpacity] = useSpring(() => ({
    opacity: 0,
    display: 'none',
  }));

  useEffect(() => {
    const isCommited = Object.values(seats)
      .filter(({ player }) => player === user.name)
      .every(({ commited }) => commited);

    setCommited(isBetting && isCommited);
  }, [isBetting, seats, user]);

  useEffect(() => {
    if (isBetting && isUserJoin) {
      setOpacity({ opacity: 1, display: 'block' });
      return;
    }

    if (isBetting) {
      setOpacity({ opacity: 0.3, display: 'block' });
      return;
    }

    setOpacity({ to: [{ opacity: 0 }, { display: 'none' }] });
  }, [setOpacity, isBetting, isUserJoin, hasCommited]);

  const enable = isBetting && isUserJoin && !hasCommited;

  return (
    <animated.div className={styles.bet} style={style}>
      <div>
        <h3 className={styles.title}>place your bets</h3>

        <Chips enable={enable} />

        <Timer total={20} countdown={countdown} />

        <Controls enable={enable} />
      </div>
    </animated.div>
  );
}
