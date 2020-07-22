import React, { useEffect } from 'react';
import Menu from './menu';
import Status from './status';
import Detail from './detail';
import Bet from './bet';
import Decision from './decision';
import { useLocation } from 'react-router-dom';
import { useSoundState, play } from '../sound';

function useBGM() {
  const location = useLocation();
  const { dispatch } = useSoundState();

  useEffect(() => {
    if (location.pathname.includes('lobby') || location.pathname.includes('game')) {
      return dispatch(play({ type: 'bgm', name: 'BG_MUSIC' }));
    }
  }, [location, dispatch]);
}

export default function UI() {
  const location = useLocation();
  useBGM();

  if (location.pathname.includes('lobby')) {
    return (
      <div className="fixedPage">
        <Status />
        <Menu />
      </div>
    );
  }

  if (location.pathname.includes('game')) {
    return (
      <div className="fixedPage">
        <Decision />
        <Bet />
        <Status />
        <Detail />
        <Menu />
      </div>
    );
  }

  return <></>;
}
