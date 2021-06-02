import React from 'react';
import { Plus, Minus, Code, Flag } from 'react-feather';
import { RiSafeLine, RiHandCoinLine } from 'react-icons/ri';
import styles from './Decision.module.scss';
import { DECISION } from '../../models';
import Control from '../components/button/Control';
import services from '../../service';

const config = [
  { item: DECISION.INSURANCE, icon: <RiSafeLine />, className: styles.indigo },
  { item: DECISION.PAY, icon: <RiHandCoinLine />, className: styles.orange },
  { item: DECISION.STAND, icon: <Minus />, className: styles.red },
  { item: DECISION.HIT, icon: <Plus />, className: styles.green },
  { item: DECISION.DOUBLE, icon: <h3>2x</h3>, className: styles.yellow },
  { item: DECISION.SPLIT, icon: <Code />, className: styles.teal },
  { item: DECISION.SURRENDER, icon: <Flag />, className: styles.gray },
];

type Props = {
  enable: boolean;
  decisions: DECISION[];
};
export default function Controls({ decisions, enable }: Props) {
  const include = (item: DECISION) => decisions.includes(item);

  return (
    <div className={styles.section}>
      {config.map(({ item, icon, className }) => (
        <Control
          key={item}
          title={item}
          icon={icon}
          className={className}
          style={{ opacity: enable && include(item) ? 1 : 0.3 }}
          enable={enable && include(item)}
          onClick={() => services.decision(item)}
        />
      ))}
    </div>
  );
}
