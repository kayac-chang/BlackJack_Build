import React, { ChangeEvent, useCallback } from 'react';
import styles from './Settings.module.scss';
import { VolumeX, Volume2 } from 'react-feather';
import { Slider, Toggle, Select } from '../../components/input';
import { useSoundState, change } from '../../../sound';
import TABLES from '../../../assets/table';
import { useSelector, useDispatch } from 'react-redux';
import store, { AppState } from '../../../store';
import { update } from '../../../store/actions';

type Props = {
  value: number;
  onChange: (event: ChangeEvent<HTMLInputElement>) => void;
};
function Volume({ value, onChange }: Props) {
  return (
    <div className={styles.volume}>
      <VolumeX color="white" />
      <Slider className={styles.slider} value={value} onChange={onChange} />
      <Volume2 color="white" />
    </div>
  );
}

function Audio() {
  const { state, dispatch } = useSoundState();

  const onVolumeChange = useCallback(
    (event: ChangeEvent<HTMLInputElement>) => dispatch(change({ volumn: Number(event.target.value) })),
    [dispatch]
  );
  const onToggleSFX = useCallback(
    (event: ChangeEvent<HTMLInputElement>) => dispatch(change({ canPlaySFX: event.target.checked })),
    [dispatch]
  );
  const onToggleBGM = useCallback(
    (event: ChangeEvent<HTMLInputElement>) => dispatch(change({ canPlayBGM: event.target.checked })),
    [dispatch]
  );

  return (
    <section>
      <div>
        <h3>audio</h3>
      </div>

      <div>
        <h4>volume</h4>
        <Volume value={state.volumn} onChange={onVolumeChange} />
      </div>

      <div className={styles.field}>
        <h4>sound effects</h4>
        <Toggle id={'sound-effects'} value={state.canPlaySFX} onChange={onToggleSFX} />
      </div>

      <div className={styles.field}>
        <h4>ambience sound</h4>
        <Toggle id={'ambience-sound'} value={state.canPlayBGM} onChange={onToggleBGM} />
      </div>
    </section>
  );
}

function Visual() {
  const dispatch = useDispatch();
  const table = useSelector((state: AppState) => state.user.table);

  const mapping = {
    [TABLES.TABLE_BLUE]: 'blue',
    [TABLES.TABLE_RED]: 'red',
    [TABLES.TABLE_GRAY]: 'gray',
    [TABLES.TABLE_GREEN]: 'green',
  };

  const onChange = useCallback(
    (event: ChangeEvent<HTMLSelectElement>) => {
      const mapping = {
        blue: 'TABLE_BLUE',
        red: 'TABLE_RED',
        gray: 'TABLE_GRAY',
        green: 'TABLE_GREEN',
      };

      const table = mapping[event.target.value as keyof typeof mapping] as keyof typeof TABLES;

      const { user } = store.getState();

      dispatch(
        update({
          ...user,
          table,
        })
      );
    },
    [dispatch]
  );

  return (
    <section>
      <div>
        <h3>visual</h3>
      </div>

      <div className={styles.field}>
        <h4>table</h4>
        <Select options={Object.values(mapping)} value={mapping[table]} onChange={onChange} />
      </div>
    </section>
  );
}

export default function Settings() {
  return (
    <div className={styles.settings}>
      <Audio />

      <Visual />
    </div>
  );
}
