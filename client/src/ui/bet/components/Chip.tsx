import React, { PropsWithChildren, HTMLAttributes } from 'react';
import styles from './Chip.module.scss';

type Props = {
  src: string;
  bet: number;
  selected: boolean;
} & PropsWithChildren<HTMLAttributes<HTMLButtonElement>>;

function format(bet: number) {
  if (bet / 1000 >= 1) {
    return String(Math.floor(bet / 1000)) + 'K';
  }

  return String(bet);
}

export default function Chip({ selected, src, bet, onClick }: Props) {
  const className = [
    //
    styles.chip,
    selected && styles.chosen,
  ]
    .filter(Boolean)
    .join(' ');

  return (
    <button className={className} onClick={onClick}>
      <h5>{format(bet)}</h5>
      <img src={src} alt={format(bet)} />
    </button>
  );
}
