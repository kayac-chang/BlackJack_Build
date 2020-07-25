import React, { useState, memo, useCallback, useEffect } from 'react';
import { Room as Model } from '../../models';
import Room from './Room';
import { useNavigate } from 'react-router-dom';
import Arrow from './Arrow';
import styles from './Lobby.module.scss';
import { useSoundState, play } from '../../sound';

const origin = [
  { left: 30, top: 35, scale: 1 },
  { left: 70, top: 35, scale: 1 },
  { left: 30, top: 75, scale: 1 },
  { left: 70, top: 75, scale: 1 },
];

function toStyle(left: number, top: number, scale: number) {
  return {
    left: `${left}%`,
    top: `${top}%`,
    transform: `translate(-50%, -50%) scale(${scale})`,
  };
}

function toFocusStyle(index: number) {
  const focus = { left: 50, top: 50, scale: 1.3 };

  const offset = {
    left: focus.left - origin[index].left,
    top: focus.top - origin[index].top,
  };

  return origin.map((style) => {
    const left = style.left + offset.left;
    const top = style.top + offset.top;

    return {
      left: left + (left - focus.left) * focus.scale,
      top: top + (top - focus.top) * focus.scale,
      scale: focus.scale,
    };
  });
}

type Props = {
  data: Model[];
  focus: number | undefined;
  setFocus: (flag: number | undefined) => void;
};

export default memo(function Rooms({ data, focus, setFocus }: Props) {
  const navTo = useNavigate();
  const [config, setConfig] = useState(origin);
  const { dispatch } = useSoundState();

  const onClick = useCallback(
    (index) => {
      if (focus === undefined) {
        setFocus(index);
        return;
      }

      navTo(`${process.env.PUBLIC_URL}/game/${data[index].id}`, {
        replace: true,
      });
    },
    [focus, setFocus, data, navTo]
  );

  useEffect(() => {
    if (focus === undefined) {
      dispatch(play({ type: 'sfx', name: 'SFX_NAV_OPEN' }));
      setConfig(origin);

      return;
    }

    dispatch(play({ type: 'sfx', name: 'SFX_NAV_CLOSE' }));
    setConfig(toFocusStyle(focus));
  }, [focus, dispatch, setConfig]);

  return (
    <>
      {config.map(({ left, top, scale }, index) => (
        <Room key={String(index)} style={toStyle(left, top, scale)} data={data[index]} onClick={() => onClick(index)} />
      ))}

      {focus !== undefined && (
        <div
          className={styles.moveUp}
          style={{
            left: `${50}%`,
            position: 'absolute',
          }}
        >
          <Arrow
            style={{
              transform: `translate(-50%, -50%) rotate(90deg)`,
            }}
          />
        </div>
      )}
    </>
  );
});
