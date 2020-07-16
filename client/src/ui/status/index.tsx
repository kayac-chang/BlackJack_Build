import React, { ReactNode, PropsWithChildren, HTMLAttributes } from 'react';
import style from './Status.module.scss';
import { currency } from '../../utils';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faWallet, faCoins } from '@fortawesome/free-solid-svg-icons';
import { useSelector } from 'react-redux';
import { AppState } from '../../store';
import { animated, useSpring } from 'react-spring';
import { Expo } from 'gsap';

type ButtonProps<T> = PropsWithChildren<T & HTMLAttributes<HTMLButtonElement>>;

type FieldProps = ButtonProps<{
  icon: ReactNode;
  title: string;
  value: number;
}>;

function Field({ icon, title, value }: FieldProps) {
  const props = useSpring({
    to: [{ color: 'rgb(255, 159, 10)' }, { color: '#ffffff' }],
    config: { duration: 250, easing: Expo.easeInOut.easeInOut },
  });

  return (
    <div className={style.field}>
      {icon}
      <div>
        <h5>{title}</h5>
        <animated.span style={props}>{currency(value)}</animated.span>
      </div>
    </div>
  );
}

export default function Status() {
  const balance = useSelector((state: AppState) => state.user.balance);
  const totalBet = useSelector((state: AppState) => state.user.totalBet);

  return (
    <div className={style.status}>
      <Field title={'balance'} value={balance} icon={<FontAwesomeIcon icon={faWallet} />} />
      <Field title={'total bet'} value={totalBet} icon={<FontAwesomeIcon icon={faCoins} />} />
    </div>
  );
}
