import React, { ReactNode, PropsWithChildren, HTMLAttributes, MouseEvent } from 'react';
import styles from './Control.module.scss';
import { Button } from './Button';
import clsx from 'clsx';

type Div<T> = PropsWithChildren<T & HTMLAttributes<HTMLButtonElement>>;

type Props = Div<{
  title: string;
  icon: ReactNode;
  enable?: boolean;
}>;

export default function Control({ title, icon, className, style, onClick, enable = true }: Props) {
  //
  function handle(evt: MouseEvent<HTMLButtonElement>) {
    if (!onClick) return;

    enable && onClick(evt);
  }

  return (
    <div className={clsx(styles.control, className)} style={style}>
      <Button onClick={handle}>{icon}</Button>
      <h5>{title}</h5>
    </div>
  );
}
