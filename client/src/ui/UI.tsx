import React from 'react';
import Menu from './menu';
import Status from './status';
import Detail from './detail';
import Bet from './bet';
import Decision from './decision';
import { useLocation } from 'react-router-dom';

export default function UI() {
  const location = useLocation();

  const inLobby = location.pathname.includes('lobby');

  return (
    <div className="fixedPage" style={{ pointerEvents: 'none' }}>
      <Menu />
      <Status />
      {!inLobby && <Detail />}
      {!inLobby && <Decision />}
      {!inLobby && <Bet />}
    </div>
  );
}
