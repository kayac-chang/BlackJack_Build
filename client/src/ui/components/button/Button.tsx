import React, { PropsWithChildren, HTMLAttributes, MouseEvent } from 'react';
import styles from './Button.module.scss';
import clsx from 'clsx';
import { useSoundState, play } from '../../../sound';

type Props = PropsWithChildren<HTMLAttributes<HTMLButtonElement>>;

export function Button({ children, className, onClick, ...props }: Props) {
  const { dispatch } = useSoundState();

  function handleClick(event: MouseEvent<HTMLButtonElement>) {
    onClick && onClick(event);
    dispatch(play({ type: 'sfx', name: 'SFX_TAP' }));
  }

  return (
    <button className={clsx(styles.button, className)} onClick={handleClick} {...props}>
      {children}
    </button>
  );
}
