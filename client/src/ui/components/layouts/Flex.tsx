import React, { PropsWithChildren, HTMLAttributes } from 'react';
import styles from './Flex.module.scss';
import clsx from 'clsx';

type Props = PropsWithChildren<HTMLAttributes<HTMLDivElement>>;

export function Center({ children, className, ...props }: Props) {
  return (
    <Flex className={clsx(styles.center, className)} {...props}>
      {children}
    </Flex>
  );
}

export function Flex({ children, className, ...props }: Props) {
  return (
    <div className={clsx(styles.default, className)} {...props}>
      {children}
    </div>
  );
}
