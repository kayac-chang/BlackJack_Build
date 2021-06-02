import React, { useEffect } from 'react';
import Menu from './menu';
import Status from './status';
import Detail from './detail';
import Bet from './bet';
import Decision from './decision';
import Reward from './reward';
import { useLocation } from 'react-router-dom';
import { useSoundState, play } from '../sound';

function GameView() {
  return (
    <div className="fixedPage">
      <Decision />
      <Bet />
      <Reward />
      <Status />
      <Detail />
      <Menu />
    </div>
  );
}

function LobbyView() {
  return (
    <div className="fixedPage">
      <Status />
      <Menu />
    </div>
  );
}

export default function UI() {
  const location = useLocation();
  const { dispatch } = useSoundState();

  useEffect(() => {
    if (location.pathname.includes('lobby') || location.pathname.includes('game')) {
      dispatch(play({ type: 'bgm', name: 'BG_MUSIC' }));
    }
  }, [location, dispatch]);

  if (location.pathname.includes('lobby')) {
    return <LobbyView />;
  }

  if (location.pathname.includes('game')) {
    return <GameView />;
  }

  return <></>;
}
