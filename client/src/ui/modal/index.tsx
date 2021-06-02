import React, { PropsWithChildren, useCallback } from 'react';
import styles from './Modal.module.scss';
import { ContextProvider } from '../../utils';
import { animated, useTransition } from 'react-spring';

type State = {
  open: boolean;
  title: string;
  msg: string;

  onConfirm?: () => void;
  onClose?: () => void;
};
type Action = { type: 'show'; state: Omit<State, 'open'> } | { type: 'close' };

const initialState = { open: false, title: '', msg: '' };
const { Provider, useState: useModelState } = ContextProvider(initialState, reducer);

function Modal() {
  const { state, dispatch } = useModelState();

  const onCloseClick = useCallback(() => {
    state.onClose && state.onClose();

    dispatch({ type: 'close' });
  }, [state, dispatch]);

  const transitions = useTransition(state.open, null, {
    from: { opacity: 0 },
    enter: { opacity: 1 },
    leave: { opacity: 0 },
    config: { duration: 150 },
  });

  return (
    <>
      {transitions.map(({ item, props, key }) => {
        if (!item) {
          return undefined;
        }

        return (
          <animated.div key={key} className={styles.modal} onClick={onCloseClick} style={props}>
            <div className={styles.content}>
              <div className={styles.header}>
                <h3>{state.title}</h3>
              </div>
              <div className={styles.body}>
                <p>{state.msg}</p>
              </div>
              <div className={styles.footer}>
                {state.onConfirm && (
                  <button className={styles.confirm} onClick={state.onConfirm}>
                    confirm
                  </button>
                )}

                <button className={styles.close} onClick={onCloseClick}>
                  close
                </button>
              </div>
            </div>
          </animated.div>
        );
      })}
    </>
  );
}

function reducer(state: State, action: Action) {
  switch (action.type) {
    case 'close': {
      return { ...state, open: false };
    }
    case 'show': {
      return { ...action.state, open: true };
    }
  }
}

function ModalProvider({ children }: PropsWithChildren<{}>) {
  return (
    <Provider>
      {children}
      <Modal />
    </Provider>
  );
}

export { ModalProvider, useModelState };
