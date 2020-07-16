import React, { PropsWithChildren, HTMLAttributes } from 'react';
import styles from './Lobby.module.scss';
import ARROW from './assets/arrow.png';
import clsx from 'clsx';

type Props = {
  reverse?: boolean;
} & PropsWithChildren<HTMLAttributes<HTMLDivElement>>;

export default function Arrow({ style, reverse = false, onClick }: Props) {
  return (
    <div className={styles.control} style={style} onClick={onClick}>
      <img className={clsx(reverse && styles.reverse)} src={ARROW} alt={ARROW} />
    </div>
  );
}
