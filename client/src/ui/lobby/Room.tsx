import React, { CSSProperties, memo } from 'react';
import styles from './Lobby.module.scss';
import { times, identity } from 'ramda';
import { Room as Model } from '../../models';
import { animated, useSpring } from 'react-spring';
import RES from '../../assets';

type Props = {
  data: Model;
  style: CSSProperties;
  onClick: () => void;
};

export default memo(function Room({ style, data, onClick }: Props) {
  const roomNum = String(data?.id || '').padStart(2, '0');
  const history = data?.history.slice(0, 20) || [];

  const max = data?.bet.max || '';
  const min = data?.bet.min || '';
  const people = data?.people || 0;

  const props = useSpring(style);

  return (
    <animated.div className={styles.room} style={props} onClick={onClick}>
      <img className={styles.table} src={RES.getBase64('TABLE')} alt={'TABLE'} />
      <img className={styles.detailImg} src={RES.getBase64('DETAIL')} alt={'DETAIL'} />

      <div>
        <img className={styles.numberImg} src={RES.getBase64('ROOM_NUM')} alt={'ROOM_NUM'} />
        <h3 className={styles.number}>{roomNum}</h3>
      </div>

      <div className={styles.detail}>
        <div>
          <h5>MAX</h5>
          <span>{max}</span>
        </div>
        <div>
          <h5>MIN</h5>
          <span>{min}</span>
        </div>
        <div>
          <h5>PEOPLE</h5>
          <span>{people}</span>
        </div>
      </div>

      <div className={styles.history}>
        {times(identity, 20).map((id) => (
          <div key={id}>
            <h5>{id + 1}</h5>
            <h4>{history[id] || ''}</h4>
          </div>
        ))}
      </div>
    </animated.div>
  );
});
