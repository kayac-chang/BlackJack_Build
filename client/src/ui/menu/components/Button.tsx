import React, { PropsWithChildren, HTMLAttributes } from 'react';
import { Menu as IconMenu, X } from 'react-feather';
import { Button } from '../../components/button/Button';
import styles from './Button.module.scss';
import clsx from 'clsx';

// ===== Trigger =====
type ButtonProps<T> = PropsWithChildren<T & HTMLAttributes<HTMLButtonElement>>;

type Props = ButtonProps<{
  open: boolean;
}>;

export function Trigger({ open, style, onClick }: Props) {
  return (
    <Button className={clsx(styles.trigger, open && styles.open)} onClick={onClick} style={style}>
      {open ? <X /> : <IconMenu />}
    </Button>
  );
}

export function Option({ open, children, onClick }: Props) {
  return (
    <Button className={clsx(styles.option, open && styles.open)} onClick={onClick}>
      {children}
    </Button>
  );
}
