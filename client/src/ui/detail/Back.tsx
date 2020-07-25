import React, { useCallback, memo } from 'react';
import { Home } from 'react-feather';
import { Button } from '../components/button/Button';
import styles from './Detail.module.scss';
import { useNavigate } from 'react-router-dom';
import services from '../../service';
import { useSoundState, play } from '../../sound';

export default memo(function Back() {
  const navTo = useNavigate();

  const { dispatch } = useSoundState();

  const onClick = useCallback(() => {
    services.leaveRoom();

    navTo(`${process.env.PUBLIC_URL}/lobby`, { replace: true });

    dispatch(play({ type: 'sfx', name: 'SFX_NAV_CLOSE' }));
  }, [navTo, dispatch]);

  return (
    <Button className={styles.back} onClick={onClick}>
      <Home />
    </Button>
  );
});
