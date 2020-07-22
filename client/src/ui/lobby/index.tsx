import React, { useState, useCallback } from 'react';
import styles from './Lobby.module.scss';
import { useSelector } from 'react-redux';
import { AppState } from '../../store';
import Rooms from './Rooms';

import RES from '../../assets';
import Carousel from './Carousel';

export default function Lobby() {
  const room = useSelector((state: AppState) => state.room);

  const [focus, setFocus] = useState<number | undefined>(undefined);

  const cancelFocus = useCallback(() => {
    if (focus !== undefined) {
      setFocus(undefined);
    }
  }, [focus]);

  return (
    <div className={styles.lobby}>
      <div>
        <img className={styles.background} src={RES.getBase64('BG')} alt={'BG'} onClick={cancelFocus} />

        <Carousel data={room} focus={focus} setFocus={setFocus}>
          {(data) => <Rooms data={data} focus={focus} setFocus={setFocus} />}
        </Carousel>
      </div>
    </div>
  );
}
