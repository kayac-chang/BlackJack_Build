import React, { useState, ChangeEvent, useCallback } from 'react';
import styles from './Toggle.module.scss';
import { useSoundState, play } from '../../../sound';

type Props = {
  id?: string;
  value?: boolean;
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
};

export default function Toggle({ id, value = false, onChange }: Props) {
  const [flag, setFlag] = useState(value);
  const { dispatch } = useSoundState();

  const handleChange = useCallback(
    function (event: ChangeEvent<HTMLInputElement>) {
      onChange && onChange(event);

      const flag = event.target.checked;

      setFlag(flag);

      dispatch(play({ type: 'sfx', name: 'SFX_TOGGLE' }));
    },
    [onChange, setFlag, dispatch]
  );

  return (
    <div className={styles.toggle}>
      <input type="checkbox" id={id} checked={flag} onChange={handleChange} />
      <label htmlFor={id}></label>
    </div>
  );
}
