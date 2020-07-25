import React, { memo } from 'react';
import styles from './Detail.module.scss';
import { useSelector } from 'react-redux';
import { AppState } from '../../store';
import { currency } from '../../utils';
import History from './History';
import Field from './Field';
import Back from './Back';

export default memo(function RoomDetail() {
  const roomID = useSelector((state: AppState) => state.game.room);
  const roundID = useSelector((state: AppState) => state.game.round);
  const { max, min } = useSelector((state: AppState) => state.game.bet);
  const history = useSelector((state: AppState) => state.room.find(({ id }) => id === roomID)?.history);

  return (
    <div className={styles.detail}>
      <div className={styles.header}>
        <Back />
        <Field title={'room'} value={String(roomID)} />
        <Field title={'round'} value={roundID} />
        <Field title={'bet'} value={`${currency(min)} - ${currency(max)}`} />
      </div>

      <History history={history} />
    </div>
  );
});
