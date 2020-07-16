import { useState, useEffect } from 'react';

type TriggerFunc = (flag?: boolean | undefined) => void;

export function useTrigger(initState = false): [boolean, TriggerFunc] {
  //
  const [state, setState] = useState(initState);

  function trigger(flag?: boolean) {
    //
    if (flag === undefined) {
      return setState(!state);
    }

    setState(flag);
  }

  return [state, trigger];
}

export function useResize<T>(fn: () => T) {
  const [state, setState] = useState(fn());

  useEffect(() => {
    let id = window.requestAnimationFrame(handler);

    function handler() {
      setState(fn());

      id = window.requestAnimationFrame(handler);
    }

    return () => window.cancelAnimationFrame(id);
  }, [fn]);

  return state;
}
