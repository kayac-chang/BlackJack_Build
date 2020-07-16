import { curry } from 'ramda';

export function Functor(x: any) {
  return {
    map: (fn: Function) => Functor(fn(x)),
    flod: (fn: Function) => fn(x),
  };
}

export const addChild = curry((children, parent) => {
  parent.addChild(...children);

  return parent;
});
