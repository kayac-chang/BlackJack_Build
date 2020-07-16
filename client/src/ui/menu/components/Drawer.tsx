import React, { ReactNode, MouseEvent, PropsWithChildren, HTMLAttributes } from 'react';
import styles from './Drawer.module.scss';
import { useSpring, animated } from 'react-spring';
import { Cubic } from 'gsap';
import { Option as OptionButton } from './Button';

type Div<T> = PropsWithChildren<T & HTMLAttributes<HTMLDivElement>>;

type Option = {
  icon: ReactNode;
  title: string;
  onClick: (event: MouseEvent) => void;
};

type DrawerProps = Div<{
  options: Option[];
  open: boolean;
}>;

function Placeholder() {
  return <div className={styles.placeholder}></div>;
}

export default function Drawer({ open, options }: DrawerProps) {
  //
  const anim = useSpring({
    opacity: open ? 1 : 0,
    transform: `translate3d(${open ? 0 : 100}%,0,0)`,
    config: {
      duration: 160,
      easing: Cubic.easeInOut.easeInOut,
    },
  });

  return (
    <animated.div className={styles.drawer} style={anim}>
      <Placeholder />

      {options.map(({ icon, title, onClick }) => (
        <OptionButton open={open} key={title} onClick={onClick}>
          {icon}
        </OptionButton>
      ))}
    </animated.div>
  );
}
