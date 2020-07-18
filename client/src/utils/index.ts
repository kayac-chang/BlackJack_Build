export * from "./device";
export * from "./format";
export * from "./check";

export function getURLParam(key: string) {
  const url = new URL(window.location.href);

  return url.searchParams.get(key) || undefined;
}

export async function getToken() {
  const res = await fetch(`https://api.sunnyland.fun/v1/tokens`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Session: "8b674aec-156c-4894-a102-f3315272f626",
    },
    body: JSON.stringify({
      game: "catpunch",
      username: "mouse1",
    }),
  });

  const { data } = await res.json();

  return data.token.access_token;
}

export function wait(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export function nextFrame() {
  return new Promise((resolve) => requestAnimationFrame(resolve));
}

export function throttleBy<T>(func: () => Promise<T>) {
  //
  let fetching = false;

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

  document.addEventListener("visibilitychange", handle);

  return () => document.removeEventListener("visibilitychange", handle);
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
  const canvas = document.createElement("canvas");

  canvas.width = img.width;
  canvas.height = img.height;
  canvas.getContext("2d")?.drawImage(img, 0, 0);

  return canvas.toDataURL();
}
