import React, { useState, ChangeEvent } from 'react';
import styles from './Select.module.scss';

type Props = {
  value?: string;
  options: string[];
  onChange?: (event: ChangeEvent<HTMLSelectElement>) => void;
};

export default function Select({ value, options, onChange }: Props) {
  const [current, setValue] = useState(value || options[0]);

  function handleChange(event: ChangeEvent<HTMLSelectElement>) {
    setValue(event.target.value);

    onChange && onChange(event);
  }

  return (
    <select className={styles.select} value={current} onChange={handleChange}>
      {options.map((value) => (
        <option key={value} value={value}>
          {value}
        </option>
      ))}
    </select>
  );
}
