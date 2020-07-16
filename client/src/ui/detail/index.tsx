import React, { useCallback } from 'react';
import { Home } from 'react-feather';
import { Button } from '../components/button/Button';
import styles from './Detail.module.scss';
import { useSelector } from 'react-redux';
import { AppState } from '../../store';
import { currency } from '../../utils';
import { useSpring, animated } from 'react-spring';
import { Expo } from 'gsap';
import { useNavigate } from 'react-router-dom';
import services from '../../services';

function Back() {
  const navTo = useNavigate();

  const onClick = useCallback(() => {
    services.leaveRoom();

    navTo(`${process.env.PUBLIC_URL}/lobby`, { replace: true });
  }, [navTo]);

  return (
    <Button className={styles.back} onClick={onClick}>
      <Home />
    </Button>
  );
}

type Props = {
  title: string;
  value: string;
};

function Field({ title, value }: Props) {
  const props = useSpring({
    to: [{ color: 'rgb(255, 159, 10)' }, { color: '#ffffff' }],
    config: { duration: 250, easing: Expo.easeInOut.easeInOut },
  });

  return (
    <div className={styles.field}>
      <h5>{title}</h5>
      <animated.span style={props}>{value}</animated.span>
    </div>
  );
}

export default function RoomDetail() {
  const roomID = useSelector((state: AppState) => state.game.room);
  const roundID = useSelector((state: AppState) => state.game.round);
  const { max, min } = useSelector((state: AppState) => state.game.bet);

  return (
    <div className={styles.detail}>
      <Back />
      <Field title={'room'} value={String(roomID)} />
      <Field title={'round'} value={roundID} />
      <Field title={'bet'} value={`${currency(min)} - ${currency(max)}`} />
    </div>
  );
}
