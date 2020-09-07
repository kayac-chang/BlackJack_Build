import React, { useState, useCallback, ReactNode } from 'react';
import { animated, useTransition } from 'react-spring';
import { useDrag } from 'react-use-gesture';
import { useSoundState, play } from '../../sound';
import usePagination from '../components/pagination';
import styles from './Lobby.module.scss';
import Arrow from './Arrow';

type Props<T> = {
  data: T[];
  children: (data: T[]) => ReactNode;
  focus: number | undefined;
  setFocus: (flag: number | undefined) => void;
};

export default function Carousel<T>({ data: source, children, focus, setFocus }: Props<T>) {
  const { dispatch } = useSoundState();
  const [isNext, setNext] = useState(true);
  const { data, page, range, next, prev } = usePagination(source, 4);

  const gesture = useDrag(({ down, movement: [mx], direction: [xDir] }) => {
    const trigger = Math.abs(mx) > 50;

    if (down || !trigger) {
      return;
    }

    if (xDir < 0) {
      setNext(true);

      next();
    } else {
      setNext(false);

      prev();
    }
  });

  const onPrevClick = useCallback(() => {
    if (focus !== undefined) {
      return setFocus(undefined);
    }

    prev();
    setNext(false);

    dispatch(play({ type: 'sfx', name: 'SFX_NAV_OPEN' }));
  }, [focus, setFocus, prev, dispatch]);

  const onNextClick = useCallback(() => {
    if (focus !== undefined) {
      return setFocus(undefined);
    }

    next();
    setNext(true);

    dispatch(play({ type: 'sfx', name: 'SFX_NAV_CLOSE' }));
  }, [focus, setFocus, next, dispatch]);

  const transitions = useTransition(page, null, {
    from: { opacity: 0, transform: isNext ? `translate3d(100%,0,0)` : `translate3d(-100%,0,0)` },
    enter: { opacity: 1, transform: `translate3d(0%,0,0)` },
    leave: { opacity: 0, transform: isNext ? `translate3d(-100%,0,0)` : `translate3d(100%,0,0)` },
  });

  return (
    <>
      {transitions.map(({ props, key }) => (
        <animated.div key={key} className={styles.rooms} style={props} {...gesture()}>
          {children(data)}
        </animated.div>
      ))}

      {page > range.min && (
        <Arrow
          style={{
            left: `${5}%`,
            top: `${50}%`,
            position: 'absolute',
            transform: `translate(-50%, -50%)`,
          }}
          onClick={onPrevClick}
        />
      )}

      {page < range.max && (
        <Arrow
          style={{
            left: `${95}%`,
            top: `${50}%`,
            position: 'absolute',
            transform: `translate(-50%, -50%) scaleX(-1)`,
          }}
          onClick={onNextClick}
        />
      )}
    </>
  );
}
