import React, { useLayoutEffect } from 'react';
import styles from './Decision.module.scss';
import { useSelector } from 'react-redux';
import { AppState } from '../../store';
import { GAME_STATE } from '../../models';
import Timer from '../components/timer';
import Controls from './Controls';
import { animated, useSpring } from 'react-spring';

export default function Decision() {
  const user = useSelector((state: AppState) => state.user);
  const { state, countdown, turn } = useSelector((state: AppState) => state.game);
  const seat = useSelector((state: AppState) => state.seat);
  const decisions = useSelector((state: AppState) => state.user.decisions);

  const isDealing = state === GAME_STATE.DEALING && countdown > 1;
  const isUserTurn = turn ? seat[turn.seat].player === user.name : false;
  const hasCommited = !(decisions.length > 0);

  const [style, setOpacity] = useSpring(() => ({
    opacity: 0,
    display: 'none',
  }));

  useLayoutEffect(() => {
    if (isDealing && isUserTurn && hasCommited) {
      setOpacity({ opacity: 0.3, display: 'block' });
      return;
    }

    if (isDealing && isUserTurn) {
      setOpacity({ opacity: 1, display: 'block' });
      return;
    }

    setOpacity({ opacity: 0, display: 'none' });
  }, [setOpacity, isDealing, isUserTurn, hasCommited]);

  return (
    <animated.div className={styles.decision} style={style}>
      <div>
        <h3 className={styles.title}>make your decision</h3>

        <Controls enable={isUserTurn && isDealing} decisions={decisions} />

        <Timer total={10} countdown={countdown} />
      </div>
    </animated.div>
  );
}
