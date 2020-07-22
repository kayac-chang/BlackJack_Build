import React, { useState, useCallback, ReactNode } from 'react';
import { animated, useTransition } from 'react-spring';
import { useDrag } from 'react-use-gesture';
import { useSoundState, play } from '../../sound';
import usePagination from '../components/pagination';
import styles from './Lobby.module.scss';
import Arrow from './Arrow';

function useCarousel<T>(source: T[], itemsPerPage: number) {
  const [isNext, setNext] = useState(true);
  const { data, page, range, next, prev } = usePagination(source, itemsPerPage);

  const [block, setBlock] = useState(false);

  const { dispatch } = useSoundState();

  const transitions = useTransition(page, {
    from: { opacity: 0, transform: isNext ? `translate3d(100%,0,0)` : `translate3d(-100%,0,0)` },
    enter: { opacity: 1, transform: `translate3d(0%,0,0)` },
    leave: { opacity: 0, transform: isNext ? `translate3d(-100%,0,0)` : `translate3d(100%,0,0)` },
    onStart: () => setBlock(true),
    onRest: () => setBlock(false),
  });

  const _next = useCallback(() => {
    if (block) return;
    setNext(true);

    next();

    dispatch(play({ type: 'sfx', name: 'SFX_NAV_OPEN' }));
  }, [next, block, dispatch]);

  const _prev = useCallback(() => {
    if (block) return;
    setNext(false);

    prev();

    dispatch(play({ type: 'sfx', name: 'SFX_NAV_CLOSE' }));
  }, [prev, block, dispatch]);

  const gesture = useDrag(({ down, movement: [mx], direction: [xDir] }) => {
    const trigger = Math.abs(mx) > 50;

    if (down || !trigger) {
      return;
    }

    xDir < 0 ? _next() : _prev();
  });

  return {
    data,
    page,
    range,
    transitions,
    gesture,
    next: _next,
    prev: _prev,
  };
}

type Props<T> = {
  data: T[];
  children: (data: T[]) => ReactNode;
  focus: number | undefined;
  setFocus: (flag: number | undefined) => void;
};

export default function Carousel<T>({ data: source, children, focus, setFocus }: Props<T>) {
  const { data, page, range, transitions, next, prev, gesture } = useCarousel(source, 4);

  const onPrevClick = useCallback(() => {
    if (focus !== undefined) {
      return setFocus(undefined);
    }

    prev();
  }, [focus, setFocus, prev]);

  const onNextClick = useCallback(() => {
    if (focus !== undefined) {
      return setFocus(undefined);
    }

    next();
  }, [focus, setFocus, next]);

  return (
    <>
      {transitions((prop) => (
        <animated.div className={styles.rooms} style={prop} {...gesture()}>
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
