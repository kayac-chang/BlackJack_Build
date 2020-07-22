import React, { ReactNode, useEffect, useState } from 'react';
import styles from './Frame.module.scss';
import { IconContext } from 'react-icons';
import { MdScreenRotation, MdTouchApp } from 'react-icons/md';
import UI from './UI';
import clsx from 'clsx';
import { isBarHidden, isMobile, isFullScreenSupport } from '../utils';

type Props = {
  children: ReactNode;
};

function Rotation() {
  return (
    <IconContext.Provider value={{ size: '50%', className: clsx(styles.rotation) }}>
      <MdScreenRotation />
    </IconContext.Provider>
  );
}

function Scroll() {
  const [hidden, setHidden] = useState(false);

  useEffect(() => {
    window.addEventListener('touchend', (event) => {
      const hidden = isBarHidden();

      hidden && window.scrollTo({ top: 0 });

      setHidden(hidden);
    });

    window.addEventListener('resize', () => {
      const hidden = isBarHidden();

      hidden && window.scrollTo({ top: 0 });

      setHidden(hidden);
    });

    window.addEventListener('orientationchange', () => {
      const hidden = isBarHidden();

      hidden && window.scrollTo({ top: 0 });

      setHidden(hidden);
    });
  }, [setHidden]);

  return (
    <div className={clsx(styles.mask, hidden && styles.hidden)}>
      <IconContext.Provider value={{ size: '50%', className: clsx(styles.scroll) }}>
        <MdTouchApp />
      </IconContext.Provider>
    </div>
  );
}

export default function Frame({ children }: Props) {
  return (
    <div className={styles.frame}>
      <Rotation />

      {isMobile() && !isFullScreenSupport() && <Scroll />}

      <div className={styles.main}>
        {children}
        <UI />
      </div>
    </div>
  );
}
