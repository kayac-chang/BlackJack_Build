export * from './device';
export * from './format';
export * from './check';
export * from './react';
export * from './url';

export function wait(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export function nextFrame() {
  return new Promise((resolve) => requestAnimationFrame(resolve));
}

export function throttleBy<T>(func: () => Promise<T>) {
  //
  let fetching = false;

  console.log('hi');

  return async function () {
    if (fetching) return;

    fetching = true;

    await func();

    fetching = false;
  };
}

export function random(min: number, max?: number) {
  if (!max) {
    max = min;
    min = 0;
  }

  if (min > max) {
    [min, max] = [max, min];
  }

  return min + (max - min) * Math.random();
}

export function whenVisibilityChange(func: (pass: number) => void) {
  let start = performance.now();

  function handle() {
    if (document.hidden) {
      start = performance.now();

      return;
    }

    const pass = performance.now() - start;

    func(pass);
  }

  handle();

  document.addEventListener('visibilitychange', handle);

  return () => document.removeEventListener('visibilitychange', handle);
}

type CondFunc = () => Promise<boolean>;

export function looper(fn: CondFunc) {
  let flag = true;

  (async function call() {
    if (flag && (await fn())) call();
  })();

  return () => (flag = false);
}

export function toBase64(img: HTMLImageElement) {
  const canvas = document.createElement('canvas');

  canvas.width = img.width;
  canvas.height = img.height;
  canvas.getContext('2d')?.drawImage(img, 0, 0);

  return canvas.toDataURL();
}
