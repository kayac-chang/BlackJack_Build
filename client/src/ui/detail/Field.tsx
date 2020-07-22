import React, { memo } from 'react';
import styles from './Detail.module.scss';
import { useSpring, animated } from 'react-spring';
import { Expo } from 'gsap';

type Props = {
  title: string;
  value: string;
};

export default memo(function Field({ title, value }: Props) {
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
});
