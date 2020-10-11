import React from 'react';
import { useSelector } from 'react-redux';
import { animated, useSpring } from 'react-spring';
import { AppState } from '../../store';
import { currency } from '../../utils';
import styles from './Reward.module.scss';

export default function Reward() {
  const [style, setOpacity] = useSpring(() => ({
    opacity: 0,
    display: 'none',
  }));

  const reward = useSelector((state: AppState) => state.user.reward);

  const showReward = reward > 0 ? 1 : 0;
  setOpacity({ opacity: showReward, display: showReward ? 'block' : 'none' });

  return (
    <animated.div className={styles.reward} style={style}>
      <div className={styles.display}>
        <h1>you win</h1>
        <h3>{currency(reward)}</h3>
      </div>
    </animated.div>
  );
}
