import React from 'react';
import Chip from './components/Chip';
import styles from './Bet.module.scss';
import { CHIP } from '../../models';
import RES from '../../assets';
import { useSelector, useDispatch } from 'react-redux';
import { AppState } from '../../store';
import { choose } from '../../store/actions';

type Props = {
  enable: boolean;
};

type CHIP_IMG =
  | 'CHIP_NORMAL_RED'
  | 'CHIP_NORMAL_GREEN'
  | 'CHIP_NORMAL_BLUE'
  | 'CHIP_NORMAL_BLACK'
  | 'CHIP_NORMAL_PURPLE'
  | 'CHIP_NORMAL_YELLOW';

const chips: { type: CHIP; src: CHIP_IMG }[] = [
  { type: CHIP.RED, src: 'CHIP_NORMAL_RED' },
  { type: CHIP.GREEN, src: 'CHIP_NORMAL_GREEN' },
  { type: CHIP.BLUE, src: 'CHIP_NORMAL_BLUE' },
  { type: CHIP.BLACK, src: 'CHIP_NORMAL_BLACK' },
  { type: CHIP.PURPLE, src: 'CHIP_NORMAL_PURPLE' },
  { type: CHIP.YELLOW, src: 'CHIP_NORMAL_YELLOW' },
];

export default function Chips({ enable }: Props) {
  const dispatch = useDispatch();
  const min = useSelector((state: AppState) => state.game.bet.min);
  const chip = useSelector((state: AppState) => state.bet.chosen?.chip);

  return (
    <div className={styles.section}>
      <div className={styles.field}>
        {chips.map(({ type, src }) => (
          <Chip
            key={type}
            selected={type === chip}
            src={RES.getBase64(src)}
            bet={type * min}
            onClick={() => enable && dispatch(choose({ chip: type, amount: type * min }))}
          />
        ))}
      </div>
    </div>
  );
}
