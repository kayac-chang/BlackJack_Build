import React, { PropsWithChildren, HTMLAttributes } from 'react';
import styles from './Lobby.module.scss';
import ARROW from './assets/arrow.png';
import clsx from 'clsx';

type Props = PropsWithChildren<HTMLAttributes<HTMLDivElement>>;

export default function Arrow({ className, style, onClick }: Props) {
  return (
    <div className={clsx(styles.control, className)} style={style} onClick={onClick}>
      <img src={ARROW} alt={ARROW} />
    </div>
  );
}
