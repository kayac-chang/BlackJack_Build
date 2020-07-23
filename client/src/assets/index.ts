import { Loader, LoaderResource } from 'pixi.js';
import { toBase64 } from '../utils';
import { PKG, ASSETS, PRELOAD } from './pkg';
import { Howl } from 'howler';

const loader = new Loader();

loader.pre(SoundHandler);

function SoundHandler(resource: LoaderResource, next: () => void) {
  const SUPPORT_FORMATS = ['mp3', 'opus', 'ogg', 'wav', 'aac', 'm4a', 'm4b', 'mp4', 'webm'];

  if (!SUPPORT_FORMATS.includes(resource.extension)) {
    return next();
  }

  const sound = new Howl({
    src: resource.url,
    onload,
    onloaderror,
  });

  function onload() {
    resource.complete();

    resource.data = sound;

    next();
  }

  function onloaderror(soundId: number, error: any) {
    resource.abort(error);

    console.error(error);

    next();
  }
}

function onProgress(func: (progress: number) => void) {
  loader.onProgress.add(() => func(loader.progress));
}

async function load(pkg: typeof ASSETS | typeof PRELOAD) {
  for (const [name, url] of Object.entries(pkg)) {
    if (loader.resources[name]) {
      continue;
    }

    loader.add(name, url);
  }

  return new Promise((resolve, reject) => {
    loader.load(resolve);

    loader.onError.add(reject);
  });
}

const cache: Record<string, string> = {};

function getBase64(res: PKG) {
  const resource = loader.resources[res];

  if (!cache[res]) {
    cache[res] = toBase64(resource.data);
  }

  return cache[res];
}

function getTexture(res: PKG) {
  return loader.resources[res].texture;
}

function getSound(res: PKG) {
  return loader.resources[res].data as Howl;
}

export default {
  load,
  onProgress,
  getBase64,
  getTexture,
  getSound,
};
