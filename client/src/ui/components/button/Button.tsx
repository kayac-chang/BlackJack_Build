import React, { PropsWithChildren, HTMLAttributes } from 'react';
import styles from './Button.module.scss';

type Props = PropsWithChildren<HTMLAttributes<HTMLButtonElement>>;

export function Button({ children, className, ...props }: Props) {
  //
  const _className = [styles.button, className].filter(Boolean).join(' ');

  return (
    <button className={_className} {...props}>
      {children}
    </button>
  );
}
