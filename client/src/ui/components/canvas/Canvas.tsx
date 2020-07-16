import React, { PropsWithChildren, HTMLProps, memo } from 'react';
import styles from './Canvas.module.scss';
import clsx from 'clsx';

type Props = PropsWithChildren<HTMLProps<HTMLCanvasElement>>;

function Canvas({ children, ...props }: Props) {
  //
  function init(canvas: HTMLCanvasElement) {
    //
    if (!canvas) {
      return;
    }

    if (children && typeof children === 'function') {
      return children(canvas);
    }

    console.error(`Canvas children must be function`);
  }

  return <canvas className={clsx(styles.shadow, styles.fitScreen)} ref={init} {...props} />;
}

export default memo(Canvas);
