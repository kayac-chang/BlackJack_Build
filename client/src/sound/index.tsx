import { Howler } from 'howler';
import { ContextProvider } from '../utils';
import RES from '../assets';
import { BGM, SFX } from '../assets/sound';

type Sound =
  | {
      type: 'bgm';
      name: keyof typeof BGM;
    }
  | {
      type: 'sfx';
      name: keyof typeof SFX;
    };

type State = {
  bgm?: Sound;
  volumn: number;
  canPlayBGM: boolean;
  canPlaySFX: boolean;
};

type StateProp = {
  volumn?: number;
  canPlayBGM?: boolean;
  canPlaySFX?: boolean;
};

type Action = { type: 'change'; state: StateProp } | { type: 'play'; sound: Sound };

const AMPLIFY = 100;
const init: State = {
  volumn: Howler.volume() * AMPLIFY,
  canPlayBGM: true,
  canPlaySFX: true,
};

function reducer(state: State, action: Action) {
  //
  if (action.type === 'change') {
    //
    if (action.state.volumn !== undefined) {
      Howler.volume(action.state.volumn / AMPLIFY);
    }

    if (action.state.canPlayBGM !== undefined && state.bgm) {
      const bgm = RES.getSound(state.bgm.name);

      bgm.mute(!action.state.canPlayBGM);
    }

    return { ...state, ...action.state };
  }

  if (action.type === 'play' && action.sound.type === 'bgm' && state.canPlayBGM) {
    const bgm = RES.getSound(action.sound.name);

    if (!bgm.playing()) {
      bgm.loop(true).play();
    }

    return { ...state, bgm: action.sound };
  }

  if (action.type === 'play' && action.sound.type === 'sfx' && state.canPlaySFX) {
    const sfx = RES.getSound(action.sound.name);

    sfx.play();

    return { ...state };
  }

  return state;
}

function change(state: StateProp): Action {
  return {
    type: 'change',
    state,
  };
}

function play(sound: Sound): Action {
  return {
    type: 'play',
    sound,
  };
}

const { Provider: SoundProvider, useState: useSoundState } = ContextProvider(init, reducer);
export { SoundProvider, useSoundState, change, play };
