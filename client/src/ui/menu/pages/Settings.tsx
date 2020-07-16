import React from 'react';
import styles from './Settings.module.scss';
import { VolumeX, Volume2 } from 'react-feather';
import { Slider, Toggle } from '../../components/input';

function Volume() {
  return (
    <div className={styles.volume}>
      <VolumeX color="white" />
      <Slider className={styles.slider} />
      <Volume2 color="white" />
    </div>
  );
}

export default function Settings() {
  return (
    <div className={styles.settings}>
      <section>
        <div>
          <h3>audio</h3>
        </div>
        <div>
          <h4>volume</h4>
          <Volume />
        </div>
        <div className={styles.toggle}>
          <h4>sound effects</h4>
          <Toggle id={'sound-effects'} />
        </div>
        <div className={styles.toggle}>
          <h4>ambience sound</h4>
          <Toggle id={'ambience-sound'} />
        </div>
      </section>
    </div>
  );
}
