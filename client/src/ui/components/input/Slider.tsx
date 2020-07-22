import React, { useState, ChangeEvent, useCallback, HTMLAttributes, memo } from 'react';
import styles from './Slider.module.scss';
import clsx from 'clsx';

type Props = {
  min?: number;
  max?: number;
  value?: number;
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
} & HTMLAttributes<HTMLDivElement>;

export default memo(function Slider({ min = 0, max = 100, value = min, onChange, className }: Props) {
  const [val, setValue] = useState(value);

  const handleChange = useCallback(
    (event: ChangeEvent<HTMLInputElement>) => {
      onChange && onChange(event);

      setValue(Number(event.target.value));
    },
    [setValue, onChange]
  );

  return (
    <div className={clsx(styles.wrapper, className)}>
      <input className={styles.slider} type="range" min={min} max={max} value={val} onChange={handleChange} />
      <output className={styles.output} style={{ left: interpret([min, max], val) }}>
        {val}
      </output>
    </div>
  );
});

function interpret([min, max]: [number, number], value: number) {
  const percentage = Number(((value - min) * 100) / (max - min));

  const newPosition = 15 - percentage * 0.4;

  return `calc(${percentage}% + (${newPosition}px))`;
}
