import React, { useCallback, useState, memo } from 'react';
import styles from './Detail.module.scss';
import { useSoundState, play } from '../../sound';
import { repeat, identity } from 'ramda';

export default memo(function History({ history = [] }: { history?: string[] }) {
  const [isOpen, toggle] = useState(false);
  const { dispatch } = useSoundState();

  const onToggle = useCallback(() => {
    toggle((isOpen) => !isOpen);

    dispatch(play({ type: 'sfx', name: 'SFX_TOGGLE' }));
  }, [toggle, dispatch]);

  return (
    <div className={styles.history} onClick={onToggle}>
      <div className={styles.content} style={{ width: (isOpen ? 100 : 0) + '%' }}>
        {repeat(identity, 20).map((_, index) => (
          <div key={String(index + 1).padStart(2, '0')} className={styles.record}>
            <h5>{String(index + 1).padStart(2, '0')}</h5>
            <h4>{history[index]}</h4>
          </div>
        ))}
      </div>

      <div className={styles.toggle}>
        <h5>History</h5>
      </div>
    </div>
  );
});
