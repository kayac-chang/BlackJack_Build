import { Loader, LoaderResource } from 'pixi.js';
import { toBase64 } from '../utils';
import PKG from './pkg';
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
    // resource.abort(error);

    // console.error(error);

    next();
  }
}

async function load() {
  for (const [name, url] of Object.entries(PKG)) {
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

function getBase64(res: keyof typeof PKG) {
  const resource = loader.resources[res];

  if (!cache[res]) {
    cache[res] = toBase64(resource.data);
  }

  return cache[res];
}

function getTexture(res: keyof typeof PKG) {
  const resource = loader.resources[res];

  return resource.texture;
}

function getSound(res: keyof typeof PKG) {
  return loader.resources[res].data as Howl;
}

export default {
  load,
  getBase64,
  getTexture,
  getSound,
  PKG,
};
