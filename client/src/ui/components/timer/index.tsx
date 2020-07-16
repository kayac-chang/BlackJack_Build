import React from 'react';
import styles from './Timer.module.scss';
import { Triangle } from '../shape';

type Props = {
  total: number;
  countdown: number;
};

export default function Timer({ total, countdown }: Props) {
  const len = Math.floor((countdown / total) * 100);

  return (
    <div className={styles.timer}>
      <Triangle direction={'left'} len={len} color={'rgba(48, 209, 88, 1)'} />
      <h5>{countdown}</h5>
      <Triangle direction={'right'} len={len} color={'rgba(48, 209, 88, 1)'} />
    </div>
  );
}
