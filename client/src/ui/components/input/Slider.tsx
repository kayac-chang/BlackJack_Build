import React, { useState, ChangeEvent, useEffect, PropsWithChildren, HTMLAttributes } from 'react';
import styles from './Slider.module.scss';

type Div<T> = PropsWithChildren<T & HTMLAttributes<HTMLDivElement>>;

type Props = Div<{
  min?: number;
  max?: number;
  onValueChange?: (value: number) => void;
}>;

export default function Slider({ min = 0, max = 100, onValueChange, className }: Props) {
  //
  const [value, setValue] = useState(min);

  useEffect(() => {
    //
    if (!onValueChange) return;

    onValueChange(value);
    //
  }, [value, onValueChange]);

  return (
    <div className={`${styles.wrapper} ${className}`}>
      <input className={styles.slider} type="range" min={min} max={max} value={value} onChange={handle} />
      <output className={styles.output} style={{ left: getPos(value) }}>
        {value}
      </output>
    </div>
  );

  function handle(event: ChangeEvent) {
    //
    const el = event.target as HTMLInputElement;

    setValue(Number(el.value));
  }

  function getPos(value: number) {
    //
    const percentage = Number(((value - min) * 100) / (max - min));

    const newPosition = 15 - percentage * 0.4;

    return `calc(${percentage}% + (${newPosition}px))`;
  }
}
