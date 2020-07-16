import React, { PropsWithChildren, HTMLAttributes } from 'react';
import { useTrigger } from '../../hooks';
import styles from './Toggle.module.scss';

type Props = PropsWithChildren<HTMLAttributes<HTMLInputElement>>;

export default function Toggle({ id }: Props) {
  const [flag, trigger] = useTrigger();

  return (
    <div className={styles.toggle}>
      <input type="checkbox" id={id} checked={flag} onChange={() => trigger()} />
      <label htmlFor={id}></label>
    </div>
  );
}
