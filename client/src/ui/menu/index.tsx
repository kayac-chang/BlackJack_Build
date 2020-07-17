import React, { ReactNode, useState } from 'react';
import styles from './Menu.module.scss';
import Drawer from './components/Drawer';
import { Trigger } from './components/Button';
import { Settings, Info, Clock, LogOut } from 'react-feather';
import { SettingsPage, HistoryPage, GameRulesPage } from './pages';
import clsx from 'clsx';

// ===== Menu =====
export default function Menu() {
  const [page, setPage] = useState<ReactNode | undefined>();
  const [isDrawerOpen, setDrawerOpen] = useState(false);

  const options = [
    {
      icon: <Info />,
      title: 'rules',
      onClick: () => setPage(<GameRulesPage />),
    },
    {
      icon: <Settings />,
      title: 'settings',
      onClick: () => setPage(<SettingsPage />),
    },
    {
      icon: <Clock />,
      title: 'history',
      onClick: () => setPage(<HistoryPage />),
    },
    {
      icon: <LogOut />,
      title: 'home',
      onClick: () => console.log('home'),
    },
  ];

  function onTrigger() {
    setDrawerOpen(!isDrawerOpen);
    setPage(undefined);
  }

  return (
    <>
      <Trigger style={{ right: 0 }} open={isDrawerOpen} onClick={onTrigger} />

      <div className={styles.menu} style={{ pointerEvents: isDrawerOpen ? 'all' : 'none' }}>
        {isDrawerOpen && (
          <div className={clsx(styles.page, page || styles.hidden)} onClick={() => !page && onTrigger()}>
            {page}
          </div>
        )}

        <Drawer options={options} open={isDrawerOpen} />
      </div>
    </>
  );
}
