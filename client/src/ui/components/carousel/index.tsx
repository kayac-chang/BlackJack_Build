import { useState, useCallback } from 'react';
import usePagination from '../pagination';
import { useTransition } from 'react-spring';
import { useDrag } from 'react-use-gesture';

export default function useCarousel<T>(source: T[], itemsPerPage: number) {
  const [isNext, setNext] = useState(true);
  const { data, page, range, next, prev } = usePagination(source, itemsPerPage);

  const [block, setBlock] = useState(false);

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
  }, [next, block]);

  const _prev = useCallback(() => {
    if (block) return;
    setNext(false);

    prev();
  }, [prev, block]);

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
