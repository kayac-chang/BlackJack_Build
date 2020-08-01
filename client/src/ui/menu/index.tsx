import React, { ReactNode, useState, useCallback } from 'react';
import styles from './Menu.module.scss';
import Drawer from './components/Drawer';
import { Trigger } from './components/Button';
import { Settings, Info, LogOut } from 'react-feather';
import { SettingsPage, GameRulesPage } from './pages';
import clsx from 'clsx';
import { useModelState } from '../modal';
import { useNavigate } from 'react-router-dom';
import { isLocalStorageSupport } from '../../utils';

// ===== Menu =====
export default function Menu() {
  const [page, setPage] = useState<ReactNode | undefined>();
  const [isDrawerOpen, setDrawerOpen] = useState(false);
  const navTo = useNavigate();
  const { dispatch } = useModelState();

  const [options] = useState([
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
    // TODO: History
    // {
    //   icon: <Clock />,
    //   title: 'history',
    //   onClick: () => setPage(<HistoryPage />),
    // },
    {
      icon: <LogOut />,
      title: 'home',
      onClick: () =>
        dispatch({
          type: 'show',
          state: {
            title: 'Back to home',
            msg: 'are you sure to exit?',
            onConfirm: () => {
              if (isLocalStorageSupport()) {
                const lobby = localStorage.getItem('lobby');

                if (lobby) {
                  return (window.location.href = lobby);
                }
              }

              return navTo(-1);
            },
          },
        }),
    },
  ]);

  const onTrigger = useCallback(() => {
    setDrawerOpen((flag) => !flag);
    setPage(undefined);
  }, [setDrawerOpen, setPage]);

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
